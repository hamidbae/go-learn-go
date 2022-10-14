package user

import "context"

type UserUsecase interface {
	GetUserByEmailSvc(ctx context.Context, email string) (result User, err error)
	InsertUserSvc(ctx context.Context, input User) (result User, err error)
}