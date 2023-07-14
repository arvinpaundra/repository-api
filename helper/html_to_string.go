package helper

import (
	"bytes"
	"html/template"
)

func RenderHTMLToString(text string, funcMap template.FuncMap, data map[string]interface{}) (string, error) {
	tmpl, err := template.New("").Funcs(funcMap).Parse(text)

	if err != nil {
		return "", err
	}

	// apply parsed HTML template data and keep the result in a buffer
	var buff bytes.Buffer

	if err := tmpl.Execute(&buff, data); err != nil {
		return "", err
	}

	return buff.String(), nil
}
