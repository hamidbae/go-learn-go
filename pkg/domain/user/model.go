package user

import (
	"final-project/pkg/domain/socialmedia"
	"time"

	"gorm.io/datatypes"
)

type User struct {
	ID        uint64         `json:"id" gorm:"column:id;primaryKey"`
	Username  string         `json:"username" gorm:"column:username;not null;unique" validate:"required"`
	Email     string         `json:"email" gorm:"column:email;not null;unique" validate:"required,email"`
	Password  string         `json:"password,omitempty" gorm:"column:password;not null" validate:"required,min=6"`
	DoB       datatypes.Date `json:"date_of_birth" gorm:"column:date_of_birth;not null" validate:"required"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;DEFAULT:current_timestamp;not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;"`
	SocialMedias *[]socialmedia.SocialMedia `json:"social_medias,omitempty" gorm:"foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserDto struct {
	Username string `json:"username" validate:"required" example:"luigi"`
	Email    string `json:"email" validate:"required,email" example:"luigi@mail.com"`
	Password string `json:"password" validate:"required,min=6" example:"password"`
	DoB      string `json:"date_of_birth" validate:"required,ISO8601date" example:"1999-09-19"`
}

type UserCreatedDto struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int `json:"age"`
}

type UserWithSocialMediaDto struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	SocialMedias []socialmedia.SocialMedia `json:"social_medias"`
}

type UserUpdateDto struct {
	Username string `json:"username" validate:"required" example:"luigi"`
	Email    string `json:"email" validate:"required,email" example:"luigi@mail.com"`
}

type UserUpdateResponseDto struct {
	TokenId string `json:"token_id"`
}