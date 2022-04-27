package api

import (
	"net/http"

	"classy.org/classymobile/sse"
	"classy.org/classymobile/store"
	"github.com/gin-gonic/gin"
)

func CalculateHomeData(donation Donation) {
	store.HomeData.RaisedThisWeek += donation.Amount
	homeDonation := store.HomeDonation{Name: donation.Name, Time: donation.CreatedOn, Campaign: donation.Campaign, Amount: donation.Amount}
	if len(store.HomeData.Donations) >= 3 {
		store.HomeData.Donations = append([]store.HomeDonation{homeDonation}, store.HomeData.Donations[1:]...)
	} else {
		store.HomeData.Donations = append([]store.HomeDonation{homeDonation}, store.HomeData.Donations...)
	}
	sse.SendMessage(store.HomeData)
}

func GetHomeData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, store.HomeData)
}
