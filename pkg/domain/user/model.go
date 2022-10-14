package user

type User struct {
	ID        uint64 `json:"id" gorm:"column:id;primaryKey"`
	FirstName string `json:"first_name" gorm:"column:first_name;not null"`
	LastName  string `json:"last_name" gorm:"column:last_name;not null"`
	Email     string `json:"email" gorm:"column:email;not null;unique"`
}
