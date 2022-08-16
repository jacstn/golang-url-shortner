package models

import "github.com/jacstn/golang-url-shortner/pkg/forms"

type TemplateData struct {
	Data map[string]interface{}
	Form *forms.Form
}
