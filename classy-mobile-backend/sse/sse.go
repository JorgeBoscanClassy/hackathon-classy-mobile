package sse

import (
	"encoding/json"
	"fmt"
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
			for clientMessageChan := range stream.TotalClients {
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

func HandleSSE() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		w := ctx.Writer
		r := ctx.Request
		log.Printf("Get handshake from client")
		// prepare the header
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		messageChan := make(chan string)
		ChannelList = append(ChannelList, messageChan)

		// close the channel after exit the function
		defer func() {
			removeChannel(messageChan)
			close(messageChan)
			log.Printf("client connection is closed")
		}()

		// prepare the flusher
		// flusher, _ := w.(http.Flusher)

		// trap the request under loop forever
		for {
			select {

			// message will received here and printed
			case message := <-messageChan:
				ctx.SSEvent("message", message)

			// connection is closed then defer will be executed
			case <-r.Context().Done():
				return
			}
		}
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

func removeChannel(channel chan string) {
	index := -1
	for i, tmpChannel := range ChannelList {
		if tmpChannel == channel {
			index = i
			break
		}
	}

	if index >= 0 {
		ChannelList = append(ChannelList[:index], ChannelList[index+1:]...)
	}
	fmt.Println("Done")
}

func TestMessage(ctx *gin.Context) {
	testData := api.Home{
		RaisedThisWeek: rand.Float32() * 10000000,
		Donations: []api.HomeDonation{
			{"Omid Borijan", time.Now(), "WorldCentral", rand.Float32() * 10000000},
			{"Tammen K", time.Now(), "Tunnels to Towers", rand.Float32() * 10000000},
			{"Emad B", time.Now(), "Tunnels to Towers", rand.Float32() * 10000000},
		},
	}

	SendMessage(testData, ctx)
}
