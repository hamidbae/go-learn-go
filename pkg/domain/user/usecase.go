package user

import (
	"context"
	"final-project/pkg/domain/response"
)

type UserUsecase interface {
	GetUserByIdSvc(ctx context.Context, input uint64) (result UserWithSocialMediaDto, err response.UsecaseError)
	UpdateUserSvc(ctx context.Context, userId uint64, input UserUpdateDto) (result UserUpdateResponseDto, err response.UsecaseError)
	DeleteUserSvc(ctx context.Context, userId uint64) (err response.UsecaseError)
}