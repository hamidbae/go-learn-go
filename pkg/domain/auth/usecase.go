package auth

import (
	"context"
	"final-project/pkg/domain/response"
	"final-project/pkg/domain/user"
)

type AuthUsecase interface {
	RegisterSvc(ctx context.Context, input user.User) (result user.UserCreatedDto, err response.UsecaseError)
	LoginSvc(ctx context.Context, input Login) (result Token, err response.UsecaseError)
}