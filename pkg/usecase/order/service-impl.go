package order

import (
	"context"
	"errors"
	"log"

	"assignment2/pkg/domain/order"
	"assignment2/pkg/domain/user"
)

type OrderUsecaseImpl struct {
	orderRepo order.OrderRepo
	userRepo user.UserRepo
}

func NewOrderUsecase(orderRepo order.OrderRepo, userRepo user.UserRepo) order.OrderUsecase {
	return &OrderUsecaseImpl{orderRepo: orderRepo, userRepo: userRepo}
}
	// GetOrderByUserEmailSvc(ctx context.Context, email string) (result Order, err error)
	// CreateOrderSvc(ctx context.Context, input Order) (result Order, err error)

func (u *OrderUsecaseImpl) GetOrderByUserEmailSvc(ctx context.Context, email string) (result []order.Order, err error) {
	log.Printf("%T - GetOrderByUserEmailSvc is invoked]\n", u)
	defer log.Printf("%T - GetOrderByUserEmailSvc executed\n", u)
	// get order from repository (database)
	log.Println("getting user from order repository")
	userFound, err := u.userRepo.GetUserByEmail(ctx, email)
	if userFound.ID <= 0 {
		// kalau tidak berarti user not found
		log.Println("user is not found")
		err = errors.New("NOT_FOUND")
		return result, err
	}

	result, err = u.orderRepo.GetOrderByUserId(ctx, userFound.ID)
	if err != nil {
		// ini berarti ada yang salah dengan connection di database
		log.Println("error when fetching data from database: " + err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
		return result, err
	}
	
	return result, err
}

func (u *OrderUsecaseImpl) CreateOrderSvc(ctx context.Context, input order.Order) (result order.Order, err error) {
	log.Printf("%T - CreateOrderSvc is invoked]\n", u)
	defer log.Printf("%T - CreateOrderSvc executed\n", u)

	user, err := u.userRepo.GetUserById(ctx, input.UserId)
	if user.ID <= 0 {
		// kalau tidak berarti user not found
		log.Println("user is not found")
		err = errors.New("NOT_FOUND")
		return result, err
	}
	
	// valid condition: NOT_FOUND
	log.Println("create order to database process")
	if err = u.orderRepo.CreateOrder(ctx, &input); err != nil {
		log.Printf("error when inserting order:%v\n", err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
	}
	return input, err
}
