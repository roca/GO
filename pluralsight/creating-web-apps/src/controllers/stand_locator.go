package controllers


import (
	"net/http"
	"text/template"
	"encoding/json"

	"github.com/GOCODE/pluralsight/creating-web-apps/src/viewmodels"
	"github.com/GOCODE/pluralsight/creating-web-apps/src/controllers/util"
)



type standLocatorController struct {
	template *template.Template
}



func (this *standLocatorController) get(w http.ResponseWriter, req *http.Request ) {
	responseWriter := util.GetResponseWriter(w,req)
	defer responseWriter.Close()

	vm := viewmodels.GetStandLocator()

    responseWriter.Header().Add("Content Type", "text/html")
    this.template.Execute(responseWriter,vm)

}

func(this *standLocatorController) apiSearch(w http.ResponseWriter, req *http.Request ) {
	responseWriter := util.GetResponseWriter(w,req)
	defer responseWriter.Close()

	vm := viewmodels.GetStandLocations()

    responseWriter.Header().Add("Content Type", "application/json")
    
    data, err := json.Marshal(vm)

    if err == nil {
    	responseWriter.Write(data)
    } else {
    	responseWriter.WriteHeader(404)
    }

}

