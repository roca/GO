package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v2"
)

var (
	kubeConfigPath string
)

func main() {
	defaultPath := fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))
	flag.StringVar(&kubeConfigPath, "kubeConfigPath", defaultPath, "Path to kube config file")
	flag.Parse()

	// Read the file
	data, err := ioutil.ReadFile(kubeConfigPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(data))

	var config kubeconfig
	//var config map[string]interface{}

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data
	fmt.Println(config)

	a := app.New()
	w := a.NewWindow("Kubeconfig Editor")
	w.Resize(fyne.NewSize(400, 400))

	//text := widget.NewRichTextWithText(string(data))
	treeMap := make(map[string][]string)

	tree := widget.NewTreeWithStrings(treeMap)

	c := container.NewMax(container.NewScroll(tree))

	treeMap["config"] = []string{"apiVersion: " + config.ApiVersion, "kind: " + config.Kind}
	tree.Root = widget.TreeNodeID("config")
	
	w.SetContent(c)

	w.ShowAndRun()
}
