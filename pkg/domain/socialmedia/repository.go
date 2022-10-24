package socialmedia

import "context"

type SocialMediaRepo interface {
	InsertSocialMedia(ctx context.Context, socialMedia *SocialMedia) (err error)
	GetById(ctx context.Context, socialMediaId uint64) (socialMedia SocialMedia, err error)
	GetByUserId(ctx context.Context, userId uint64) (socialMedias []SocialMedia, err error)
	UpdateSocialMedia(ctx context.Context, socialMedia *SocialMedia, input UpdateSocialMediaInput) (err error)
	DeleteSocialMediaById(ctx context.Context, socialMediaId uint64) (err error)
}