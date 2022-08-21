package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jacstn/golang-url-shortner/config"
	"github.com/jacstn/golang-url-shortner/internal/forms"
	"github.com/jacstn/golang-url-shortner/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewHandlers(c *config.AppConfig) {
	app = c
}

func Home(w http.ResponseWriter, r *http.Request) {

	displayData := make(map[string]interface{})
	displayData["variable"] = "variable to display"

	app.Session.Put(r.Context(), "remote_ip", r.Host)

	var urls []models.Url
	urls = models.ListUrls(app.DB)

	displayData["list_of_urls"] = urls

	renderTemplate(w, "home", &models.TemplateData{
		Data: displayData,
	})
}

func About(w http.ResponseWriter, r *http.Request) {
	ip := app.Session.GetString(r.Context(), "remote_ip")
	data := models.TemplateData{}
	fmt.Printf(ip)
	//models.TemplateData{ data["remote_ip"] = string(ip)

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

	form := forms.New(r.PostForm)
	form.Has("surl", r)
	form.ValidUrl("surl", r)
	urlModel := models.Url{Name: r.Form.Get("surl")}

	data := make(map[string]interface{})
	data["csrf_token"] = nosurf.Token(r)
	data["url_model"] = urlModel
	models.SaveUrl(app.DB, urlModel)

	renderTemplate(w, "new-url", &models.TemplateData{
		Form: form,
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
