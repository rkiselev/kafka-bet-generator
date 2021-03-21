package model

import "time"

type Bet struct {
	Bettor string `json:"bettor"`
	Match string `json:"match"`
	Amount int `json:"amount"`
	Time time.Time `json:"time"`
}

type Event struct {
	Match string `json:"match"`
	Score Score `json:"score"`
	Time time.Time `json:"time"`
}

type Score struct {
	Home int `json:"home"`
	Away int `json:"away"`
}
