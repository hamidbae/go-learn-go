package photo

import (
	"final-project/pkg/domain/photo"
	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PhotoHdlImpl struct {
	photoUsecase photo.PhotoUsecase
}

func NewPhotoHandler(photoUsecase photo.PhotoUsecase) photo.PhotoHandler {
	return &PhotoHdlImpl{photoUsecase: photoUsecase}
}

func (u *PhotoHdlImpl) AddPhotoHdl(ctx *gin.Context){
	userId := ctx.GetUint64("user_id")

	reqBody := photo.AddPhotoInput{}
	if err := ctx.ShouldBind(&reqBody); err != nil {
		responseMessage := response.Response{
			Message: "error while binding body",
			InvalidArg: &response.InvalidArg{
				ErrorType:    errortype.INVALID_INPUT,
				ErrorMessage: err.Error(),
			},
		}

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Build(
				http.StatusBadRequest,
				responseMessage,
			),
		)
		return
	}


	validate := validator.New()
	err := validate.Struct(reqBody)
	if err != nil {
		responseMessage := response.Response{
			Message: "error while validating body",
			InvalidArg: &response.InvalidArg{
				ErrorType:    errortype.VALIDATION_FAILED,
				ErrorMessage: err.Error(),
			},
		}

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Build(
				http.StatusBadRequest,
				responseMessage,
			),
		)
		return
	}
	reqBody.UserId = userId

	result, usecaseError := u.photoUsecase.AddPhotoSvc(ctx, reqBody)
	if usecaseError.Error != nil {
		responseMessage := response.Response{
			Message: usecaseError.Message,
			InvalidArg: &response.InvalidArg{
				ErrorType:    usecaseError.ErrorType,
				ErrorMessage: usecaseError.Error.Error(),
			},
		}

		ctx.AbortWithStatusJSON(
			usecaseError.HttpCode,
			response.Build(
				usecaseError.HttpCode,
				responseMessage,
			),
		)
		return
	}
	result.User = nil

	responseMessage := response.Response{
		Message: "add photo success",
		Data: result,
	}

	ctx.JSONP(
		http.StatusAccepted,
		response.Build(
			http.StatusAccepted,
			responseMessage,
		),
	)
}