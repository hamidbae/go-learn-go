package photo

import "github.com/gin-gonic/gin"

type PhotoHandler interface {
	AddPhotoHdl(ctx *gin.Context)
	GetPhotoByIdHdl(ctx *gin.Context)
	GetPhotoByUserIdHdl(ctx *gin.Context)
	UpdatePhotoHdl(ctx *gin.Context)
	DeletePhotoByIdHdl(ctx *gin.Context)
}