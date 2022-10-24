package socialmedia

import "final-project/pkg/domain/user"

type SocialMedia struct {
	ID     uint64     `json:"id" gorm:"column:id;primaryKey"`
	Name   string     `json:"name" gorm:"column:name;not null" validate:"required"`
	URL    string     `json:"url" gorm:"column:url;not null" validate:"required"`
	UserId uint64     `json:"user_id" gorm:"column:user_id;not null"`
	User   *user.User `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type AddSocialMediaInput struct {
	Name string `json:"name" validate:"required"`
	URL  string `json:"url" validate:"required"`
}

type UpdateSocialMediaInput struct {
	URL string `json:"url" validate:"required"`
}