package main

import (
	"books/initializers"
	"books/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}
func main() {
	initializers.DB.AutoMigrate(&models.Video{})
}
