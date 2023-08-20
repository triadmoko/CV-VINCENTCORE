package helpers

import (
	"bytes"
	"html/template"
)

func ParseTemplate(html string, data any) (string, error) {
	body, err := template.New("template").Parse(html)
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	if err := body.Execute(&tpl, data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}
