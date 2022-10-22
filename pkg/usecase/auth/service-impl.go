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

func (u *AuthUsecaseImpl) RegisterSvc(ctx context.Context, input user.User) (result user.UserCreatedDto, usecaseError response.UsecaseError) {
	userCheck, err := u.userRepo.GetUserByEmail(ctx, input.Email)
	// fmt.Println(userCheck)

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

	usernameCheck, err := u.userRepo.GetUserByUsername(ctx, input.Username)
	if usernameCheck.Username == input.Username {
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
		ID:       input.ID,
		Username: input.Username,
		Email:    input.Email,
		Age:      age,
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

func (u *AuthUsecaseImpl) RefreshTokenSvc(ctx context.Context, userId uint64) (result auth.Token, usecaseError response.UsecaseError) {
	
	user, err := u.userRepo.GetUserById(ctx, userId)
	if err != nil {
		err := errors.New("did not recognize user after middleware")
		usecaseError := response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "internal server errror",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}

	// create access token
	timeNow := time.Now()
	claimAccess := auth.ClaimAccess{
		JWTID:          uuid.New(),
		Subject:        userId,
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
		Subject:        userId,
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
		Username: user.Username,
		Email:    user.Email,
		DOB:      time.Time(user.DoB),
	}
	tokenID, _ := token.CreateJWT(ctx, claimID)

	result.AccessToken = accessToken
	result.RefreshToken = refreshToken
	result.TokenID = tokenID

	return result, usecaseError
}