package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title       string `gorm:"size:255;not null;" json:"title"`
	Description string `gorm:"size:255;not null;" json:"description"`
}

func GetPosts() ([]Post, error) {

	var posts []Post

	var err error
	err = DB.Find(&posts).Error
	if err != nil {
		return []Post{}, err
	}
	return posts, nil
}

func GetOnePostById(id int) (Post, error) {
	var post Post

	var err error
	err = DB.Where("id=?", id).First(&post).Error
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

func (p *Post) SavePost() (*Post, error){
	var err error
	err = DB.Create(&p).Error 
	if err != nil {
		return &Post{}, err
	}
	return p, nil
}

func (p *Post) UpdatePost() (*Post, error){
	var err error
	err = DB.Save(&p).Error 
	if err != nil {
		return &Post{}, err
	}
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