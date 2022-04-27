package sse

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SSEPayload struct {
	Type           []string    `json:"type,omitempty"`
	RaisedThisWeek float32     `json:"raisedThisWeek,omitempty"`
	Highlights     []Highlight `json:"highlights,omitempty"`
	Donations      []Donations `json:"donations,omitempty"`
	ChartData      string      `json:"chartData,omitempty"`
}

type Highlight struct {
	Name   string  `json:"name,omitempty"`
	Amount float32 `json:"amount,omitempty"`
}

type Donations struct {
	Name     string    `json:"name,omitempty"`
	Time     time.Time `json:"time,omitempty"`
	Campaign string    `json:"campaign,omitempty"`
	Amount   float32   `json:"amount,omitempty"`
}

var ChannelList []chan string

func HandleSSEGin() gin.HandlerFunc {
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
		flusher, _ := w.(http.Flusher)

		// trap the request under loop forever
		for {
			select {

			// message will received here and printed
			case message := <-messageChan:
				fmt.Fprintf(w, "%s\n", message)
				flusher.Flush()

			// connection is closed then defer will be executed
			case <-r.Context().Done():
				return
			}
		}
	}
}

func SendMessage(message SSEPayload) {
	data, _ := json.Marshal(message)
	for _, channel := range ChannelList {
		channel <- string(data)
	}
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
