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

	strMap := make(map[string]string)
	strMap["variable"] = "variable to display"
	tmplData := models.TemplateData{
		Data: strMap,
	}
	app.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)

	renderTemplate(w, "home", &tmplData)
}

func About(w http.ResponseWriter, r *http.Request) {
	ip := app.Session.GetString(r.Context(), "remote_ip")
	data := models.TemplateData{}
	fmt.Printf(ip)
	//models.TemplateData{ data["remote_ip"] = string(ip)

	renderTemplate(w, "about", &data)
}

func NewUrl(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "new-url", nil)
}

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// model := models.Url{
	// 	Url: r.Form.Get("surl"),
	// }
	form := forms.New(r.PostForm)

	if form.Has("surl", r) == false {

	}

	fmt.Println("create url", ":")
	renderTemplate(w, "new-url", nil)
}

func renderTemplate(w http.ResponseWriter, templateName string, data *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateName+".go.tmpl", "./templates/base.layout.go.tmpl")

	err := parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "This is about page")
	}
}
