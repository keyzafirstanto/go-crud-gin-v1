package controllers

import (
	"go-crud-gin-v1/initializers"
	"go-crud-gin-v1/models"

	"github.com/gin-gonic/gin"
)

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsCreate(c *gin.Context) {
	//Get Data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsShow(c *gin.Context) {
	// Get ID of url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	// respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//  Get the id off the url
	id := c.Param("id")

	// Get the data off the req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	// Find the data
	var post models.Post
	initializers.DB.First(&post, id)

	// update the data
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.Delete(&post, id)

	// respond
	c.Status(200)
}
