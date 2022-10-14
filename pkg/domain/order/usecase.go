package order

import "context"

type OrderUsecase interface {
	GetOrderByUserEmailSvc(ctx context.Context, email string) (result []Order, err error)
	CreateOrderSvc(ctx context.Context, input Order) (result Order, err error)
}