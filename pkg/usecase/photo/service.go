package photo

import (
	"context"
	"errors"
	"final-project/pkg/domain/photo"
	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"log"
	"net/http"
)

type PhotoUsecaseImpl struct {
	photoRepo photo.PhotoRepo
}

func NewPhotoUsecase(photoRepo photo.PhotoRepo) photo.PhotoUsecase {
	return &PhotoUsecaseImpl{photoRepo: photoRepo}
}

func (u *PhotoUsecaseImpl) 	AddPhotoSvc(ctx context.Context, input photo.AddPhotoInput) (result photo.Photo, usecaseError response.UsecaseError){
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



