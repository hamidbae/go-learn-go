package socialmedia

import (
	"context"
	"errors"
	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"final-project/pkg/domain/socialmedia"
	"final-project/pkg/domain/user"
	"log"
	"net/http"
)

type SocialMediaUsecaseImpl struct {
	socialMediaRepo socialmedia.SocialMediaRepo
	userRepo user.UserRepo
}

func NewSocialMediaUsecase(socialMediaRepo socialmedia.SocialMediaRepo, userRepo user.UserRepo) socialmedia.SocialMediaUsecase {
	return &SocialMediaUsecaseImpl{socialMediaRepo: socialMediaRepo, userRepo: userRepo}
}

func (u *SocialMediaUsecaseImpl) AddSocialMediaSvc(ctx context.Context, input socialmedia.AddSocialMediaInput, userId uint64) (result socialmedia.SocialMedia, usecaseError response.UsecaseError){
	socialMedia := socialmedia.SocialMedia{
		Name: input.Name,
		URL: input.URL,
		UserId: userId,
	}

	err := u.socialMediaRepo.InsertSocialMedia(ctx, &socialMedia)
	if err != nil{
		log.Printf("error when inserting user:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "insert social media",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}

	return socialMedia, usecaseError
}

// func (u *SocialMediaUsecaseImpl) GetSocialMediaByIdSvc(ctx context.Context, socialMediaId uint64) (result socialMedia.SocialMedia, usecaseError response.UsecaseError){
	
// 	result, err := u.socialMediaRepo.GetById(ctx, socialMediaId)
// 	if err != nil{
// 		log.Printf("error when getting socialMedia:%v\n", err.Error())
// 		err = errors.New("internal server error")
// 		usecaseError = response.UsecaseError{
// 			HttpCode:  http.StatusInternalServerError,
// 			Message:   "get socialMedia failed",
// 			ErrorType: errortype.INTERNAL_SERVER_ERROR,
// 			Error:     err,
// 		}
// 		return result, usecaseError
// 	}

// 	if(result.ID == 0){
// 		err := errors.New("socialMedia not found")
// 		usecaseError = response.UsecaseError{
// 			HttpCode:  http.StatusOK,
// 			Message:   "get socialMedia failed",
// 			ErrorType: errortype.INVALID_INPUT,
// 			Error:     err,
// 		}
// 		return result, usecaseError
// 	}

// 	return result, usecaseError
// }

func (u *SocialMediaUsecaseImpl) GetSocialMediasByUserIdSvc(ctx context.Context, userId uint64) (result []socialmedia.SocialMedia, usecaseError response.UsecaseError){
	result, err := u.socialMediaRepo.GetByUserId(ctx, userId)
	if err != nil{
		log.Printf("error when getting socialMedia:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "get socialMedia failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}

	return result, usecaseError
}

func (u *SocialMediaUsecaseImpl) UpdateSocialMediaSvc(ctx context.Context, socialMediaId uint64, userId uint64, input socialmedia.UpdateSocialMediaInput) (result socialmedia.SocialMedia, usecaseError response.UsecaseError){
	socialMedia, err := u.socialMediaRepo.GetById(ctx, socialMediaId)
	if(err != nil){
		log.Printf("error when updating social media:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "update social media failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}
	if(socialMedia.ID == 0){
		err := errors.New("social media not found")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusOK,
			Message:   "update social media failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return result, usecaseError
	}
	if(userId != socialMedia.UserId){
		log.Printf("unauthorized user\n")
		err = errors.New("unauthorized")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusUnauthorized,
			Message:   "update social media failed",
			ErrorType: errortype.INVALID_SCOPE,
			Error:     err,
		}
		return result, usecaseError
	}

	err = u.socialMediaRepo.UpdateSocialMedia(ctx, &socialMedia, input)
	if(err != nil){
		log.Printf("error when update social media:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "update social media failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}

	return socialMedia, usecaseError
}

func (u *SocialMediaUsecaseImpl) DeleteSocialMediaByIdSvc(ctx context.Context, userId uint64, socialMediaId uint64) (usecaseError response.UsecaseError){
	socialMedia, err := u.socialMediaRepo.GetById(ctx, socialMediaId)
	if(err != nil){
		log.Printf("error when updating social media:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "delete social media failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return usecaseError
	}
	if(socialMedia.ID == 0){
		err := errors.New("social media not found")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusOK,
			Message:   "delete social media failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return usecaseError
	}
	if(userId != socialMedia.UserId){
		log.Printf("unauthorized user\n")
		err = errors.New("unauthorized")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusUnauthorized,
			Message:   "delete social media failed",
			ErrorType: errortype.INVALID_SCOPE,
			Error:     err,
		}
		return usecaseError
	}

	err = u.socialMediaRepo.DeleteSocialMediaById(ctx, socialMediaId)
	if(err != nil){
		log.Printf("error when deleting socialMedia:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "delete social media failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return usecaseError
	}

	return usecaseError
}



