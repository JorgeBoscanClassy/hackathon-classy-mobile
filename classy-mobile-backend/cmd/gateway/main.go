package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"classy.org/classymobile/api"
	"classy.org/classymobile/sse"
	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.Int("port", 4000, "specify a port to use http rather than AWS Lambda")

	CreateStartData()
	fmt.Println("Running Gin implementation on http://localhost:" + strconv.Itoa(*port))

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// StatusCheck
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Active")
	})

	// SSE
	r.GET("/sse/:orgId/subscribe", sse.HandleSSEGin())
	r.GET("/sse/test", sse.TestMessage)

	// Donations
	r.GET("/donations/:id", api.GetDonationById)
	r.POST("/donations", api.PostDonation)

	// Events
	r.GET("/checkins/:id", api.GetCheckinById)
	r.POST("/checkins/", api.PostCheckin)
	r.GET("/events/:id", api.GetEventById)
	r.GET("/events/:id/attendees", api.GetAttendeesByEventId)
	r.POST("/events/", api.GetEvents)

	// Home
	r.GET("/home", api.GetHomeData)

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
			Amount:   rand.Int31n(10000),
			Name:     name,
			Email:    strings.ToLower(strings.TrimSpace(name)) + "@classy.org",
			Campaign: orgs[rand.Intn(len(orgs))],
		})
	}
}
