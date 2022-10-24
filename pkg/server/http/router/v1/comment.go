package v1

import (
	engine "final-project/config/gin"
	"final-project/pkg/domain/comment"
	"final-project/pkg/server/http/middleware"
	"final-project/pkg/server/http/router"

	"github.com/gin-gonic/gin"
)

type CommentRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	commentHandler comment.CommentHandler
}

func NewCommentRouter(ginEngine engine.HttpServer, commentHandler comment.CommentHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/api/v1/comment")
	return &CommentRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, commentHandler: commentHandler}
}

func (u *CommentRouterImpl) post() {
	u.routerGroup.POST("", middleware.CheckJwtAuth, u.commentHandler.AddCommentHdl)
}

func (u *CommentRouterImpl) put() {
	u.routerGroup.PUT("/:id", middleware.CheckJwtAuth, u.commentHandler.UpdateCommentHdl)
}

func (u *CommentRouterImpl) delete() {
	u.routerGroup.DELETE("/:id", middleware.CheckJwtAuth, u.commentHandler.DeleteCommentByIdHdl)
}

func (u *CommentRouterImpl) Routers() {
	u.post()
	u.put()
	u.delete()
}