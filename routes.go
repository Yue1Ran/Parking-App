package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/create", func(c *gin.Context) {
		size, _ := strconv.Atoi(c.PostForm("size"))
		NewParkingLot(size)
		c.String(http.StatusOK, "Created parking of %d slots", size)
	})

	r.POST("/park", func(c *gin.Context) {
		reg := c.PostForm("plate")
		msg, _ := Park(reg)
		c.String(http.StatusOK, msg)
	})

	r.POST("/leave", func(c *gin.Context) {
		reg := c.PostForm("plate")
		hours, _ := strconv.Atoi(c.PostForm("hours"))
		msg, _ := Leave(reg, hours)
		c.String(http.StatusOK, msg)
	})

	r.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, Status())
	})

	return r
}
