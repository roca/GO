package controllers


import (
	"net/http"
	"text/template"

	"github.com/GOCODE/pluralsight/creating-web-apps/src/viewmodels"
)


type homeController struct {
	template *template.Template
}


func (this *homeController) get(w http.ResponseWriter, req *http.Request ) {

	vm := viewmodels.GetHome()

	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w,vm)
	
}