package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

func renderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateName+".tmpl", "./templates/base.layout.tmpl")

	data := make(map[string]string)
	data["variable"] = "this is variable passed"
	err := parsedTemplate.Execute(w, &data)
	if err != nil {
		fmt.Fprintf(w, "This is about page")
	}
}
