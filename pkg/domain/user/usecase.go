package user

import (
	"context"
	"final-project/pkg/domain/response"
)

type UserUsecase interface {
	GetUserByIdSvc(ctx context.Context, input uint64) (result UserGetDto, err response.UsecaseError)
	UpdateUserSvc(ctx context.Context, userId uint64, input UserUpdateDto) (result string, err response.UsecaseError)
	DeleteUserSvc(ctx context.Context, userId uint64) (err response.UsecaseError)
}