package models

import "github.com/jacstn/golang-url-shortner/pkg/forms"

type TemplateData struct {
	Data map[string]string
	Form *forms.Form
}
