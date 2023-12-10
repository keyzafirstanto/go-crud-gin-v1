package main

import (
	"go-crud-gin-v1/initializers"
	"go-crud-gin-v1/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
