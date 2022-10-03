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
	public.POST("/posts",controllers.AddPost)
	public.PUT("/posts/:id",controllers.UpdatePost)
	public.DELETE("/posts/:id",controllers.DeletePost)
	
	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user",controllers.CurrentUser)

	r.Run(":8080")

}