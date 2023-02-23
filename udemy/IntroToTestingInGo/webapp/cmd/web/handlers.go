package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/roca/go-toolkit/v2"
)

var pathToTemplates = "./templates/"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
}

type TemplateData struct {
	IP   string
	Data map[string]interface{}
}

func (app *application) render(w http.ResponseWriter, r *http.Request, t string, data *TemplateData) error {
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}
	data.IP = app.ipFromContext(r.Context())
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {

	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var tools toolkit.Tools

	err := tools.ReadJSON(w, r, &data)
	if err != nil {
		return
	}

	fmt.Println(data.Email, data.Password)

}
