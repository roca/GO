package controllers

import (
	"fmt"
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

	fmt.Println(req.Method)

	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")
		member, err := models.GetMember(email, password)

		fmt.Println(member)

		if err == nil {
			session, err := models.CreateSession(member)
			if err == nil {
				var cookie http.Cookie
				cookie.Name = "sessionId"
				cookie.Value = session.SessionId()
				fmt.Println(cookie)
				responseWriter.Header().Add("Set-Cookie", cookie.String())
			} else {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	this.loginTemplate.Execute(responseWriter, vm)

}
