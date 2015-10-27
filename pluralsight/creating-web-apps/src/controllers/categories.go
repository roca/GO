package controllers


import (
	"net/http"
	"text/template"

	"github.com/GOCODE/pluralsight/creating-web-apps/src/viewmodels"
	"github.com/GOCODE/pluralsight/creating-web-apps/src/controllers/util"
	"github.com/gorilla/mux"
	"github.com/GOCODE/pluralsight/creating-web-apps/src/models"
	"github.com/GOCODE/pluralsight/creating-web-apps/src/converters"
	"strconv"
)


type categoriesController struct {
	template *template.Template
}


func (this *categoriesController) get(w http.ResponseWriter, req *http.Request ) {

	categories := models.GetCategories()
	
	categoriesVM := []viewmodels.Category{}
	for _, category := range categories {
		categoriesVM = append(categoriesVM, converters.ConvertCategoyToViewModel(category))
	}

	vm := viewmodels.GetCategories()
	vm.Categories = categoriesVM


	responseWriter := util.GetResponseWriter(w,req)
	defer responseWriter.Close()
    responseWriter.Header().Add("Content Type", "text/html")
	
	this.template.Execute(responseWriter,vm)	
}

type categoryController struct {
	template *template.Template
}



func (this *categoryController) get(w http.ResponseWriter, req *http.Request ) {

     vars := mux.Vars(req)

     idRaw := vars["id"]

     id,err := strconv.Atoi(idRaw)

    if err == nil {
		vm := viewmodels.GetProducts(id)

	responseWriter := util.GetResponseWriter(w,req)
	defer responseWriter.Close()
    responseWriter.Header().Add("Content Type", "text/html")
	
	this.template.Execute(responseWriter,vm)
	} else {
			w.WriteHeader(404)
	}
	
}