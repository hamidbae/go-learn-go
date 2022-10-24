package v1

import (
	engine "final-project/config/gin"
	"final-project/pkg/domain/socialmedia"
	"final-project/pkg/server/http/middleware"
	"final-project/pkg/server/http/router"

	"github.com/gin-gonic/gin"
)

type SocialMediaRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	socialMediaHandler socialmedia.SocialMediaHandler
}

func NewSocialMediaRouter(ginEngine engine.HttpServer, socialMediaHandler socialmedia.SocialMediaHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/v1/social-media")
	return &SocialMediaRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, socialMediaHandler: socialMediaHandler}
}

func (u *SocialMediaRouterImpl) get() {
	u.routerGroup.GET("", middleware.CheckJwtAuth, u.socialMediaHandler.GetSocialMediaByUserIdHdl)
}

func (u *SocialMediaRouterImpl) post() {
	u.routerGroup.POST("", middleware.CheckJwtAuth, u.socialMediaHandler.AddSocialMediaHdl)
}

func (u *SocialMediaRouterImpl) put() {
	u.routerGroup.PUT("/:id", middleware.CheckJwtAuth, u.socialMediaHandler.UpdateSocialMediaHdl)
}

func (u *SocialMediaRouterImpl) delete() {
	u.routerGroup.DELETE("/:id", middleware.CheckJwtAuth, u.socialMediaHandler.DeleteSocialMediaByIdHdl)
}

func (u *SocialMediaRouterImpl) Routers() {
	u.get()
	u.post()
	u.put()
	u.delete()
}