package main

import (
	DB "./database"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./views/index.gohtml"))
}
func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/startMonitoring", startMonitor)
	http.ListenAndServe(":8080", nil)

}

func root(w http.ResponseWriter, r *http.Request) {
	data := DB.Select("SELECT * FROM log")
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}
func startMonitor(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w, "page.gohtml", r.Form)
}
