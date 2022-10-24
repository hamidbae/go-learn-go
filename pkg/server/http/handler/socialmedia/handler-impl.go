package socialMedia

import (
	"errors"
	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"final-project/pkg/domain/socialmedia"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SocialMediaHdlImpl struct {
	socialMediaUsecase socialmedia.SocialMediaUsecase
}

func NewSocialMediaHandler(socialMediaUsecase socialmedia.SocialMediaUsecase) socialmedia.SocialMediaHandler {
	return &SocialMediaHdlImpl{socialMediaUsecase: socialMediaUsecase}
}

func (u *SocialMediaHdlImpl) AddSocialMediaHdl(ctx *gin.Context){
	userId := ctx.GetUint64("user_id")
	reqBody := socialmedia.AddSocialMediaInput{}
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

	result, usecaseError := u.socialMediaUsecase.AddSocialMediaSvc(ctx, reqBody, userId)
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
		Message: "add socialMedia success",
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


func (u *SocialMediaHdlImpl) GetSocialMediaByUserIdHdl(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	result, usecaseError := u.socialMediaUsecase.GetSocialMediasByUserIdSvc(ctx, userId)
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

	responseMessage := response.Response{
		Message: "get social media success",
		Data:    result,
	}

	ctx.JSONP(
		http.StatusOK,
		response.Build(
			http.StatusOK,
			responseMessage,
		),
	)
}

func (u *SocialMediaHdlImpl) UpdateSocialMediaHdl(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	paramId := ctx.Params.ByName("id")
	socialMediaId, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil{
		err := errors.New("params should be a number")
		responseMessage := response.Response{
			Message: "update socialMedia failed",
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

	reqBody := socialmedia.UpdateSocialMediaInput{}
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
	err = validate.Struct(reqBody)
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

	result, usecaseError := u.socialMediaUsecase.UpdateSocialMediaSvc(ctx, socialMediaId, userId, reqBody)
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

	responseMessage := response.Response{
		Message: "update socialMedia success",
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


func (u *SocialMediaHdlImpl) DeleteSocialMediaByIdHdl(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	paramId := ctx.Params.ByName("id")
	socialMediaId, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil{
		err := errors.New("params should be a number")
		responseMessage := response.Response{
			Message: "delete socialMedia failed",
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
	usecaseError := u.socialMediaUsecase.DeleteSocialMediaByIdSvc(ctx, userId, socialMediaId)
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

	responseMessage := response.Response{
		Message: "delete socialMedia success",
		Data: nil,
	}

	ctx.JSONP(
		http.StatusAccepted,
		response.Build(
			http.StatusAccepted,
			responseMessage,
		),
	)
}