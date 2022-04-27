package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Donation struct {
	Id        string    `json:"id"`
	Amount    uint64    `json:"amount" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Campaign  string    `json:"campaign" binding:"required"`
	CreatedOn time.Time `json:"createdOn"`
}

var Donations map[string]Donation = make(map[string]Donation)
var nextId int = 0

func GetDonationById(c *gin.Context) {
	id := c.Param("id")
	if donation, ok := Donations[id]; ok {
		c.JSON(http.StatusOK, donation)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "donation id does not exist"})
}

func PostDonation(c *gin.Context) {
	var donation Donation
	if err := c.BindJSON(&donation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	donation = AddDonation(donation)
	c.JSON(http.StatusCreated, donation)
}

func AddDonation(donation Donation) Donation {
	id := strconv.Itoa(nextId)
	nextId += 1

	donation.Id = id
	donation.CreatedOn = time.Now()
	Donations[id] = donation

	go CalculateHomeData(donation)
	return donation
}

func GetDonations(c *gin.Context) {
	var result []Donation

	i := 0
	for _, v := range Donations {
		result = append(result, v)
		i++

		if i == 25 {
			break
		}
	}

	c.JSON(http.StatusOK, result)
}
