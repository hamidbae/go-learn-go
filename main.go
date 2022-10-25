package main

import (
	engine "final-project/config/gin"
	"final-project/config/postgres"
	"final-project/docs"
	"os"

	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"

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

	socialmediarepo "final-project/pkg/repository/socialmedia"
	socialmediahandler "final-project/pkg/server/http/handler/socialmedia"
	socialmediarouter "final-project/pkg/server/http/router/v1"
	socialmediausecase "final-project/pkg/usecase/socialmedia"

	"github.com/gin-gonic/gin"
)

// @title           MyGram API Documentation
// @version         1.0
// @description     This is API for user to post something
// @termsOfService  http://swagger.io/terms/

// @contact.name   Hamid Baehaqi
// @contact.email  hamid1bae1@gmail.com

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description "Type 'Bearer TOKEN' to correctly set the API Key"
func main() {
	// uncomment below when run on local
	// err := godotenv.Load(".env")

	// if err != nil {
	//   log.Fatalf("Error loading .env file")
	// }

	postgresCln := postgres.NewPostgresConnection()
	
	host := os.Getenv("HOST") 
	port := os.Getenv("PORT")
	ginEngine := engine.NewGinHttp(engine.Config{
		Port: ":" + port,
	})

	ginEngine.GetGin().Use(
		gin.Recovery(),
		gin.Logger(),
	)


	docs.SwaggerInfo.Host = host + ":" + port
	docs.SwaggerInfo.BasePath = "/api"
	ginEngine.GetGin().GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
	
	userRepo := userrepo.NewUserRepo(postgresCln)
	photoRepo := photorepo.NewPhotoRepo(postgresCln)
	commentRepo := commentrepo.NewCommentRepo(postgresCln)
	socialMediaRepo := socialmediarepo.NewSocialMediaRepo(postgresCln)
	
	authUsecase := authusecase.NewAuthUsecase(userRepo)
	userUsecase := userusecase.NewUserUsecase(userRepo)
	photoUsecase := photousecase.NewPhotoUsecase(photoRepo, userRepo)
	commentUsecase := commentusecase.NewCommentUsecase(commentRepo, userRepo)
	socialMediaUsecase := socialmediausecase.NewSocialMediaUsecase(socialMediaRepo, userRepo)
	
	authHandler := authhandler.NewAuthHandler(authUsecase)
	userHandler := userhandler.NewUserHandler(userUsecase)
	photoHandler := photohandler.NewPhotoHandler(photoUsecase)
	commentHandler := commenthandler.NewCommentHandler(commentUsecase)
	socialMediaHandler := socialmediahandler.NewSocialMediaHandler(socialMediaUsecase)
	
	userrouter.NewUserRouter(ginEngine, userHandler).Routers()
	authrouter.NewAuthRouter(ginEngine, authHandler).Routers()
	photorouter.NewPhotoRouter(ginEngine, photoHandler).Routers()
	commentrouter.NewCommentRouter(ginEngine, commentHandler).Routers()
	socialmediarouter.NewSocialMediaRouter(ginEngine, socialMediaHandler).Routers()
	
	ginEngine.Serve()
}