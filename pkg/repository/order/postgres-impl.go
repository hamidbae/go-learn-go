package order

import (
	"context"
	"log"

	"assignment2/config/postgres"
	"assignment2/pkg/domain/order"
)

type OrderRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewOrderRepo(pgCln postgres.PostgresClient) order.OrderRepo {
	return &OrderRepoImpl{pgCln: pgCln}
}

func (u *OrderRepoImpl) GetOrderByUserId(ctx context.Context, user_id uint64) (result []order.Order, err error) {
	log.Printf("%T - GetOrderByUserId is invoked]\n", u)
	defer log.Printf("%T - GetOrderByUserId executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()
	// insert new order
	db.Model(&order.Order{}).Where("user_id = ?", user_id).Find(&result)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when getting order with user_id = %v\n", user_id)
	}
	return result, err
}

func (u *OrderRepoImpl) CreateOrder(ctx context.Context, insertedOrder *order.Order) (err error) {
	log.Printf("%T - CreateOrder is invoked]\n", u)
	defer log.Printf("%T - CreateOrder executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()
	// insert new order
	db.Model(&order.Order{}).Create(&insertedOrder)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when creating order with email\n")
	}
	return err
}
