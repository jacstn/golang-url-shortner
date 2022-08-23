package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jacstn/golang-url-shortner/config"
	"github.com/jacstn/golang-url-shortner/internal/forms"
	"github.com/jacstn/golang-url-shortner/internal/helpers"
	"github.com/jacstn/golang-url-shortner/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewHandlers(c *config.AppConfig) {
	app = c
}

func Home(w http.ResponseWriter, r *http.Request) {

	displayData := make(map[string]interface{})

	app.Session.Put(r.Context(), "remote_ip", r.Host)
	displayData["list_of_urls"] = models.ListUrls(app.DB)

	renderTemplate(w, "home", &models.TemplateData{
		Data: displayData,
	})
}

func About(w http.ResponseWriter, r *http.Request) {
	//ip := app.Session.GetString(r.Context(), "remote_ip")
	data := models.TemplateData{}

	renderTemplate(w, "about", &data)
}

func NewUrl(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["url_model"] = models.Url{}
	data["csrf_token"] = nosurf.Token(r)

	renderTemplate(w, "new-url", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	app.Session.Pop(r.Context(), "saved_id")
	form := forms.New(r.PostForm)
	form.Has("surl", r)
	form.ValidUrl("surl", r)
	urlModel := models.Url{Name: r.Form.Get("surl")}

	data := make(map[string]interface{})
	data["csrf_token"] = nosurf.Token(r)
	data["url_model"] = urlModel

	id := models.SaveUrl(app.DB, urlModel)
	if id > 0 {
		fmt.Println(id)
		app.Session.Put(r.Context(), "saved_id", id)
		ViewUrl(w, r)
		return
	}
	fmt.Println("saving url error", id)

	renderTemplate(w, "new-url", &models.TemplateData{
		Form: form,
		Data: data,
	})
}

func ViewUrl(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	id := app.Session.GetString(r.Context(), "saved_id")
	fmt.Println("got id from session", id)
	url := models.GetUrlById(app.DB, id)

	if url.Id == 0 {
		data["link"] = "URL not found"
	} else {
		data["link"] = r.Host + "/" + helpers.IntToCode(int(url.Id), app.CharArr)
	}

	renderTemplate(w, "view-url", &models.TemplateData{
		Data: data,
	})
}

func renderTemplate(w http.ResponseWriter, templateName string, data *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateName+".go.tmpl", "./templates/base.layout.go.tmpl")

	err := parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error handling template page!!", err)
	}
}
