package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"classy.org/classymobile/api"
	"classy.org/classymobile/sse"
	"github.com/gin-gonic/gin"
)

var attendees map[string]Attendee = make(map[string]Attendee)
var nextAttendeeId = 0

var events map[string]Event = make(map[string]Event)
var nextEventId = 0

func main() {
	port := flag.Int("port", 4000, "specify a port to use http rather than AWS Lambda")

	CreateStartData()
	fmt.Println("Running Gin implementation on http://localhost:" + strconv.Itoa(*port))

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// SSE
	stream := sse.NewServer()
	sseRoute := r.Group("/sse")
	sseRoute.Use(HeadersMiddleware())
	sseRoute.Use(stream.ServeHTTP())
	sseRoute.GET("/subscribe", sse.StreamHandler)
	sseRoute.GET("/test", sse.TestMessage)

	r.GET("/donations/:id", api.GetDonationById)
	r.POST("/donations/", api.PostDonation)

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

	r.Run(fmt.Sprintf(":%d", *port))
}

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}

var names []string = []string{"Tammen B", "Patrick C", "Omid B", "Emad B", "Jorge B"}
var orgs []string = []string{"World Central", "Tunnels to Towers"}

func CreateStartData() {
	for i := 0; i < 100; i++ {
		name := names[rand.Intn(len(names))]
		api.AddDonation(api.Donation{
			Amount:   rand.Float32() * 1000,
			Name:     name,
			Email:    strings.ToLower(strings.TrimSpace(name)) + "@classy.org",
			Campaign: orgs[rand.Intn(len(orgs))],
		})
	}
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
