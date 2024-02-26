package main

import (
	"books/controllers"
	"books/initializers"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {

	http.HandleFunc("/api/v1/videos/create", controllers.VideosCreate)
	http.HandleFunc("/api/v1/videos", controllers.VideosIndex)
	http.HandleFunc("/api/v1/videos/", controllers.VideosShow)
	http.HandleFunc("/api/v1/videos/update/", controllers.VideosUpdate)
	http.HandleFunc("/api/v1/videos/delete/", controllers.VideosDelete)

	http.ListenAndServe(":3000", nil)
}
