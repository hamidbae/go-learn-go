package auth

import (
	"context"
	"errors"
	"final-project/pkg/domain/auth"
	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"final-project/pkg/domain/user"
	"final-project/pkg/usecase/auth/token"
	"final-project/pkg/usecase/helper"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecaseImpl struct {
	userRepo user.UserRepo
}

func NewAuthUsecase(userRepo user.UserRepo) auth.AuthUsecase {
	return &AuthUsecaseImpl{userRepo: userRepo}
}

// func (u *AuthUsecaseImpl) GetAuthByEmailSvc(ctx context.Context, email string) (result user.Auth, err error) {
// 	log.Printf("%T - GetAuthByEmail is invoked]\n", u)
// 	defer log.Printf("%T - GetAuthByEmail executed\n", u)
// 	// get user from repository (database)
// 	log.Println("getting user from user repository")
// 	result, err = u.userRepo.GetAuthByEmail(ctx, email)
// 	if err != nil {
// 		// ini berarti ada yang salah dengan connection di database
// 		log.Println("error when fetching data from database: " + err.Error())
// 		err = errors.New("INTERNAL_SERVER_ERROR")
// 		return result, err
// 	}
// 	// check user id > 0 ?
// 	log.Println("checking user id")
// 	if result.ID <= 0 {
// 		// kalau tidak berarti user not found
// 		log.Println("user is not found: " + email)
// 		err = errors.New("NOT_FOUND")
// 		return result, err
// 	}
// 	return result, err
// }

func (u *AuthUsecaseImpl) RegisterSvc(ctx context.Context, input user.User) (result user.UserCreatedDto, usecaseError response.UsecaseError) {
	userCheck, err := u.userRepo.GetUserByEmail(ctx, input.Email)
	fmt.Println(userCheck)

	if userCheck.Email == input.Email {
		err = errors.New(fmt.Sprintf("user with email %s has been registered", input.Email))
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusBadRequest,
			Message:   "register failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return result, usecaseError
	}

	if userCheck.Username == input.Username {
		err = errors.New(fmt.Sprintf("user with username %s has been registered", input.Username))
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusBadRequest,
			Message:   "register failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return result, usecaseError
	}

	if err = u.userRepo.InsertUser(ctx, &input); err != nil {
		log.Printf("error when inserting user:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "register failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}

	age := helper.CountAge(time.Time(input.DoB)) 

	result = user.UserCreatedDto{
		ID: input.ID,
		Username: input.Username,
		Email: input.Email,
		Age: age,
	}

	return result, usecaseError
}

func (u *AuthUsecaseImpl) LoginSvc(ctx context.Context, input auth.Login) (result auth.Token, usecaseError response.UsecaseError) {
	// cek email exist
	userCheck, err := u.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "login failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}
	if userCheck.Email == "" {
		err = errors.New("user not found")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusBadRequest,
			Message:   "login failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return result, usecaseError
	}

	// cek password match
	err = bcrypt.CompareHashAndPassword([]byte(userCheck.Password), []byte(input.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		err = errors.New("password not match")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusBadRequest,
			Message:   "login failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return result, usecaseError
	}

	// create access token
	timeNow := time.Now()
	claimAccess := auth.ClaimAccess{
		JWTID:          uuid.New(),
		Subject:        userCheck.ID,
		Issuer:         "go-fga.com",
		Audience:       "user.go-fga.com",
		Scope:          "user",
		Type:           "ACCESS_TOKEN",
		IssuedAt:       timeNow.Unix(),
		NotValidBefore: timeNow.Unix(),
		ExpiredAt:      timeNow.Add(24 * time.Hour).Unix(),
	}
	accessToken, _ := token.CreateJWT(ctx, claimAccess)

	claimRefresh := auth.ClaimAccess{
		JWTID:          uuid.New(),
		Subject:        userCheck.ID,
		Issuer:         "go-fga.com",
		Audience:       "user.go-fga.com",
		Scope:          "user",
		Type:           "REFRESH_TOKEN",
		IssuedAt:       timeNow.Unix(),
		NotValidBefore: timeNow.Unix(),
		ExpiredAt:      timeNow.Add(1000 * time.Hour).Unix(),
	}
	refreshToken, _ := token.CreateJWT(ctx, claimRefresh)

	claimID := auth.ClaimID{
		JWTID:    uuid.New(),
		Username: userCheck.Username,
		Email:    userCheck.Email,
		DOB:      time.Time(userCheck.DoB),
	}
	tokenID, _ := token.CreateJWT(ctx, claimID)

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken
	result.TokenID = tokenID

	return result, usecaseError
}