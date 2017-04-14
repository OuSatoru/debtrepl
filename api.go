package main

import (
	"net/http"
	"fmt"
)

type Reply struct {
	ReplyNo       string     `json:"reply_no"`
	ReplyType     string     `json:"reply_type"`
	ReplyYear     string     `json:"reply_year"`
	Borrower      string     `json:"borrower"`
	IDNo          string     `json:"id_no"`
	TotalAmount   int        `json:"total_amount"`
	Head          string     `json:"head"`
	Comprehensive Security   `json:"comprehensive"`
	Interim       []Security `json:"interim"`
}

type Security struct {
	Mortgage  []Mortgage  `json:"mortgage"`
	Pledge    []Pledge    `json:"pledge"`
	Guarantee []Guarantee `json:"guarantee"`
	Credit    []Credit    `json:"credit"`
}

type Mortgage struct {
	Loan
	Collateral string `json:"collateral"`
	Owner      string `json:"owner"`
	Position   string `json:"position"`
	Area       string `json:"area"`
	Value      string `json:"value"`
}

type Pledge struct {
	Loan
	Material string `json:"material"`
	Owner    string `json:"owner"`
	Value    string `json:"value"`
}

type Guarantee struct {
	Loan
	GuarantorKind string `json:"guarantor_kind"`
	Guarantor     string `json:"guarantor"`
	IDNo          string `json:"id_no"`
}

type Credit struct {
	Loan
}

type Loan struct {
	Category     string `json:"category"`
	BeginDate    string `json:"begin_date"`
	EndDate      string `json:"end_date"`
	Span         int    `json:"span"`
	DebtProduct  string `json:"debt_product"`
	DebtKind     string `json:"debt_kind"`
	Warranty     string `json:"warranty"`
	InterestRate string `json:"interest_rate"`
	Limitation   string `json:"limitation"`
	RiskTip      string `json:"risk_tip"`
	Notes        string `json:"notes"`
	Branch       string `json:"branch"`
	Responsible  string `json:"responsible"`
	Operator     string `json:"operator"`
	Time         string `json:"time"`
}

// get::  para:: rpno= deleted=0/1 idno= compname=
//        json:: rpno=
// post:: para = get::json
func api(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Fprintf(w, "age: %v", r.Form["pfbh"])

	}
}

func get() {

}

func post() {

}
