package controllers

import (
	"crud/initializers"
	"crud/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	
	c.JSON(200, gin.H{
		"post": post,
	})

	fmt.Println(post.Title)
}

func GetPosts(c *gin.Context) {
	
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostByIndex(c *gin.Context) {
	
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	
	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}