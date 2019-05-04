package server

import (
	"github.com/gin-gonic/gin"
	"github.com/naughtydevelopment/todo-go/controllers"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")

	{
		todos := v1.Group("todos/")
		todos.GET("/", func(context *gin.Context) {
			context.String(http.StatusOK, "todos")
		})
	}

	return router
}
