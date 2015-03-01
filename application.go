package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, _ := template.ParseFiles(tmpl)
	fmt.Println(t)
	t.Execute(w, nil)
}

func Welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	renderTemplate(w, "templates/welcome.html")
}

func Static(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	router := httprouter.New()
	router.GET("/", Welcome)
	router.GET("/static", Static)

	log.Fatal(http.ListenAndServe(":5000", router))
}
