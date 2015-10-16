package controllers


import (
	"net/http"
	"text/template"

	"github.com/GOCODE/pluralsight/creating-web-apps/src/viewmodels"
)


type categoriesController struct {
	template *template.Template
}


func (this *categoriesController) get(w http.ResponseWriter, req *http.Request ) {

	vm := viewmodels.GetCategories()

	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w,vm)
	
}