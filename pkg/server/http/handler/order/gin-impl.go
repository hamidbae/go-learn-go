package order

import (
	"log"
	"net/http"

	"assignment2/pkg/domain/message"
	"assignment2/pkg/domain/order"

	"github.com/gin-gonic/gin"
)

type OrderHdlImpl struct {
	orderUsecase order.OrderUsecase
}

func NewOrderHandler(orderUsecase order.OrderUsecase) order.OrderHandler {
	return &OrderHdlImpl{orderUsecase: orderUsecase}
}

func (u *OrderHdlImpl)  GetOrderByUserEmailHdl(ctx *gin.Context) {
	log.Printf("%T - GetOrderByUserEmailHdl is invoked]\n", u)
	defer log.Printf("%T - GetOrderByUserEmailHdl executed\n", u)

	email := ctx.Params.ByName("email")
	println(email)
	if email == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "param can't be empty string"})
	}

	order, err := u.orderUsecase.GetOrderByUserEmailSvc(ctx, email)

	if err != nil {
		switch err.Error() {
		case "BAD_REQUEST":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  80,
				Error: "invalid processing payload",
			})
			return
		case "INTERNAL_SERVER_ERROR":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Code:  99,
				Error: "something went wrong",
			})
			return
		case "NOT_FOUND":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  99,
				Error: "user not found",
			})
			return
		}
	}
	// response result for the order if success
	ctx.JSONP(http.StatusOK, message.Response{
		Code:    0,
		Message: "success get order",
		Data:    order,
	})
}

func (u *OrderHdlImpl) CreateOrderHdl(ctx *gin.Context) {
	log.Printf("%T - CreateOrderHdl is invoked]\n", u)
	defer log.Printf("%T - CreateOrderHdl executed\n", u)

	// binding / mendapatkan body payload dari request
	log.Println("binding body payload from request")
	var order order.Order
	if err := ctx.ShouldBind(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Code:  80,
			Error: "failed to bind payload",
		})
		return
	}

	// call service/usecase untuk menginsert data
	log.Println("calling insert service usecase")
	result, err := u.orderUsecase.CreateOrderSvc(ctx, order)
	if err != nil {
		switch err.Error() {
		case "BAD_REQUEST":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  80,
				Error: "invalid processing payload",
			})
			return
		case "INTERNAL_SERVER_ERROR":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Code:  99,
				Error: "something went wrong",
			})
			return
		case "NOT_FOUND":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Code:  99,
				Error: "user not found",
			})
			return
		}
	}
	// response result for the order if success
	ctx.JSONP(http.StatusOK, message.Response{
		Code:    0,
		Message: "success insert order",
		Data:    result,
	})
}
