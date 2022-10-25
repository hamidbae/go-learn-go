package user

import (
	"context"
	"errors"
	"final-project/pkg/domain/auth"
	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"final-project/pkg/domain/user"
	"final-project/pkg/usecase/auth/token"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}

func (u *UserUsecaseImpl) GetUserByIdSvc(ctx context.Context, input uint64) (result user.UserWithSocialMediaDto, usecaseError response.UsecaseError){
	user, err := u.userRepo.GetUserWithSocialMediaById(ctx, input)
	if err != nil {
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "get user failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}
	if(user.Username == ""){
		err = errors.New("user not found")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusBadRequest,
			Message:   "get user failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return result, usecaseError
	}
	result.ID = user.ID
	result.Username = user.Username
	result.SocialMedias = *user.SocialMedias

	return result, usecaseError
}

func (u *UserUsecaseImpl) UpdateUserSvc(ctx context.Context, userId uint64, input user.UserUpdateDto) (result user.UserUpdateResponseDto, usecaseError response.UsecaseError){
	// find logged in user
	userLoggedIn, err := u.userRepo.GetUserById(ctx, userId)
	if err != nil {
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "update user failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}

	// is email exist
	userCheck, err := u.userRepo.GetUserByEmail(ctx, input.Email)
	if input.Email != userLoggedIn.Email && input.Email == userCheck.Email{
		err = errors.New(fmt.Sprintf("user with email %s exist", input.Email))
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusBadRequest,
			Message:   "update user failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return result, usecaseError
	}

	usernameCheck, err := u.userRepo.GetUserByUsername(ctx, input.Username)
	if input.Username != userLoggedIn.Username && input.Username == usernameCheck.Username{
		err = errors.New(fmt.Sprintf("user with username %s exist", input.Username))
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusBadRequest,
			Message:   "update user failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return result, usecaseError
	}

	user, err := u.userRepo.UpdateUser(ctx, userLoggedIn.ID, &input)
	if(err != nil){
		log.Printf("error when updating user:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "update user failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}

	claimID := auth.ClaimID{
		JWTID:    uuid.New(),
		Username: user.Username,
		Email:    user.Email,
		DOB:      time.Time(user.DoB),
	}
	tokenID, _ := token.CreateJWT(ctx, claimID)
	result.TokenId = tokenID

	return result, usecaseError
}

func (u *UserUsecaseImpl) DeleteUserSvc(ctx context.Context, userId uint64) (usecaseError response.UsecaseError){
	err := u.userRepo.DeleteUserById(ctx, userId)
	if(err != nil){
		log.Printf("error when deleting user:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "delete user failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return usecaseError
	}

	return usecaseError
}