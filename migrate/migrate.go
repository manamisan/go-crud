package main

import (
	"github.com/manami/go-crud/initializers"
	"github.com/manami/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
