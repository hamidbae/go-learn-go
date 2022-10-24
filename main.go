package main

import (
	engine "final-project/config/gin"
	"final-project/config/postgres"

	userrepo "final-project/pkg/repository/user"
	userhandler "final-project/pkg/server/http/handler/user"
	userrouter "final-project/pkg/server/http/router/v1"
	userusecase "final-project/pkg/usecase/user"

	authhandler "final-project/pkg/server/http/handler/auth"
	authrouter "final-project/pkg/server/http/router/v1"
	authusecase "final-project/pkg/usecase/auth"

	photorepo "final-project/pkg/repository/photo"
	photohandler "final-project/pkg/server/http/handler/photo"
	photorouter "final-project/pkg/server/http/router/v1"
	photousecase "final-project/pkg/usecase/photo"

	commentrepo "final-project/pkg/repository/comment"
	commenthandler "final-project/pkg/server/http/handler/comment"
	commentrouter "final-project/pkg/server/http/router/v1"
	commentusecase "final-project/pkg/usecase/comment"

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
	photoRepo := photorepo.NewPhotoRepo(postgresCln)
	commentRepo := commentrepo.NewCommentRepo(postgresCln)
	
	authUsecase := authusecase.NewAuthUsecase(userRepo)
	userUsecase := userusecase.NewUserUsecase(userRepo)
	photoUsecase := photousecase.NewPhotoUsecase(photoRepo, userRepo)
	commentUsecase := commentusecase.NewCommentUsecase(commentRepo, userRepo)
	
	authHandler := authhandler.NewAuthHandler(authUsecase)
	userHandler := userhandler.NewUserHandler(userUsecase)
	photoHandler := photohandler.NewPhotoHandler(photoUsecase)
	commentHandler := commenthandler.NewCommentHandler(commentUsecase)
	
	userrouter.NewUserRouter(ginEngine, userHandler).Routers()
	authrouter.NewAuthRouter(ginEngine, authHandler).Routers()
	photorouter.NewPhotoRouter(ginEngine, photoHandler).Routers()
	commentrouter.NewCommentRouter(ginEngine, commentHandler).Routers()
	
	ginEngine.Serve()
}