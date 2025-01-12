package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pandakn/sa-65-example/controller"
	"github.com/pandakn/sa-65-example/entity"
	"github.com/pandakn/sa-65-example/middlewares"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	entity.SetupDatabase()
	r := gin.Default()

	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			// Room Routes
			r.GET("/rooms", controller.ListRooms)
			r.GET("/room/:id", controller.GetRoom)
			r.POST("/rooms", controller.CreateRoom)

			// RoomType Routes
			r.GET("/room-type/:id", controller.GetRoomType)
			r.GET("/room-types", controller.ListRoomTypes)

			// RoomZone Routes
			r.GET("/room-zone/:id", controller.GetRoomZone)
			r.GET("/room-zones", controller.ListRoomZones)

			// Admin Routes
			r.GET("/admin/:id", controller.GetAdmin)
			r.GET("/admins", controller.ListAdmins)

		}
	}

	// login User Route
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}
