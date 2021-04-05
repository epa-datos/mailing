package mailing

import (
	"bytes"
	"html/template"
)

func ParseMailTemplate(baseTemplate Template, data interface{}) (string, error) {
	t, err := template.New("template").Parse(baseTemplate.String())
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
