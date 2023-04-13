package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manami/go-crud/initializers"
	"github.com/manami/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the post
	var post []models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post []models.Post
	initializers.DB.First(&post, id)

	// update a post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Delete a post
	initializers.DB.Delete(&models.Post{}, id)

	// return it
	c.Status(200)
}
