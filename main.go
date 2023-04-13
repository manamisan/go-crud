package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manami/go-crud/controllers"
	"github.com/manami/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.GET("/", controllers.PostsCreate)

	r.Run()
}
