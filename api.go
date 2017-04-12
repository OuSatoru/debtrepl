package main

import "net/http"

type Reply struct {
	ReplyNo     string `json:"reply_no"`
	ReplyType   string `json:"reply_type"`
	ReplyYear   string `json:"reply_year"`
	Borrower    string `json:"borrower"`
	IdNo        string `json:"id_no"`
	TotalAmount int    `json:"total_amount"`
	Head        string `json:"head"`
}

type In struct {
	BeginDate   string `json:"begin_date"`
	EndDate     string `json:"end_date"`
	Span        int    `json:"span"`
	DebtProduct string `json:"debt_product"`
	DebtKind    string `json:"debt_kind"`
}

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

func get() {

}

func post() {

}
