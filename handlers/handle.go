package handlers

import (
	"html/template"
	"log"
	"net/http"

	"../invest"
	"../stock"
)

var templates = template.Must(template.ParseGlob("static/*"))

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := templates.ExecuteTemplate(w, "indexPage", nil)
	if err != nil {
		log.Print("template ExecuteTemplate error: ", err)
	}
}

func ResourcesPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := templates.ExecuteTemplate(w, "resourcePage", nil)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
}

func StockPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	p := stock.GetDaily()
	err := templates.ExecuteTemplate(w, "stockPage", p)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
}

func SubmitPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	i := r.Form.Get("init")
	iv := invest.InvestVars{
		Init:     i,
		Invested: invest.Calculate(i),
	}
	w.WriteHeader(http.StatusOK)
	err := templates.ExecuteTemplate(w, "submitPage", iv)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
}
