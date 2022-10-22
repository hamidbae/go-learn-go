package v1

import (
	engine "final-project/config/gin"
	"final-project/pkg/domain/auth"
	"final-project/pkg/server/http/middleware"
	"final-project/pkg/server/http/router"

	"github.com/gin-gonic/gin"
)

type AuthRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	authHandler auth.AuthHandler
}

func NewAuthRouter(ginEngine engine.HttpServer, authHandler auth.AuthHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/v1/auth")
	return &AuthRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, authHandler: authHandler}
}

func (u *AuthRouterImpl) post() {
	u.routerGroup.POST("/register", u.authHandler.RegisterHdl)
	u.routerGroup.POST("/login", u.authHandler.LoginHdl)
	u.routerGroup.POST("/refresh-token", middleware.CheckJwtAuth, u.authHandler.RefreshTokenHdl)
}

func (u *AuthRouterImpl) Routers() {
	// u.get()
	u.post()
}