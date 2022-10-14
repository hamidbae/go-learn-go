package v1

import (
	engine "assignment2/config/gin"
	"assignment2/pkg/domain/order"
	"assignment2/pkg/server/http/router"

	"github.com/gin-gonic/gin"
)

type OrderRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	orderHandler order.OrderHandler
}

func NewOrderRouter(ginEngine engine.HttpServer, orderHandler order.OrderHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/v1/order")
	return &OrderRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, orderHandler: orderHandler}
}

func (u *OrderRouterImpl) get() {
	// all path for get method are here
	u.routerGroup.GET("/:email", u.orderHandler.GetOrderByUserEmailHdl)
}

func (u *OrderRouterImpl) post() {
	// all path for post method are here
	u.routerGroup.POST("", u.orderHandler.CreateOrderHdl)
}

func (u *OrderRouterImpl) Routers() {
	u.get()
	u.post()
}
