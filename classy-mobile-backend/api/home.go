package api

import "time"

type Home struct {
	RaisedThisWeek uint64         `json:"raisedThisWeek,omitempty"`
	Donations      []HomeDonation `json:"donations,omitempty"`
}

type HomeHighlight struct {
	Name   string `json:"name,omitempty"`
	Amount uint64 `json:"amount,omitempty"`
}

type HomeDonation struct {
	Name     string    `json:"name,omitempty"`
	Time     time.Time `json:"time,omitempty"`
	Campaign string    `json:"campaign,omitempty"`
	Amount   uint64    `json:"amount,omitempty"`
}

var homeData Home = Home{}

func CalculateHomeData(donation Donation) {
	homeData.RaisedThisWeek += donation.Amount
	homeDonation := HomeDonation{Name: donation.Name, Time: donation.CreatedOn, Campaign: donation.Campaign, Amount: donation.Amount}
	if len(homeData.Donations) >= 3 {
		homeData.Donations = append(homeData.Donations[1:], homeDonation)
	} else {
		homeData.Donations = append(homeData.Donations, homeDonation)
	}
}
