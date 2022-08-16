package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jacstn/golang-url-shortner/config"
	"github.com/jacstn/golang-url-shortner/pkg/forms"
	"github.com/jacstn/golang-url-shortner/pkg/models"
)

var app *config.AppConfig

func NewHandlers(c *config.AppConfig) {
	app = c
}

func Home(w http.ResponseWriter, r *http.Request) {

	strMap := make(map[string]interface{})
	strMap["variable"] = "variable to display"

	app.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)

	renderTemplate(w, "home", &models.TemplateData{
		Data: strMap,
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
	strMap := make(map[string]interface{})
	strMap["new_data"] = "new data passed before POST"
	renderTemplate(w, "new-url", &models.TemplateData{
		Form: forms.New(nil),
		Data: strMap,
	})
}

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	form := forms.New(r.PostForm)

	url := models.Url{}
	url.Url = r.Form.Get("surl")
	data := make(map[string]interface{})
	data["url_info"] = url
	data["new_data"] = "new data after POST"
	form.Has("surl", r)
	if !form.Valid() {

		renderTemplate(w, "new-url", &models.TemplateData{
			Data: data,
			Form: form,
		})
		return
	}
	fmt.Println("create url", ":", url.Url)
	renderTemplate(w, "new-url", &models.TemplateData{
		Data: data,
		Form: form,
	})
}

func renderTemplate(w http.ResponseWriter, templateName string, data *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateName+".go.tmpl", "./templates/base.layout.go.tmpl")

	err := parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error handling template page!!", err)
	}
}
