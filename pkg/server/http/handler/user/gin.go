package user

import (
	"errors"
	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"final-project/pkg/domain/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHdlImpl struct {
	userUsecase user.UserUsecase
}

func NewUserHandler(userUsecase user.UserUsecase) user.UserHandler {
	return &UserHdlImpl{userUsecase: userUsecase}
}

// @Summary get user by id
// @Description get user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int false "user id"
// @Success 200 {object} user.UserWithSocialMediaDto
// @Router /v1/user/{id} [get]
func (u *UserHdlImpl) GetUserByIdHdl(ctx *gin.Context) {
	paramId := ctx.Params.ByName("id")
	if paramId == ""{
		err := errors.New("params can't be null")
		responseMessage := response.Response{
			Message: "get user failed",
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

	userId, err := strconv.ParseUint(ctx.Params.ByName("id"), 10, 64)
	if err != nil{
		err := errors.New("params should be a number")
		responseMessage := response.Response{
			Message: "get user failed",
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

	result, usecaseError := u.userUsecase.GetUserByIdSvc(ctx, userId)
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
		Message: "get user success",
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

// @Summary update user
// @Description update user, auth required
// @Tags user
// @Security Bearer
// @Accept json
// @Produce json
// @Param user body user.UserUpdateDto true "user info"
// @Success 201 {object} user.UserUpdateResponseDto
// @Router /v1/user [put]
func (u *UserHdlImpl) UpdateUserHdl(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")

	reqBody := user.UserUpdateDto{}
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


	result, usecaseError := u.userUsecase.UpdateUserSvc(ctx, userId, reqBody)
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
		Message: "update user success",
		Data: result,
	}

	ctx.JSONP(
		http.StatusCreated,
		response.Build(
			http.StatusCreated,
			responseMessage,
		),
	)
}

// @Summary delete user
// @Description delete user, auth required
// @Tags user
// @Security Bearer
// @Accept json
// @Produce json
// @Success 201
// @Router /v1/user [delete]
func (u *UserHdlImpl) DeleteUserHdl(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")

	usecaseError := u.userUsecase.DeleteUserSvc(ctx, userId)
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
		Message: "delete user success",
		Data: nil,
	}

	ctx.JSONP(
		http.StatusCreated,
		response.Build(
			http.StatusCreated,
			responseMessage,
		),
	)
}