package models

import "fmt"

type Like struct {
	UserId uint `gorm:"size:10;not null;" json:"user_id"`
	PostId uint `gorm:"size:10;not null;" json:"post_id"`
}

func LikePost(postId int, userId int) error {
	var like Like
	// var like = Like{PostId: uint(postId), UserId: uint(userId)}
	var err error

	err = DB.Where("user_id=? AND post_id=?", userId, postId).First(&like).Error
	// err = DB.First(&like).Error
	// fmt.Println(err.Error())
	fmt.Println(like)
	if err != nil {
		if err.Error() == "record not found" {
			like = Like{UserId: uint(userId), PostId: uint(postId)}
			DB.Create(&like)
			return nil
		}
		return err
	}
	DB.Delete(&like)

	return nil
}