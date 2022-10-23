package photo

import "github.com/gin-gonic/gin"

type PhotoHandler interface {
	AddPhotoHdl(ctx *gin.Context)
}