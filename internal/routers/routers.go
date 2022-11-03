package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/isther/clinic/internal/routers/api"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.MaxMultipartMemory = 8 << 20
	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", api.NewUserFormApi().Create)
		userGroup.POST("/query", api.NewUserFormApi().Query)
		userGroup.POST("/upload", api.NewUserFormApi().Upload)
	}

	adminGroup := router.Group("/admin")
	{ // admin api
		adminGroup.POST("/put/status", api.NewAdminFormApi().Status)
		adminGroup.GET("/get/unreviewed", api.NewAdminFormApi().GetUnreviewed)
		adminGroup.GET("/get/reviewed", api.NewAdminFormApi().GetReviewed)
		adminGroup.GET("/get/cannot", api.NewAdminFormApi().GetCannot)
		adminGroup.GET("/get/done", api.NewAdminFormApi().GetDone)
	}

	return router
}
