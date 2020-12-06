package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	// In case of map, use key: value
	names := []string{"Allen", "James"}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", names)
	if err != nil {
		log.Fatalln(err)
	}
}
