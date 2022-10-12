package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title       string `gorm:"size:255;not null;" json:"title"`
	Description string `gorm:"size:255;not null;" json:"description"`
	UserId uint 	`gorm:"size:10;not null;" json:"user_id"`
	User 		User    `gorm:"foreignKey:UserId;references:ID" json:"user"`
	// User 		User    `gorm:"foreignKey:UserId;" json:"user"`
	Likes       []User `gorm:"many2many:likes;" json:"likes"`
}

func GetPosts() ([]Post, error) {

	var posts []Post
	// var post Post

	var err error 
	err = DB.Preload("User").Preload("Likes").Find(&posts).Error
	if err != nil {
		return []Post{}, err
	}
	return posts, nil
}

func GetOnePostById(id int) (Post, error) {
	var post Post

	var err error
	err = DB.Preload("User").Preload("Likes").Where("id=?", id).First(&post).Error
	if err != nil {
		return Post{}, err
	}
	post.User.Password = ""
	return post, nil
}

func (p *Post) SavePost() (*Post, error){
	var err error
	err = DB.Create(&p).Error 
	if err != nil {
		return &Post{}, err
	}
	p.User = User{}
	return p, nil
}

func (p *Post) UpdatePost() (*Post, error){
	var err error
	err = DB.Save(&p).Error 
	if err != nil {
		return &Post{}, err
	}
	p.User = User{}
	return p, nil
}

func (p *Post) DeletePost() (error){
	var err error
	err = DB.Delete(&p).Error 
	if err != nil {
		return err
	}
	return nil
}