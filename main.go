package main

import (
	engine "assignment2/config/gin"
	"assignment2/config/postgres"
	orderrepo "assignment2/pkg/repository/order"
	userrepo "assignment2/pkg/repository/user"
	orderhandler "assignment2/pkg/server/http/handler/order"
	userhandler "assignment2/pkg/server/http/handler/user"
	orderrouter "assignment2/pkg/server/http/router/v1"
	userrouter "assignment2/pkg/server/http/router/v1"
	orderusecase "assignment2/pkg/usecase/order"
	userusecase "assignment2/pkg/usecase/user"

	"github.com/gin-gonic/gin"
)

func main() {
	
	// models.ConnectDataBase()

	// r := gin.Default()
	
	// public := r.Group("/api")
	// public.POST("/register", controllers.Register)
	// public.POST("/login",controllers.Login)
	// public.GET("/posts",controllers.GetAllPost)
	// public.GET("/posts/:id",controllers.GetOnePostById)
	// public.POST("/posts",controllers.AddPost)
	// public.PUT("/posts/:id",controllers.UpdatePost)
	// public.DELETE("/posts/:id",controllers.DeletePost)
	
	// protected := r.Group("/api/admin")
	// protected.Use(middlewares.JwtAuthMiddleware())
	// protected.GET("/user",controllers.CurrentUser)

	// r.Run(":8080")

	// generate postgres config and connect to postgres
	// this postgres client, will be used in repository layer
	postgresCln := postgres.NewPostgresConnection()

	// gin engine
	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":8080",
	})

	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger(),
	)

	
	// generate user repository
	userRepo := userrepo.NewUserRepo(postgresCln)
	orderRepo := orderrepo.NewOrderRepo(postgresCln)
	// initiate use case
	userUsecase := userusecase.NewUserUsecase(userRepo)
	orderUsecase := orderusecase.NewOrderUsecase(orderRepo, userRepo)
	// initiate handler
	useHandler := userhandler.NewUserHandler(userUsecase)
	orderHandler := orderhandler.NewOrderHandler(orderUsecase)
	// initiate router
	userrouter.NewUserRouter(ginEngine, useHandler).Routers()
	orderrouter.NewOrderRouter(ginEngine, orderHandler).Routers()

	// ASSESSMENT
	// buat API
	// - get user
	// sebelum membuat order
	//	- table dengan relasi order -> user (FOREIGN KEY)
	// 			ref:https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-create-table/
	// 	- code base untuk repo, usecase, dll
	// - create order
	// - get order by user

	// running the service
	ginEngine.Serve()
}