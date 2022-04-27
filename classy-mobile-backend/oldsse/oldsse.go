package oldsse

import (
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

var MessageList []chan string

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

		// instantiate the channel
		var messageChan chan string
		messageChan = make(chan string)
		MessageList = append(MessageList, messageChan)

		// close the channel after exit the function
		defer func() {
			close(messageChan)
			messageChan = nil
			log.Printf("client connection is closed")
		}()

		// trap the request under loop forever
		for {

			select {

			// message will received here and printed
			case message := <-messageChan:

				ctx.Stream(func(w io.Writer) bool {
					// Stream message to client from message channel
					ctx.SSEvent("message", message)
					return true
				})

			// connection is closed then defer will be executed
			case <-r.Context().Done():
				return
			}
		}
	}
}

func SendMessageGin(message string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, messageChan := range MessageList {
			// send the message through the available channel
			messageChan <- message
		}
	}
}