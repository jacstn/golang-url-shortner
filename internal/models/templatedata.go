package models

import "github.com/jacstn/golang-url-shortner/internal/forms"

type TemplateData struct {
	Data map[string]interface{}
	Form *forms.Form
}
