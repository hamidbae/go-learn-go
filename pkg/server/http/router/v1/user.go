package v1

import (
	engine "assignment2/config/gin"
	"assignment2/pkg/domain/user"
	"assignment2/pkg/server/http/router"

	"github.com/gin-gonic/gin"
)

type UserRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	userHandler user.UserHandler
}

func NewUserRouter(ginEngine engine.HttpServer, userHandler user.UserHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/v1/user")
	return &UserRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler}
}

func (u *UserRouterImpl) get() {
	// all path for get method are here
	u.routerGroup.GET("/:email", u.userHandler.GetUserByEmailHdl)
}

func (u *UserRouterImpl) post() {
	// all path for post method are here
	u.routerGroup.POST("", u.userHandler.InsertUserHdl)
}

func (u *UserRouterImpl) Routers() {
	u.get()
	u.post()
}
