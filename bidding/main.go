package main

import (
	"log"
	"net/http"

	"github.com/aditi23/greedygame/bidding/handlers"
)

func main() {

	// populate bidding data
	InitAds()

	// Handler call
	http.HandleFunc("/bid", handlers.Bidding)

	//Run application at provided port
	log.Fatal(http.ListenAndServe(":8888", nil))

}
