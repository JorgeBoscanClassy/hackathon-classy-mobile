package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var donations map[string]Donation = make(map[string]Donation)
var nextId int = 0

func main() {
	fmt.Println("Running Gin implementation on http://localhost:4000")

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

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
