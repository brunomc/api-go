package router

import (
	"github.com/brunomc/api-go/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("api/v1")
	{
		v1.GET("openings", handler.CreateOpeningHandler)
		v1.GET("opening", handler.ShowOpeningHandler)
		v1.POST("opening", handler.CreateOpeningHandler)
		v1.DELETE("opening", handler.DeleteOpeningHandler)
		v1.PUT("opening", handler.UpdateOpeningHandler)
	}

}
