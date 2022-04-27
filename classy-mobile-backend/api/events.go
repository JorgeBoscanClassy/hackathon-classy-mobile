package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var attendees map[string]Attendee = make(map[string]Attendee)
var nextAttendeeId = 0

var events map[string]Event = make(map[string]Event)
var nextEventId = 0

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

func GetCheckinById(c *gin.Context) {
	id := c.Param("id")
	if attendee, ok := attendees[id]; ok {
		c.JSON(http.StatusOK, attendee)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "attendee id does not exist"})
}

func PostCheckin(c *gin.Context) {
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
}

func GetEventById(c *gin.Context) {
	id := c.Param("id")
	if event, ok := events[id]; ok {
		c.JSON(http.StatusOK, event)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "event id does not exist"})
}

func GetAttendeesByEventId(c *gin.Context) {
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
}

func GetEvents(c *gin.Context) {
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
}
