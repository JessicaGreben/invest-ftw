package main

import (
	"log"
	"net/http"

	"./handlers"
)

func main() {
	serverStart()
}

func serverStart() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/submit", handlers.SubmitPage)
	http.HandleFunc("/resources", handlers.ResourcesPage)
	http.HandleFunc("/dailystock", handlers.StockPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
