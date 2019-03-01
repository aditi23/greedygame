package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type AdObject struct {
	AdID     string  `json:"ad_id"`
	BidPrice float64 `json:"bid_price"`
}

var Ads []AdObject

// Bidding handler
func Bidding(w http.ResponseWriter, r *http.Request) {

	// Input Validation
	adPlacementID := r.URL.Query()["ad_placement_id"][0]
	if adPlacementID == "" {
		http.Error(w, "Invalid Parameters Received", 204)
		return
	}

	adObject := fetchBidding(adPlacementID)

	response, err := json.Marshal(adObject)
	if err != nil || adObject == nil {
		http.Error(w, err.Error(), 204)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
}

// fetchBidding method shuffle ads and return the adobject if adId is matched
func fetchBidding(adPlacementID string) *AdObject {
	shuffleAds()
	adPlacement := Ads[rand.Intn(len(Ads))]
	if adPlacement.AdID == adPlacementID {
		return &adPlacement
	}
	return &AdObject{}
}

// shuffleAds method to shuffle the list of ads to pick the random object every time
func shuffleAds() {
	dest := make([]AdObject, len(Ads))
	perm := rand.Perm(len(Ads))
	for i, v := range perm {
		dest[v] = Ads[i]
	}
}
