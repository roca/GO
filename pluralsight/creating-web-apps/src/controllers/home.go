package controllers

import (
	"net/http"
	"text/template"

	"github.com/GOCODE/pluralsight/creating-web-apps/src/controllers/util"
	"github.com/GOCODE/pluralsight/creating-web-apps/src/models"
	"github.com/GOCODE/pluralsight/creating-web-apps/src/viewmodels"
)

type homeController struct {
	template      *template.Template
	loginTemplate *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {

	vm := viewmodels.GetHome()

	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	responseWriter.Header().Add("Content Type", "text/html")

	this.template.Execute(responseWriter, vm)

}

func (this *homeController) login(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content Type", "text/html")

	vm := viewmodels.GetLogin()

	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")
		member, err := models.GetMember(email, password)
	}

	this.loginTemplate.Execute(responseWriter, vm)

}
