package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.come/salmanfaris22/Skill-map/models"
)

func listuser(ctx *gin.Context) {
	user := models.User{"salman", "as"}
	ctx.JSON(http.StatusOK, []models.User{user})
}
