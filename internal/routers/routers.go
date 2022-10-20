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

	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", api.NewUserFormApi().Create)
		userGroup.POST("/query", api.NewUserFormApi().Query)
	}

	adminGroup := router.Group("/admin")
	{ // TODO: admin api
		adminGroup.POST("/put/done", api.NewAdminFormApi().Done)
		adminGroup.GET("/get/todo", api.NewAdminFormApi().GetTodo)
		adminGroup.GET("/get/history", api.NewAdminFormApi().GetHistory)
	}

	return router
}
