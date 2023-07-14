package generator

import (
	"strings"
	"text/template"
)

type ModelGeneratorOptions struct {
	ClassName string
	FileName  string
}

func GenerateModel(opt ModelGeneratorOptions, templateContent string) (string, error) {
	tmpl, err := template.New("model").Parse(templateContent)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	var buf strings.Builder
	err = tmpl.Execute(&buf, opt)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
