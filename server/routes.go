package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jahankohan/go_event_listener/handlers"
)

func NewRouter() *gin.Engine {
	
	gin.SetMode(gin.DebugMode)

	router := gin.New()

	// MiddleWares
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.Default())

	v1 := router.Group("v1") 
	{
		events := v1.Group("events") 
		{
			handlers := handlers.EventHandler{}

			events.GET("/getAllEvents", handlers.GetAllEvents)
		}
	}
	return router
}