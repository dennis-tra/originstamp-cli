package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TemplateData struct {
	Version string
}

func main() {
	versionFile, err := os.Open("VERSION")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := versionFile.Close(); err != nil {
			panic(err)
		}
	}()

	versionBytes, err := ioutil.ReadAll(versionFile)
	if err != nil {
		panic(err)
	}
	tmpl := template.Must(template.ParseFiles(filepath.Join("pkg", "version.go.tpl")))

	data := TemplateData{
		Version: strings.TrimSpace(string(versionBytes)),
	}
	// open output file
	fo, err := os.Create(filepath.Join("pkg", "version.go"))
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	err = tmpl.Execute(fo, data)
	if err != nil {
		panic(err)
	}
}
