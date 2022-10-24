package photo

import (
	"context"
	"errors"
	"final-project/pkg/domain/photo"
	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"final-project/pkg/domain/user"
	"log"
	"net/http"
)

type PhotoUsecaseImpl struct {
	photoRepo photo.PhotoRepo
	userRepo user.UserRepo
}

func NewPhotoUsecase(photoRepo photo.PhotoRepo, userRepo user.UserRepo) photo.PhotoUsecase {
	return &PhotoUsecaseImpl{photoRepo: photoRepo, userRepo: userRepo}
}

func (u *PhotoUsecaseImpl) AddPhotoSvc(ctx context.Context, input photo.AddPhotoInput) (result photo.Photo, usecaseError response.UsecaseError){
	photo := photo.Photo{
		Title: input.Title,
		Caption: input.Caption,
		URL: input.URL,
		UserId: input.UserId,
	}

	err := u.photoRepo.InsertPhoto(ctx, &photo)
	if err != nil{
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

	return photo, usecaseError
}

func (u *PhotoUsecaseImpl)	GetPhotoByIdSvc(ctx context.Context, photoId uint64) (result photo.Photo, usecaseError response.UsecaseError){
	
	result, err := u.photoRepo.GetById(ctx, photoId)
	if err != nil{
		log.Printf("error when getting photo:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "get photo failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}

	if(result.ID == 0){
		err := errors.New("photo not found")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusOK,
			Message:   "get photo failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return result, usecaseError
	}

	return result, usecaseError
}

func (u *PhotoUsecaseImpl) 	GetPhotosByUserIdSvc(ctx context.Context, userId uint64) (result []photo.Photo, usecaseError response.UsecaseError){
	user, _ := u.userRepo.GetUserById(ctx, userId)
	if(user.ID == 0){
		err := errors.New("user not found")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusOK,
			Message:   "get photo failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return result, usecaseError
	}

	result, err := u.photoRepo.GetByUserId(ctx, userId)
	if err != nil{
		log.Printf("error when getting photo:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "get photo failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}

	return result, usecaseError
}

func (u *PhotoUsecaseImpl) UpdatePhotoByIdSvc(ctx context.Context, photoId uint64, userId uint64, input photo.UpdatePhotoInput) (result photo.Photo, usecaseError response.UsecaseError){
	photo, err := u.photoRepo.GetById(ctx, photoId)
	if(userId != photo.UserId){
		log.Printf("unauthorized user\n")
		err = errors.New("unauthorized")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusUnauthorized,
			Message:   "update photo failed",
			ErrorType: errortype.INVALID_SCOPE,
			Error:     err,
		}
		return result, usecaseError
	}

	err = u.photoRepo.UpdatePhoto(ctx, photo, input)
	if(err != nil){
		log.Printf("error when update photo:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "delete user failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}

	return photo, usecaseError
}

func (u *PhotoUsecaseImpl) DeletePhotoByIdSvc(ctx context.Context, userId uint64, photoId uint64) (usecaseError response.UsecaseError){
	photo, err := u.photoRepo.GetById(ctx, photoId)
	if(err != nil){
		log.Printf("error when deleting photo:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "delete photo failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return usecaseError
	}
	if(userId != photo.UserId){
		log.Printf("unauthorized user\n")
		err = errors.New("unauthorized")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusUnauthorized,
			Message:   "delete photo failed",
			ErrorType: errortype.INVALID_SCOPE,
			Error:     err,
		}
		return usecaseError
	}

	err = u.photoRepo.DeletePhotoById(ctx, photoId)
	if(err != nil){
		log.Printf("error when deleting photo:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "delete photo failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return usecaseError
	}

	return usecaseError
}



