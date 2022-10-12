package main

import (
	"jwt-medium/controllers"
	"jwt-medium/middlewares"
	"jwt-medium/models"

	"github.com/gin-gonic/gin"
)

func main() {
	
	models.ConnectDataBase()

	r := gin.Default()
	
	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login",controllers.Login)
	public.GET("/posts",controllers.GetAllPost)
	public.GET("/posts/:id",controllers.GetOnePostById)
	
	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/profile",controllers.CurrentUser)
	protected.POST("/posts",controllers.AddPost)
	protected.PUT("/posts/:id",controllers.UpdatePost)
	protected.DELETE("/posts/:id",controllers.DeletePost)
	protected.GET("/posts/:id/like",controllers.LikePost)

	r.Run(":8080")

}