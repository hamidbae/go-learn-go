package order

import (
	"assignment2/pkg/domain/user"
)

type Order struct {
	ID       uint64 	`json:"id" gorm:"column:id;primaryKey"`
	UserId   uint64 	`json:"user_id" gorm:"column:user_id;not null"`
	ItemName string 	`json:"item_name" gorm:"column:item_name;not null"`
	// User     user.User	`json:"user,omitempty" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type OrderWithUser struct{
	Order Order
	User user.User 		`json:"user,omitempty" gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
