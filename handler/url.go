package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAPI(r *gin.Engine) {
	userGrup := r.Group("/user")

	userGrup.GET("")
	userGrup.POST("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "This Open OK",
		})
	})
	userGrup.PATCH("", listuser)
}
