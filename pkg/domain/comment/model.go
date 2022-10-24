package comment

import (
	"time"

	"final-project/pkg/domain/user"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint64         `json:"id" gorm:"column:id;primaryKey"`
	Message   string         `json:"message" gorm:"column:message;not null" validate:"required"`
	UserId    uint64         `json:"user_id" gorm:"column:user_id;not null"`
	PhotoId   uint64         `json:"photo_id" gorm:"column:photo_id;not null" validate:"required"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;DEFAULT:current_timestamp;not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;"`
	User      *user.User      `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type AddCommentInput struct {
	Message   string         `json:"message" validate:"required"`
	PhotoId   uint64         `json:"photo_id" validate:"required"`
}

type UpdateCommentInput struct {
	Message   string         `json:"message" validate:"required"`
}