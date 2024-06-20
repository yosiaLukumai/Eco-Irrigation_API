package utils

import (
	"bytes"
	"html/template"
)

type VerificationEmailDataTemplate struct {
	AppName    string
	VerifyLink string
	Name       string
	Year       int
}

func ParseHtmlVariables(htmlTemplate string, data any) (string, error) {
	tmpl, err := template.New("emailTemplate").Parse(htmlTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
