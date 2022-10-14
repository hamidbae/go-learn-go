package order

import "github.com/gin-gonic/gin"

// this handler will use GIN GONIC as http web framework
type OrderHandler interface {
	GetOrderByUserEmailHdl(ctx *gin.Context)
	CreateOrderHdl(ctx *gin.Context)
}
