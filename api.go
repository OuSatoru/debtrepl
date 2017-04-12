package main

import "net/http"

// get::  para:: rpno= deleted=0/1 idno= compname=
//        json:: rpno=
// post:: para = get::json
func api(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
	} else if r.Method == "POST" {
		r.ParseForm()
	}
}

func get()  {
	
}

func post()  {
	
}
