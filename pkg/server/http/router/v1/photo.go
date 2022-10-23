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

func (u *PhotoRouterImpl) post() {
	u.routerGroup.POST("", middleware.CheckJwtAuth, u.photoHandler.AddPhotoHdl)
}

func (u *PhotoRouterImpl) Routers() {
	u.post()
}