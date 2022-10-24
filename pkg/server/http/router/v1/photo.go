package v1

import (
	engine "final-project/config/gin"
	"final-project/pkg/domain/photo"
	"final-project/pkg/server/http/middleware"
	"final-project/pkg/server/http/router"

	"github.com/gin-gonic/gin"
)

type PhotoRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	photoHandler photo.PhotoHandler
}

func NewPhotoRouter(ginEngine engine.HttpServer, photoHandler photo.PhotoHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/v1/photo")
	return &PhotoRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, photoHandler: photoHandler}
}

func (u *PhotoRouterImpl) get() {
	u.routerGroup.GET("", middleware.CheckJwtAuth, u.photoHandler.GetPhotoByUserIdHdl)
	u.routerGroup.GET("/:id", middleware.CheckJwtAuth, u.photoHandler.GetPhotoByIdHdl)
}

func (u *PhotoRouterImpl) post() {
	u.routerGroup.POST("", middleware.CheckJwtAuth, u.photoHandler.AddPhotoHdl)
}

func (u *PhotoRouterImpl) put() {
	u.routerGroup.PUT("/:id", middleware.CheckJwtAuth, u.photoHandler.UpdatePhotoHdl)
}

func (u *PhotoRouterImpl) delete() {
	u.routerGroup.DELETE("/:id", middleware.CheckJwtAuth, u.photoHandler.DeletePhotoByIdHdl)
}

func (u *PhotoRouterImpl) Routers() {
	u.get()
	u.post()
	u.put()
	u.delete()
}