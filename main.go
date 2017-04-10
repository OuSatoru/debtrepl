package main

import (
	"net/http"
	"html/template"
)

const replyPage = "web/tmpl/reply.html"

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/", reply)
	http.ListenAndServe(":2333", nil)
}

func reply(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		t, _ := template.ParseFiles(replyPage)
		t.Execute(w, nil)
	}
}
