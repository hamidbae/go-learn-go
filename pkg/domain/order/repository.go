package order

import "context"

type OrderRepo interface {
	GetOrderByUserId(ctx context.Context, user_id uint64) (orders []Order, err error)
	CreateOrder(ctx context.Context, order *Order) (err error)
}