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

type Part struct {
	Name  string
	Count int
}

func main() {
	templateText := "Dot is: {{.}}\n"
	executeTemplate(templateText, "ABC")
	executeTemplate(templateText, 123.5)

	templateText = "start {{if .}}Dot is true!{{end}} finish\n"
	executeTemplate(templateText, true)
	executeTemplate(templateText, false)

	templateText = "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n" //first and last do contain slice.  middle dot contains element
	executeTemplate(templateText, []string{"do", "re", "mi"})

	templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	executeTemplate(templateText, []float64{1.25, 0.99, 27})

	templateText = "Name: {{.Name}}\tCount: {{.Count}}\n"
	executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
	executeTemplate(templateText, Part{Name: "Cables", Count: 2})
}
