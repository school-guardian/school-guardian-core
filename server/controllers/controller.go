package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {

	requestID, _ := c.Get("RequestID")

	c.JSON(http.StatusOK, gin.H{
		"requestID": requestID,
		"status":    "pong",
	})

}
