package handlers

import (
	"encoding/xml"
	"fmt"
	"myapp/data"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/roca/celeritas"
)

type Handlers struct {
	App    *celeritas.Celeritas
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) GoPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.GoPage(w, r, "home", nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) JetPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.JetPage(w, r, "jet-template", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) SessionsTest(w http.ResponseWriter, r *http.Request) {
	myData := "this is my data"

	h.sessionPut(r.Context(), "myData", myData)

	myValue := h.App.Session.GetString(r.Context(), "myData")

	vars := make(jet.VarMap)
	vars.Set("myData", myValue)

	err := h.App.Render.JetPage(w, r, "sessions", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) JSON(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		ID      int64    `json:"id"`
		Name    string   `json:"name"`
		Hobbies []string `json:"hobbies"`
	}{
		ID:      10,
		Name:    "John Doe",
		Hobbies: []string{"hiking", "biking", "swimming"},
	}

	err := h.App.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println("error rendering JSON:", err)
	}
}

func (h *Handlers) XML(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		XMLName xml.Name `xml:"person"`
		ID      int64    `xml:"id"`
		Name    string   `xml:"name"`
		Hobbies []string `xml:"hobbies>hobby"`
	}{
		ID:      10,
		Name:    "John Doe",
		Hobbies: []string{"hiking", "biking", "swimming"},
	}

	err := h.App.WriteXML(w, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println("error rendering XML:", err)
	}
}

func (h *Handlers) DownloadFile(w http.ResponseWriter, r *http.Request) {
	err := h.App.DownloadFile(w, r, "./public/images", "celeritas.jpg")
	if err != nil {
		h.App.ErrorLog.Println("error downloading file:", err)
	}
}

func (h *Handlers) TestCrypto(w http.ResponseWriter, r *http.Request) {
	plainText := "Hello, world"
	fmt.Fprint(w, "Unencrpted: " + plainText + "\n")
	encrypted, err := h.encrypt(plainText)
	if err != nil {
		h.App.ErrorLog.Println(err)
		h.App.Error500(w,r)
		return
	}

	fmt.Fprint(w,"Encrypted: " + encrypted + "\n")

	decrypted, err := h.decrypt(encrypted)
	if err != nil {
		h.App.ErrorLog.Println(err)
		h.App.Error500(w,r)
		return
	}

	fmt.Fprint(w,"Decrypted: " + decrypted + "\n")
}
