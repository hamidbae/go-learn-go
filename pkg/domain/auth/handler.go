package auth

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	RegisterHdl(ctx *gin.Context)
	LoginHdl(ctx *gin.Context)
}