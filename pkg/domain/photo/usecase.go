package photo

import (
	"context"
	"final-project/pkg/domain/response"
)

type PhotoUsecase interface {
	AddPhotoSvc(ctx context.Context, input AddPhotoInput) (result PhotoDto, err response.UsecaseError)
	GetPhotoByIdSvc(ctx context.Context, photoId uint64) (result PhotoDetailDto, err response.UsecaseError)
	GetPhotosByUserIdSvc(ctx context.Context, userId uint64) (result []PhotoDetailDto, err response.UsecaseError)
	UpdatePhotoByIdSvc(ctx context.Context, photoId uint64, userId uint64, input UpdatePhotoInput) (result PhotoDto, err response.UsecaseError)
	DeletePhotoByIdSvc(ctx context.Context, userId uint64, photoId uint64) (err response.UsecaseError)
}