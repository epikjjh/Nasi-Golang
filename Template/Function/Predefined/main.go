package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseFiles("index.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", []string{"one", "two", "three"})
	if err != nil {
		log.Fatalln(err)
	}
}
