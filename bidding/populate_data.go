package main

import (
	"github.com/aditi23/greedygame/bidding/handlers"
)

func InitAds() {
	handlers.Ads = []handlers.AdObject{
		{
			AdID:     "12345",
			BidPrice: 10.5,
		},
		{
			AdID:     "12345",
			BidPrice: 20.5,
		},
		{
			AdID:     "12345",
			BidPrice: 18.5,
		},
		{
			AdID:     "12345",
			BidPrice: 12.5,
		},
		{
			AdID:     "12345",
			BidPrice: 29.5,
		},
		{
			AdID:     "12345",
			BidPrice: 25.5,
		},
	}
}
