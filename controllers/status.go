package controllers

import (
	"assignment3/model"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetStatus(c *gin.Context) {
	var status model.Status

	status.Water = 1 + rand.Intn(100-1)
	status.Wind = 1 + rand.Intn(100-1)

	switch {
	case status.Water <= 5:
		status.WaterStatus = "Aman"
	case (status.Water >= 6) && (status.Water <= 8):
		status.WaterStatus = "Siaga"
	case status.Water > 8:
		status.WaterStatus = "Bahaya"
	default:
		status.WaterStatus = "Something went wrong"
	}

	switch {
	case status.Wind <= 6:
		status.WindStatus = "Aman"
	case (status.Wind >= 7) && (status.Wind <= 15):
		status.WindStatus = "Siaga"
	case status.Wind > 15:
		status.WindStatus = "Bahaya"
	default:
		status.WindStatus = "Something went wrong"
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"status": status,
	})
}
