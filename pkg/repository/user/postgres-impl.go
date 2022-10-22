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
		log.Printf("error when getting user with email %v\n",email)
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
	}
	return err
}