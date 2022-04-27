package store

import "time"

type Home struct {
	RaisedThisWeek int32          `json:"raisedThisWeek,omitempty"`
	Donations      []HomeDonation `json:"donations,omitempty"`
}

type HomeHighlight struct {
	Name   string `json:"name,omitempty"`
	Amount int32  `json:"amount,omitempty"`
}

type HomeDonation struct {
	Name     string    `json:"name,omitempty"`
	Time     time.Time `json:"time,omitempty"`
	Campaign string    `json:"campaign,omitempty"`
	Amount   int32     `json:"amount,omitempty"`
}

var HomeData Home = Home{}
