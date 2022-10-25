package socialmedia

type SocialMedia struct {
	ID     uint64 `json:"id" gorm:"column:id;primaryKey"`
	Name   string `json:"name" gorm:"column:name;not null"`
	URL    string `json:"url" gorm:"column:url;not null"`
	UserId uint64 `json:"user_id" gorm:"column:user_id;not null"`
}

type AddSocialMediaInput struct {
	Name string `json:"name" validate:"required" example:"instagram"`
	URL  string `json:"url" validate:"required" example:"url"`
}

type UpdateSocialMediaInput struct {
	URL string `json:"url" validate:"required" example:"url"`
}