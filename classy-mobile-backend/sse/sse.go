package sse

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"time"

	"classy.org/classymobile/api"
	"github.com/gin-gonic/gin"
)

type ClientChan chan string

var ChannelList []chan string
var Stream *Event

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

// Initialize event and Start processing requests
func NewServer() *Event {

	Stream = &Event{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}

	go Stream.listen()

	return Stream
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
			i := 1
			log.Printf("Sending Message")
			for clientMessageChan := range stream.TotalClients {
				log.Printf("Sending Message to Client: %d", i)
				clientMessageChan <- eventMsg
			}
		}
	}
}

func (stream *Event) ServeHTTP() gin.HandlerFunc {
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

func StreamHandler(c *gin.Context) {
	c.Stream(func(w io.Writer) bool {
		// Stream message to client from message channel
		if msg, ok := <-Stream.Message; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}

func SendMessage(message interface{}, ctx *gin.Context) {
	data, _ := json.Marshal(message)
	Stream.Message <- string(data)
}

func TestMessage(ctx *gin.Context) {
	testData := api.Home{
		RaisedThisWeek: rand.Int31n(10000),
		Donations: []api.HomeDonation{
			{"Omid Borijan", time.Now(), "WorldCentral", rand.Int31n(10000)},
			{"Tammen K", time.Now(), "Tunnels to Towers", rand.Int31n(10000)},
			{"Emad B", time.Now(), "Tunnels to Towers", rand.Int31n(10000)},
		},
	}

	SendMessage(testData, ctx)
}

func OldSse(c *gin.Context) {

}
