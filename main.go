package main

import (
	"html/template"
	"io"
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
	http.HandleFunc("/stopMonitoring", stopMonitor)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	data := selectQ("SELECT * FROM log")
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}
func startMonitor(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	ip := r.FormValue("serverIp")
	go startMonitoring(ip)

	io.WriteString(w, "Monitoring Started...")
}
func stopMonitor(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	ip := r.FormValue("serverIp")
	go stopMonitoring(ip)

	io.WriteString(w, "Monitoring Stop...")
}
