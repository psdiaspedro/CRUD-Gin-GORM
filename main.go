package main

import (
	"crud/controllers"
	"crud/initializers"

	"github.com/gin-gonic/gin"
)

func main() {

	initializers.LoadEnv()
	initializers.ConnectDB()

	r := gin.Default()
	r.POST("/post", controllers.CreatePost)
	r.GET("/post", controllers.GetPosts)
	r.GET("/post/:id", controllers.GetPostByIndex)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)
	r.Run() // listen and serve on 0.0.0.0:8080
}