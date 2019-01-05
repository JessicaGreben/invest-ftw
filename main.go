package main

import (
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("static/*"))

func main() {
	serverStart()
}

func serverStart() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/submit", submitPage)
	http.HandleFunc("/resources", resourcesPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := templates.ExecuteTemplate(w, "indexPage", nil)
	if err != nil {
		log.Print("template ExecuteTemplate error: ", err)
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
	w.WriteHeader(http.StatusOK)
	err := templates.ExecuteTemplate(w, "submitPage", iv)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
}

func calculate(init string) string {
	principal, _ := strconv.ParseFloat(init, 64)
	frequency := 12.0
	years := 30.0
	rate := 0.07
	tRate := 1 + rate/frequency
	time := frequency * years
	total := principal * math.Pow(tRate, time)
	return strconv.Itoa(int(total))
}

func resourcesPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := templates.ExecuteTemplate(w, "resourcePage", nil)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
}
