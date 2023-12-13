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

func CreateUser(c *gin.Context) {

	return

}

func GetUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
