package main

import (
	"log"
	"net/http"
	"time"

	"github.com/aditi23/greedygame/auction/handlers"
)

func main() {

	// Handler call
	http.HandleFunc("/auction", handlers.Auction)

	// server object with address and write timeout
	s := &http.Server{
		Addr:         ":9999",
		WriteTimeout: 200 * time.Millisecond,
	}

	log.Fatal(s.ListenAndServe())

}
