package user

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `json:"id" gorm:"column:id;primaryKey"`
	Username  string         `json:"username" gorm:"column:username;not null;unique" validate:"required"`
	Email     string         `json:"email" gorm:"column:email;not null;unique" validate:"required,email"`
	Password  string         `json:"password,omitempty" gorm:"column:password;not null" validate:"required,min=6"`
	DoB       datatypes.Date `json:"date_of_birth" gorm:"column:date_of_birth;not null" validate:"required"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;DEFAULT:current_timestamp;not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;"`
	DeletedAt gorm.DeletedAt      `json:"deleted_at" gorm:"column:deleted_at;"`
}

type UserDto struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	DoB      string `json:"date_of_birth" validate:"required,ISO8601date"`
}

type UserCreatedDto struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int `json:"age"`
}

type UserGetDto struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

type UserUpdateDto struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}