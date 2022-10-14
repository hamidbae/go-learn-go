package user

import "github.com/gin-gonic/gin"

// this handler will use GIN GONIC as http web framework
type UserHandler interface {
	GetUserByEmailHdl(ctx *gin.Context)
	InsertUserHdl(ctx *gin.Context)
}
