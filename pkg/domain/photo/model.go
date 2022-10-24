package photo

import (
	"time"

	"final-project/pkg/domain/comment"
	"final-project/pkg/domain/user"

	"gorm.io/gorm"
)

type Photo struct {
	ID        uint64         `json:"id" gorm:"column:id;primaryKey"`
	Title     string         `json:"title" gorm:"column:title;not null" validate:"required"`
	Caption   string         `json:"caption" gorm:"column:caption;not null" validate:"required"`
	URL       string         `json:"url" gorm:"column:url;not null" validate:"required"`
	UserId    uint64         `json:"user_id" gorm:"column:user_id;not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;DEFAULT:current_timestamp;not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;"`
	User      *user.User      `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments  *[]comment.Comment `json:"comments,omitempty" gorm:"foreignkey:PhotoId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type AddPhotoInput struct {
	Title     string         `validate:"required"`
	Caption   string         `validate:"required"`
	URL       string         `validate:"required"`
	UserId    uint64 
}

type UpdatePhotoInput struct {
	Title     string         `validate:"required"`
	Caption   string         `validate:"required"`
}