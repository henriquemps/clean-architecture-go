package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"project-name-here/api/handler"
	"project-name-here/api/middleware"
)

func Execute() {

	r := gin.Default()

	r.Use(cors.Default())
	r.Use(middleware.HelloWorld)

	users := r.Group("/users")
	{
		users.GET("/", handler.ListUsers)
		users.GET("/:user", handler.ShowUser)
		users.POST("/", handler.CreateUser)
	}

	r.Run()
}
