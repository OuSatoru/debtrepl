package main

import (
	"net/http"
	"html/template"
	"fmt"
)

const replyPage = "web/reply.html"

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/", reply)
	http.ListenAndServe(":2333", nil)
}

func reply(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == "GET" {
			t, _ := template.ParseFiles(replyPage)
			t.Execute(w, nil)
		} else {
			r.ParseForm()
			fmt.Println(r.Form)
			fmt.Fprintf(w, "age: %v", r.Form["age"])
		}
	}
}
