package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/roca/celeritas/mailer"
)

func (a *application) routes() *chi.Mux {
	// middle must come before any routes
	a.use(a.Middleware.CheckRemember)

	// add routes here
	a.get("/", a.Handlers.Home)
	a.get("/go-page", a.Handlers.GoPage)
	a.get("/jet-page", a.Handlers.JetPage)
	a.get("/sessions", a.Handlers.SessionsTest)

	a.get("/users/login", a.Handlers.UserLogin)
	a.post("/users/login", a.Handlers.PostUserLogin)
	a.get("/users/logout", a.Handlers.Logout)
	a.get("/users/forgot-password", a.Handlers.Forgot)
	a.post("/users/forgot-password", a.Handlers.PostForgot)

	a.get("/form", a.Handlers.Form)
	a.post("/form", a.Handlers.PostForm)

	a.get("/json", a.Handlers.JSON)
	a.get("/xml", a.Handlers.XML)
	a.get("/download-file", a.Handlers.DownloadFile)

	a.get("/create-user", a.Handlers.CreateUser)
	a.get("/get-user/{id}", a.Handlers.GetUserByID)
	a.get("/update-user/{id}", a.Handlers.UpdateUserByID)
	a.get("/test_crypto", a.Handlers.TestCrypto)
	a.get("/test_cache", a.Handlers.ShowCachePage)
	a.post("/api/save-in-cache", a.Handlers.SaveInCache)
	a.post("/api/get-from-cache", a.Handlers.GetFromCache)
	a.post("/api/delete-from-cache", a.Handlers.DeleteFromCache)
	a.post("/api/empty-cache", a.Handlers.EmptyCache)

	a.get("/test_mail", func(w http.ResponseWriter, r *http.Request) {
		msg := mailer.Message{
			From:        "test@example.com",
			To:          "you@there.com",
			Subject:     "Test Subject - sent using channel",
			Template:    "test",
			Attachments: nil,
			Data:        nil,
		}

		// a.App.Mail.Jobs <- msg
		// res := <-a.App.Mail.Results
		// if res.Error != nil {
		// 	a.App.ErrorLog.Println(res.Error)
		// }
		err := a.App.Mail.SendSMTPMessage(msg)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}
		fmt.Fprint(w, "Sent mail!")
	})

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
