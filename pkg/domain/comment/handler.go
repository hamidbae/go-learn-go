package comment

import "github.com/gin-gonic/gin"

type CommentHandler interface {
	AddCommentHdl(ctx *gin.Context)
	UpdateCommentHdl(ctx *gin.Context)
	DeleteCommentByIdHdl(ctx *gin.Context)
}