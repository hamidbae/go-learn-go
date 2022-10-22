package user

import "context"

type UserRepo interface {
	GetUserByEmail(ctx context.Context, email string) (user User, err error)
	GetUserByUsername(ctx context.Context, username string) (user User, err error)
	GetUserById(ctx context.Context, id uint64) (user User, err error)
	InsertUser(ctx context.Context, user *User) (err error)
}