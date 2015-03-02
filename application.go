package main

import (
	"os"
	"log"
	"net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
)

var base_tmpl = "templates/base.html"

type Page struct {
	Title string
}

func render_template(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(base_tmpl, tmpl)
	t.Execute(w, p)
}

func Welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := &Page{Title: "Welcome"}
	render_template(w, "templates/welcome.html", p)
}

func Projects(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := &Page{Title: "Projects"}
	render_template(w, "templates/projects.html", p)
}

func main() {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.GET("/", Welcome)
	router.GET("/projects", Projects)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(http.ListenAndServe(":" + port, router))
}
