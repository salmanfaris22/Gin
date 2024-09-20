package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("salman$ go get -u github.com/gin-gonic/gin")
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Skill Map openig",
		})
	})
	r.Run()
}
