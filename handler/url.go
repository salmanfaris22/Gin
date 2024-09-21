package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAPI(r *gin.Engine) {
	userGrup := r.Group("/user")

	userGrup.GET("", listuser)
	userGrup.POST("", creatuser)
	userGrup.PATCH("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "This Open OK",
		})
	})
}
