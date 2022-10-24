package socialmedia

import "github.com/gin-gonic/gin"

type SocialMediaHandler interface {
	AddSocialMediaHdl(ctx *gin.Context)
	// GetSocialMediaByIdHdl(ctx *gin.Context)
	GetSocialMediaByUserIdHdl(ctx *gin.Context)
	UpdateSocialMediaHdl(ctx *gin.Context)
	DeleteSocialMediaByIdHdl(ctx *gin.Context)
}