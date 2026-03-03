package main

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v3"
)

// KubeConfig represents the top-level kubeconfig structure.
type KubeConfig struct {
	APIVersion     string         `yaml:"apiVersion"`
	Kind           string         `yaml:"kind"`
	CurrentContext string         `yaml:"current-context"`
	Clusters       []ClusterEntry `yaml:"clusters"`
	Contexts       []ContextEntry `yaml:"contexts"`
	Users          []UserEntry    `yaml:"users"`
}

type ClusterEntry struct {
	Name    string        `yaml:"name"`
	Cluster ClusterDetail `yaml:"cluster"`
}

type ClusterDetail struct {
	Server                string `yaml:"server"`
	CertificateAuthority  string `yaml:"certificate-authority"`
	CertificateAuthorData string `yaml:"certificate-authority-data"`
}

type ContextEntry struct {
	Name    string        `yaml:"name"`
	Context ContextDetail `yaml:"context"`
}

type ContextDetail struct {
	Cluster   string `yaml:"cluster"`
	Namespace string `yaml:"namespace"`
	User      string `yaml:"user"`
}

type UserEntry struct {
	Name string     `yaml:"name"`
	User UserDetail `yaml:"user"`
}

type UserDetail struct {
	Exec  *ExecConfig `yaml:"exec"`
	Token string      `yaml:"token"`
}

type ExecConfig struct {
	APIVersion string    `yaml:"apiVersion"`
	Command    string    `yaml:"command"`
	Args       []string  `yaml:"args"`
	Env        []EnvVar  `yaml:"env"`
}

type EnvVar struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func main() {
	a := app.New()
	w := a.NewWindow("Kubeconfig Viewer")
	w.Resize(fyne.NewSize(800, 600))

	kubeconfig, err := loadKubeConfig()
	if err != nil {
		w.SetContent(widget.NewLabel(fmt.Sprintf("Error loading kubeconfig: %v", err)))
		w.ShowAndRun()
		return
	}

	tabs := container.NewAppTabs(
		container.NewTabItem("Clusters", buildClustersTab(kubeconfig)),
		container.NewTabItem("Contexts", buildContextsTab(kubeconfig)),
		container.NewTabItem("Users", buildUsersTab(kubeconfig)),
	)

	w.SetContent(tabs)
	w.ShowAndRun()
}

func loadKubeConfig() (*KubeConfig, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("getting home dir: %w", err)
	}

	data, err := os.ReadFile(filepath.Join(home, ".kube", "config"))
	if err != nil {
		return nil, fmt.Errorf("reading kubeconfig: %w", err)
	}

	var cfg KubeConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing kubeconfig: %w", err)
	}

	return &cfg, nil
}

func buildClustersTab(cfg *KubeConfig) fyne.CanvasObject {
	var cards []fyne.CanvasObject
	for _, c := range cfg.Clusters {
		certInfo := "embedded (base64)"
		if c.Cluster.CertificateAuthority != "" {
			certInfo = c.Cluster.CertificateAuthority
		}

		form := widget.NewForm(
			widget.NewFormItem("Server", widget.NewLabel(c.Cluster.Server)),
			widget.NewFormItem("CA", widget.NewLabel(certInfo)),
		)

		card := widget.NewCard(c.Name, "", form)
		cards = append(cards, card)
	}

	return container.NewVScroll(container.NewVBox(cards...))
}

func buildContextsTab(cfg *KubeConfig) fyne.CanvasObject {
	var cards []fyne.CanvasObject
	for _, c := range cfg.Contexts {
		ns := c.Context.Namespace
		if ns == "" {
			ns = "default"
		}

		form := widget.NewForm(
			widget.NewFormItem("Cluster", widget.NewLabel(c.Context.Cluster)),
			widget.NewFormItem("Namespace", widget.NewLabel(ns)),
			widget.NewFormItem("User", widget.NewLabel(c.Context.User)),
		)

		isCurrent := c.Name == cfg.CurrentContext
		var card fyne.CanvasObject
		if isCurrent {
			title := canvas.NewText("★ "+c.Name+" (current)", color.RGBA{R: 0, G: 180, B: 0, A: 255})
			title.TextStyle = fyne.TextStyle{Bold: true}
			title.TextSize = 16
			card = container.NewVBox(title, form)
		} else {
			card = widget.NewCard(c.Name, "", form)
		}

		cards = append(cards, card)
	}

	return container.NewVScroll(container.NewVBox(cards...))
}

func buildUsersTab(cfg *KubeConfig) fyne.CanvasObject {
	var cards []fyne.CanvasObject
	for _, u := range cfg.Users {
		var items []*widget.FormItem

		if u.User.Token != "" {
			items = append(items, widget.NewFormItem("Auth", widget.NewLabel("token")))
		}

		if u.User.Exec != nil {
			items = append(items, widget.NewFormItem("Command", widget.NewLabel(u.User.Exec.Command)))
			if len(u.User.Exec.Args) > 0 {
				items = append(items, widget.NewFormItem("Args", widget.NewLabel(strings.Join(u.User.Exec.Args, " "))))
			}
			for _, env := range u.User.Exec.Env {
				items = append(items, widget.NewFormItem("Env: "+env.Name, widget.NewLabel(env.Value)))
			}
		}

		if len(items) == 0 {
			items = append(items, widget.NewFormItem("Auth", widget.NewLabel("(no details)")))
		}

		form := widget.NewForm(items...)
		card := widget.NewCard(u.Name, "", form)
		cards = append(cards, card)
	}

	return container.NewVScroll(container.NewVBox(cards...))
}
