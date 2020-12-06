package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	owner := user {
		Name: "Allen",
		Age: 25,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", owner)
	if err != nil {
		log.Fatalln(err)
	}
}
