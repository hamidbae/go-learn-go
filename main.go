package main

import (
	engine "final-project/config/gin"
	"final-project/config/postgres"
	userrepo "final-project/pkg/repository/user"
	authhandler "final-project/pkg/server/http/handler/auth"
	authrouter "final-project/pkg/server/http/router/v1"
	authusecase "final-project/pkg/usecase/auth"

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

	authHandler := authhandler.NewAuthHandler(authUsecase)
	
	authrouter.NewAuthRouter(ginEngine, authHandler).Routers()
	
	ginEngine.Serve()
}