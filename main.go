package main

import (
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

func main() {
	serverStart()
}

func serverStart() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/submit", submitPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

type InvestVars struct {
	Init     string
	Invested string
}

func submitPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	i := r.Form.Get("init")
	iv := InvestVars{
		Init:     i,
		Invested: calculate(i),
	}
	t, err := template.ParseFiles("./static/invest.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, iv)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func calculate(init string) string {
	principal, _ := strconv.ParseFloat(init, 64)
	frequency := 12.0
	years := 30.0
	rate := 0.07
	t := 1 + rate/frequency
	y := frequency * years
	total := principal * math.Pow(t, y)
	return strconv.Itoa(int(total))
}
