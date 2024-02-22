package main

import (
	"books/controllers"
	"books/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	video := r.Group("/api/v1")
	{
		video.GET("/videos", controllers.VideosIndex)
		video.GET("/videos/:id", controllers.VideosShow)
		video.POST("/videos", controllers.VideosCreate)
		video.PUT("/videos/:id", controllers.VideosUpdate)
		video.DELETE("/videos/:id", controllers.VideosDelete)
	}
	r.Run()
}
