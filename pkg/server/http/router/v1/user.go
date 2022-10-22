package v1

import (
	engine "final-project/config/gin"
	"final-project/pkg/domain/user"
	"final-project/pkg/server/http/middleware"
	"final-project/pkg/server/http/router"

	"github.com/gin-gonic/gin"
)

type UserRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	userHandler user.UserHandler
}

func NewUserRouter(ginEngine engine.HttpServer, userHandler user.UserHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/v1/user")
	return &UserRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler}
}

func (u *UserRouterImpl) get() {
	u.routerGroup.GET("/:id", u.userHandler.GetUserByIdHdl)
}

func (u *UserRouterImpl) put() {
	u.routerGroup.PUT("", middleware.CheckJwtAuth, u.userHandler.UpdateUserHdl)
}

func (u *UserRouterImpl) delete() {
	u.routerGroup.DELETE("", middleware.CheckJwtAuth, u.userHandler.DeleteUserHdl)
}

func (u *UserRouterImpl) Routers() {
	u.get()
	u.put()
	u.delete()
}