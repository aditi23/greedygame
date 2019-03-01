package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Response struct {
	AdID     string  `json:"ad_id"`
	BidPrice float64 `json:"bid_price"`
}

// Auction handler
func Auction(w http.ResponseWriter, r *http.Request) {

	var (
		bidAmount float64
		response  Response
	)
	resp := &http.Response{}
	// Input Validation
	adPlacementID := r.URL.Query()["ad_placement_id"][0]
	if adPlacementID == "" {
		http.Error(w, "Invalid Parameters Received", 204)
		return
	}

	// channel of http response
	ch := make(chan *http.Response)

	// go routines for the concurrent calls for making bid
	for range [5]int{} {
		go makeBid(ch, adPlacementID)
	}

	for range [5]int{} {
		resp = <-ch
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &response)
		if err == nil || resp.StatusCode == 200 {
			if response.BidPrice > bidAmount {
				bidAmount = response.BidPrice
			}
		}
	}

	if bidAmount == 0 {
		http.Error(w, "", 204)
		return
	}

	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, fmt.Sprintf("%.2f", bidAmount))
}

// makeBid method makes the http request to place the bid for corresponding ad id
func makeBid(ch chan<- *http.Response, adPlacementID string) {
	url := fmt.Sprintf("http://localhost:8811/bid?ad_placement_id=%s", adPlacementID)
	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{
		Timeout: time.Duration(200) * time.Millisecond,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[Auction] Error occured while calling bid API", err)
	}
	ch <- resp
}
