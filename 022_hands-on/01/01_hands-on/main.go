package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type BasicStruc struct {
	Title   string
	Message string
}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "basic.gohtml", BasicStruc{"Index", "This is the index"})
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "basic.gohtml", BasicStruc{"Dog", "Dog goes bow wow"})
	if err != nil {
		log.Fatalln(err)
	}
}

func me(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "basic.gohtml", BasicStruc{"Me", "Hello Bennett"})
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
