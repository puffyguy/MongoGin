package main

import (
	"MongoGin/configs"
	"MongoGin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"data": "Hello, Welcome",
			})
		})
	}

	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router)
	
	router.Run(":4000")
}
