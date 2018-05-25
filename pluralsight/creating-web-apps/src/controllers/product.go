package controllers


import (
	"net/http"
	"text/template"

	"github.com/GOCODE/pluralsight/creating-web-apps/src/viewmodels"
	"github.com/GOCODE/pluralsight/creating-web-apps/src/controllers/util"
	"github.com/gorilla/mux"
	"strconv"
)




type productController struct {
	template *template.Template
}



func (this *productController) get(w http.ResponseWriter, req *http.Request ) {

     vars := mux.Vars(req)

     idRaw := vars["id"]

     id,err := strconv.Atoi(idRaw)

    if err == nil {
		vm := viewmodels.GetProduct(id)

	responseWriter := util.GetResponseWriter(w,req)
	defer responseWriter.Close()
    responseWriter.Header().Add("Content Type", "text/html")
	
	this.template.Execute(responseWriter,vm)		} else {
			w.WriteHeader(404)
	}
	
}