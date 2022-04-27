package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

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

	// SSE
	stream := sse.NewServer()
	// sseRoute := r.Group("/sse")
	// sseRoute.Use(HeadersMiddleware())
	// sseRoute.Use(stream.ServeHTTP())
	// sseRoute.GET("/subscribe", sse.StreamHandler)
	// sseRoute.GET("/test", sse.TestMessage)

	r.GET("/stream", func(c *gin.Context) {
		// We are streaming current time to clients in the interval 10 seconds
		go func() {
			for {
				time.Sleep(time.Second * 10)
				now := time.Now().Format("2006-01-02 15:04:05")
				currentTime := fmt.Sprintf("The Current Time Is %v", now)

				// Send current time to clients message channel
				stream.Message <- currentTime
			}
		}()

		c.Stream(func(w io.Writer) bool {
			// Stream message to client from message channel
			if msg, ok := <-stream.Message; ok {
				c.SSEvent("message", msg)
				return true
			}
			return false
		})
	})

	r.GET("/test", func(ctx *gin.Context) {
		stream.Message <- "Test Message"
	})

	r.Use(HeadersMiddleware())
	r.Use(stream.ServeHTTP())
	r.GET("/sse/subscribe", sse.StreamHandler)
	r.GET("/sse/test", sse.TestMessage)

	// Donations
	r.GET("/donations/:id", api.GetDonationById)
	r.POST("/donations/", api.PostDonation)

	// Events
	r.GET("/checkins/:id", api.GetCheckinById)
	r.POST("/checkins/", api.PostCheckin)
	r.GET("/events/:id", api.GetEventById)
	r.GET("/events/:id/attendees", api.GetAttendeesByEventId)
	r.POST("/events/", api.GetEvents)

	r.Run(fmt.Sprintf(":%d", *port))
}

//It keeps a list of clients those are currently attached
//and broadcasting events to those clients.
type Event struct {
	// Events are pushed to this channel by the main events-gathering routine
	Message chan string

	// New client connections
	NewClients chan chan string

	// Closed client connections
	ClosedClients chan chan string

	// Total client connections
	TotalClients map[chan string]bool
}

// New event messages are broadcast to all registered client connection channels
type ClientChan chan string

// Initialize event and Start procnteessing requests
func NewServer() (event *Event) {

	event = &Event{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}

	go event.listen()

	return
}

//It Listens all incoming requests from clients.
//Handles addition and removal of clients and broadcast messages to clients.
func (stream *Event) listen() {
	for {
		select {
		// Add new available client
		case client := <-stream.NewClients:
			stream.TotalClients[client] = true
			log.Printf("Client added. %d registered clients", len(stream.TotalClients))

		// Remove closed client
		case client := <-stream.ClosedClients:
			delete(stream.TotalClients, client)
			log.Printf("Removed client. %d registered clients", len(stream.TotalClients))

		// Broadcast message to client
		case eventMsg := <-stream.Message:
			for clientMessageChan := range stream.TotalClients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

func (stream *Event) serveHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Initialize client channel
		clientChan := make(ClientChan)

		// Send new connection to event server
		stream.NewClients <- clientChan

		defer func() {
			// Send closed connection to event server
			stream.ClosedClients <- clientChan
		}()

		go func() {
			// Send connection that is closed by client to event server
			<-c.Done()
			stream.ClosedClients <- clientChan
		}()

		c.Next()
	}
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
