package main

import (
	"fmt"
	"math/rand"
	"time"

	"classy.org/classymobile/sse"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Running Gin implementation on http://localhost:4000")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

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

	r.Run(":4000")
}
