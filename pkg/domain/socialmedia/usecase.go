package socialmedia

import (
	"context"
	"final-project/pkg/domain/response"
)

type SocialMediaUsecase interface {
	AddSocialMediaSvc(ctx context.Context, input AddSocialMediaInput, userId uint64) (result SocialMedia, err response.UsecaseError)
	// GetSocialMediaByIdSvc(ctx context.Context, socialMediaId uint64) (result SocialMedia, err response.UsecaseError)
	GetSocialMediasByUserIdSvc(ctx context.Context, userId uint64) (result []SocialMedia, err response.UsecaseError)
	UpdateSocialMediaSvc(ctx context.Context, socialMediaId uint64, userId uint64, input UpdateSocialMediaInput) (result SocialMedia, err response.UsecaseError)
	DeleteSocialMediaByIdSvc(ctx context.Context, socialMediaId uint64, userId uint64) (err response.UsecaseError)
}