package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"classy.org/classymobile/sse"
	"github.com/gin-gonic/gin"
)

var donations map[string]Donation = make(map[string]Donation)
var nextId int = 0

var attendees map[string]Attendee = make(map[string]Attendee)
var nextAttendeeId = 0

var events map[string]Event = make(map[string]Event)
var nextEventId = 0

func main() {
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
		if donation, ok := donations[id]; ok {
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
		id := strconv.Itoa(nextId)
		nextId += 1

		donation.Id = id
		donation.CreatedOn = time.Now()
		donations[id] = donation
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

	r.GET("/events/:id", func(c *gin.Context) {
		id := c.Param("id")
		if event, ok := events[id]; ok {
			c.JSON(http.StatusOK, event)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "event id does not exist"})
	})

	r.GET("/events/:id/attendees", func(c *gin.Context) {
		id := c.Param("id")
		attendants := []Attendee{}

		if event, ok := events[id]; ok {
			for _, attendee := range attendees {
				if attendee.EventId == event.Id {
					attendants = append(attendants, attendee)
				}
			}

			c.JSON(http.StatusOK, attendants)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "event id does not exist"})
	})

	r.POST("/events/", func(c *gin.Context) {
		var event Event
		if err := c.BindJSON(&event); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := strconv.Itoa(nextEventId)
		nextEventId += 1
		event.Id = id
		event.CreatedOn = time.Now()
		events[id] = event

		c.JSON(http.StatusCreated, event)
	})

	r.Run(":4000")
}

type Donation struct {
	Id           string    `json:"id"`
	Amount       string    `json:"amount" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	Organization string    `json:"organization" binding:"required"`
	CreatedOn    time.Time `json:"createdOn"`
}

type Attendee struct {
	Id        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	EventId   string    `json:"eventId" binding:"required"`
	CreatedOn time.Time `json:"createdOn"`
}

type Event struct {
	Id        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedOn time.Time `json:"createdOn"`
}
