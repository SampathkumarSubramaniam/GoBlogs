package main

import (
	"os"
	"text/template"
)

type Princess struct {
	Name  string
	Class string
}

func main() {
	//exampleOne()
	parseHtml()
}

func parseHtml() {
	tpl, err := template.ParseFiles("web/test.gohtml")
	if err != nil {
		panic(err)
	}
	fp, err := os.Create("index.html")
	err = tpl.Execute(fp, nil)
	if err != nil {
		panic(err)
	}
}

func exampleOne() {
	nadhira := Princess{Name: "Nadhira", Class: "first"}
	tmpl, err := template.New("test").Parse("{{.Name}} princess is studying in {{.Class}} grade.")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nadhira)
	if err != nil {
		panic(err)
	}
}
