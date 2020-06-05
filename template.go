package sdk

import (
	"bytes"
	"text/template"
)

// process applies the data structure 'vars' onto an already
// parsed template 't', and returns the resulting string.
func process(t *template.Template, vars interface{}) (string, error) {
	var tmplBytes bytes.Buffer

	err := t.Execute(&tmplBytes, vars)
	if err != nil {
		return "", err
	}
	return tmplBytes.String(), nil
}

// ProcessString parses the supplied string and compiles it using
// the given variables.
func ProcessString(str string, vars interface{}) (string, error) {
	tmpl, err := template.New("tmpl").Parse(str)

	if err != nil {
		return "", err
	}
	return process(tmpl, vars)
}

// ProcessFile parses the supplied filename and compiles its template
// using the given variables.
func ProcessFile(fileName string, vars interface{}) (string, error) {
	tmpl, err := template.ParseFiles(fileName)

	if err != nil {
		return "", err
	}
	return process(tmpl, vars)
}