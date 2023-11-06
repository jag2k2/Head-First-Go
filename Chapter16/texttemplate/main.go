package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	template, err := template.New("test").Parse(text)
	check(err)
	err = template.Execute(os.Stdout, data)
	check(err)
}

func main() {
	templateText := "Dot is: {{.}}\n"
	executeTemplate(templateText, "ABC")
	executeTemplate(templateText, 123.5)
}
