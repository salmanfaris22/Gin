package handler

import "github.com/gin-gonic/gin"

func Logginehanler(ctx *gin.Context) {

	var json struct {
		Username string
		Password string
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"status": "http://localhost:8080/login",
	})
}
