package main

import (
	"github.com/gin-gonic/gin"
	"github.come/salmanfaris22/Skill-map/handler"
)

func main() {
	server := gin.Default()

	server.POST("/login", handler.Logginehanler)
	server.Run()
}
