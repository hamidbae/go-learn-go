package photo

import (
	"errors"
	"final-project/pkg/domain/photo"
	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"net/http"
	"strconv"

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

func (u *PhotoHdlImpl) GetPhotoByIdHdl(ctx *gin.Context){
	paramId := ctx.Params.ByName("id")
	if paramId == ""{
		err := errors.New("params can't be null")
		responseMessage := response.Response{
			Message: "get photo failed",
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

	photoId, err := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)
	if err != nil{
		err := errors.New("params should be a number")
		responseMessage := response.Response{
			Message: "get photo failed",
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

	result, usecaseError := u.photoUsecase.GetPhotoByIdSvc(ctx, photoId)
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
		Message: "get photo success",
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

func (u *PhotoHdlImpl) GetPhotoByUserIdHdl(ctx *gin.Context) {
	var userId  uint64
	userId = ctx.GetUint64("user_id")
	userIdString, _ := ctx.GetQuery("user_id")
	var err error
	if userIdString != "" {
		userId, err = strconv.ParseUint(userIdString, 10, 64)
		if err != nil{
			err := errors.New("params should be a number")
			responseMessage := response.Response{
				Message: "get photo failed",
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

	}

	result, usecaseError := u.photoUsecase.GetPhotosByUserIdSvc(ctx, userId)
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
		Message: "get photo success",
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


func (u *PhotoHdlImpl) UpdatePhotoHdl(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")

	paramId := ctx.Params.ByName("id")
	// if paramId == ""{
	// 	err := errors.New("params can't be null")
	// 	responseMessage := response.Response{
	// 		Message: "get photo failed",
	// 		InvalidArg: &response.InvalidArg{
	// 			ErrorType:    errortype.INVALID_INPUT,
	// 			ErrorMessage: err.Error(),
	// 		},
	// 	}

	// 	ctx.AbortWithStatusJSON(
	// 		http.StatusBadRequest,
	// 		response.Build(
	// 			http.StatusBadRequest,
	// 			responseMessage,
	// 		),
	// 	)
	// 	return
	// }

	photoId, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil{
		err := errors.New("params should be a number")
		responseMessage := response.Response{
			Message: "get photo failed",
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

	reqBody := photo.UpdatePhotoInput{}
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

	// validate req body
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

	result, usecaseError := u.photoUsecase.UpdatePhotoByIdSvc(ctx, photoId, userId, reqBody)
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
		Message: "update photo success",
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


func (u *PhotoHdlImpl) DeletePhotoByIdHdl(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	paramId := ctx.Params.ByName("id")
	// if paramId == ""{
	// 	err := errors.New("params can't be null")
	// 	responseMessage := response.Response{
	// 		Message: "get photo failed",
	// 		InvalidArg: &response.InvalidArg{
	// 			ErrorType:    errortype.INVALID_INPUT,
	// 			ErrorMessage: err.Error(),
	// 		},
	// 	}

	// 	ctx.AbortWithStatusJSON(
	// 		http.StatusBadRequest,
	// 		response.Build(
	// 			http.StatusBadRequest,
	// 			responseMessage,
	// 		),
	// 	)
	// 	return
	// }

	photoId, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil{
		err := errors.New("params should be a number")
		responseMessage := response.Response{
			Message: "delete photo failed",
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
	usecaseError := u.photoUsecase.DeletePhotoByIdSvc(ctx, userId, photoId)
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
		Message: "delete photo success",
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