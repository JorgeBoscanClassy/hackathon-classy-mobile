package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"classy.org/classymobile/sse"
	"github.com/gin-gonic/gin"
)

var Donations map[string]Donation = make(map[string]Donation)
var nextId int = 0

var attendees map[string]Attendee = make(map[string]Attendee)
var nextAttendeeId = 0

func main() {
	CreateStartData()
	fmt.Println("Running Gin implementation on http://localhost:4000")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// SSE
	r.GET("/sse/subscribe", sse.HandleSSEGin())
	r.GET("/sse/test", func(ctx *gin.Context) {
		testData := sse.SSEPayload{
			Type:           []string{"highlights", "test", "donations", "raised-this-week"},
			RaisedThisWeek: rand.Float32() * 10000000,
			Highlights: []sse.Highlight{
				{"Average Transaction Site", rand.Float32() * 10000000},
				{"Total Transactions", rand.Float32() * 10000000},
			},
			Donations: []sse.Donations{
				{"Omid Borijan", time.Now(), "WorldCentral", rand.Float32() * 10000000},
				{"Tammen K", time.Now(), "Tunnels to Towers", rand.Float32() * 10000000},
				{"Emad B", time.Now(), "Tunnels to Towers", rand.Float32() * 10000000},
			},
			ChartData: "Stub",
		}

		sse.SendMessage(testData)
	})

	r.GET("/donations/:id", func(c *gin.Context) {
		id := c.Param("id")
		if donation, ok := Donations[id]; ok {
			c.JSON(http.StatusOK, donation)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "donation id does not exist"})
	})

	r.POST("/donations/", func(c *gin.Context) {
		var donation Donation
		if err := c.BindJSON(&donation); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		donation = AddDonation(donation)
		c.JSON(http.StatusCreated, donation)
	})

	r.GET("/checkins/:id", func(c *gin.Context) {
		id := c.Param("id")
		if attendee, ok := attendees[id]; ok {
			c.JSON(http.StatusOK, attendee)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "attendee id does not exist"})
	})

	r.POST("/checkins/", func(c *gin.Context) {
		var attendee Attendee
		if err := c.BindJSON(&attendee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := strconv.Itoa(nextAttendeeId)
		nextAttendeeId += 1
		attendee.Id = id
		attendee.CreatedOn = time.Now()
		attendees[id] = attendee

		c.JSON(http.StatusCreated, attendee)
	})

	r.Run(":4000")
}

func AddDonation(donation Donation) Donation {
	id := strconv.Itoa(nextId)
	nextId += 1

	donation.Id = id
	donation.CreatedOn = time.Now()
	Donations[id] = donation

	return donation
}

type Donation struct {
	Id           string    `json:"id"`
	Amount       float32   `json:"amount" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	Organization string    `json:"organization" binding:"required"`
	CreatedOn    time.Time `json:"createdOn"`
}

var names []string = []string{"Tammen B", "Patrick C", "Omid B", "Emad B", "Jorge B"}
var orgs []string = []string{"World Central", "Tunnels to Towers"}

func CreateStartData() {
	for i := 0; i < 100; i++ {
		name := names[rand.Intn(len(names))]
		AddDonation(Donation{
			Amount:       rand.Float32() * 10000000,
			Name:         name,
			Email:        strings.ToLower(strings.TrimSpace(name)) + "@classy.org",
			Organization: orgs[rand.Intn(len(orgs))],
		})
	}
}

type Attendee struct {
	Id        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedOn time.Time `json:"createdOn"`
}
