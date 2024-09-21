package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.come/salmanfaris22/Skill-map/models"
)

func listuser(ctx *gin.Context) {
	user1 := models.User{
		"salman",
		"987452894",
	}
	user2 := models.User{
		"jasuim",
		"322894",
	}
	user3 := models.User{
		"rnus",
		"3242323",
	}
	ctx.JSON(http.StatusOK, []models.User{user1, user2, user3})
}

func creatuser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "This Open OK",
	})
}
