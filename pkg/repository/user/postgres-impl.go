package user

import (
	"context"
	"log"

	"final-project/config/postgres"
	"final-project/pkg/domain/user"
)

type UserRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewUserRepo(pgCln postgres.PostgresClient) user.UserRepo {
	return &UserRepoImpl{pgCln: pgCln}
}

func (u *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (result user.User, err error) {
	log.Printf("%T - GetUserByEmail is invoked]\n", u)
	defer log.Printf("%T - GetUserByEmail executed\n", u)
	db := u.pgCln.GetClient()
	db.Model(&user.User{}).Where("email = ?", email).Find(&result)
	if err = db.Error; err != nil {
		log.Printf("error when find user with email %v\n",email)
		return result, err
	}
	return result, err
}

func  (u *UserRepoImpl) GetUserById(ctx context.Context, userId uint64) (result user.User, err error){
	log.Printf("%T - GetUserById is invoked]\n", u)
	defer log.Printf("%T - GetUserById executed\n", u)
	db := u.pgCln.GetClient()
	db.Model(&user.User{}).Where("id = ?", userId).Find(&result)
	if err = db.Error; err != nil {
		log.Printf("error when getting user with id %v\n",userId)
		return result, err
	}
	return result, err
}

func  (u *UserRepoImpl) GetUserWithSocialMediaById(ctx context.Context, userId uint64) (result user.User, err error){
	log.Printf("%T - GetUserById is invoked]\n", u)
	defer log.Printf("%T - GetUserById executed\n", u)
	db := u.pgCln.GetClient()
	// db.Model(&user.User{}).Where("id = ?", userId).Preload("SocialMedia").Find(&result).Where()
	// db.Preload("SocialMedia").Model(&user.User{}).Where("id = ?", userId).Find(&result)
	db.Preload("SocialMedias").Find(&result, userId)
	if err = db.Error; err != nil {
		log.Printf("error when getting user with id %v\n",userId)
		return result, err
	}
	return result, err
}

func  (u *UserRepoImpl) GetUserByUsername(ctx context.Context, username string) (result user.User, err error){
	log.Printf("%T - GetUserByUsername is invoked]\n", u)
	defer log.Printf("%T - GetUserByUsername executed\n", u)
	db := u.pgCln.GetClient()
	db.Model(&user.User{}).Where("username = ?", username).Find(&result)
	if err = db.Error; err != nil {
		log.Printf("error when getting user with username %v\n",username)
		return result, err
	}
	return result, err
}

func (u *UserRepoImpl) InsertUser(ctx context.Context, insertedUser *user.User) (err error) {
	log.Printf("%T - InsertUser is invoked]\n", u)
	defer log.Printf("%T - InsertUser executed\n", u)
	db := u.pgCln.GetClient()
	db.Model(&user.User{}).Create(&insertedUser)

	if err = db.Error; err != nil {
		log.Printf("error when inserting user with email %v\n",insertedUser.Email)
		return err
	}
	return err
}

func (u *UserRepoImpl) UpdateUser(ctx context.Context, userId uint64, userUpdate *user.UserUpdateDto) (user user.User, err error) {
	log.Printf("%T - UpdateUser is invoked]\n", u)
	defer log.Printf("%T - UpdateUser executed\n", u)

	db := u.pgCln.GetClient()
	db.Model(&user).Where("id = ?", userId).Update("username", userUpdate.Username).Update("email", userUpdate.Email)
	if err = db.Error; err != nil {
		log.Printf("error when update user with email %v\n",user.Email)
		return user, err
	}
	return user, err
}

func (u *UserRepoImpl) DeleteUserById(ctx context.Context, userId uint64) (err error) {
	log.Printf("%T - DeleteUserById is invoked]\n", u)
	defer log.Printf("%T - DeleteUserById executed\n", u)

	db := u.pgCln.GetClient()
	// db.Model(&user.User{}).Where("id = ?", userId).Delete()
	db.Delete(&user.User{}, userId)
	if err = db.Error; err != nil {
		log.Printf("error when update user with id %v\n", userId)
	}
	return err
}