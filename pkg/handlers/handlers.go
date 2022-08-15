package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jacstn/golang-url-shortner/config"
	"github.com/jacstn/golang-url-shortner/pkg/models"
)

var app *config.AppConfig

func NewHandlers(c *config.AppConfig) {
	app = c
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := models.TemplateData{}
	app.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)

	renderTemplate(w, "home", &data)
}

func About(w http.ResponseWriter, r *http.Request) {
	ip := app.Session.GetString(r.Context(), "remote_ip")
	data := models.TemplateData{}
	fmt.Printf(ip)
	data["remote_ip"] = string(ip)

	renderTemplate(w, "about", &data)
}

func renderTemplate(w http.ResponseWriter, templateName string, data *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateName+".tmpl", "./templates/base.layout.tmpl")

	err := parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "This is about page")
	}
}
