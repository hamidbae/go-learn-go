package main

import (
	engine "final-project/config/gin"
	"final-project/config/postgres"
	userrepo "final-project/pkg/repository/user"
	authhandler "final-project/pkg/server/http/handler/auth"
	userhandler "final-project/pkg/server/http/handler/user"
	authrouter "final-project/pkg/server/http/router/v1"
	userrouter "final-project/pkg/server/http/router/v1"
	authusecase "final-project/pkg/usecase/auth"
	userusecase "final-project/pkg/usecase/user"

	"github.com/gin-gonic/gin"
)

func main() {
	postgresCln := postgres.NewPostgresConnection()
	
	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":8080",
	})

	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger(),
	)
	
	userRepo := userrepo.NewUserRepo(postgresCln)
	
	authUsecase := authusecase.NewAuthUsecase(userRepo)
	userUsecase := userusecase.NewUserUsecase(userRepo)
	
	authHandler := authhandler.NewAuthHandler(authUsecase)
	userHandler := userhandler.NewUserHandler(userUsecase)
	
	userrouter.NewUserRouter(ginEngine, userHandler).Routers()
	authrouter.NewAuthRouter(ginEngine, authHandler).Routers()
	
	ginEngine.Serve()
}