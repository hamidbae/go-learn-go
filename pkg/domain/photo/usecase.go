package photo

import (
	"context"
	"final-project/pkg/domain/response"
)

type PhotoUsecase interface {
	AddPhotoSvc(ctx context.Context, input AddPhotoInput) (result Photo, err response.UsecaseError)
	// GetPhotosByUserIdSvc(ctx context.Context, userId uint64) (result []Photo, err response.UsecaseError)
	// GetPhotoByIdSvc(ctx context.Context, userId uint64) (result Photo, err response.UsecaseError)
	// UpdatePhotoByIdSvc(ctx context.Context, userId uint64, input Photo) (result Photo, err response.UsecaseError)
	// DeletePhotoByIdSvc(ctx context.Context, userId uint64) (err response.UsecaseError)
}