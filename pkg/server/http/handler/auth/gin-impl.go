package auth

import (
	"final-project/pkg/domain/auth"
	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"final-project/pkg/domain/user"
	customvalidator "final-project/pkg/server/http/handler/custom-validator"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

type AuthHdlImpl struct {
	authUsecase auth.AuthUsecase
}

func NewAuthHandler(authUsecase auth.AuthUsecase) auth.AuthHandler {
	return &AuthHdlImpl{authUsecase: authUsecase}
}

// @Summary register
// @Description field date_of_birth should use pattern "yyyy-mm-dd"
// @Tags auth
// @Accept json
// @Produce json
// @Param user body user.UserDto true "user credential"
// @Success 201 {object} user.UserCreatedDto
// @Router /v1/auth/register [post]
func (u *AuthHdlImpl) RegisterHdl(ctx *gin.Context) {
	log.Printf("%T - RegisterHdl is invoked]\n", u)
	defer log.Printf("%T - RegisterHdl executed\n", u)

	var userDto user.UserDto
	if err := ctx.ShouldBind(&userDto); err != nil {
		responseMessage := response.Response{
			Message: "error while binding body",
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

	validate := validator.New()
	validate.RegisterValidation("ISO8601date", customvalidator.IsISO8601Date)
	err := validate.Struct(userDto)
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

	fmt.Println(userDto)
	dob, err := time.Parse("2006-01-02", userDto.DoB)
	if err != nil {
		responseMessage := response.Response{
			Message: "error while parsing date",
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		responseMessage := response.Response{
			Message: "error while hashing password",
			InvalidArg: &response.InvalidArg{
				ErrorType:    errortype.INTERNAL_CONNECTION_PROBLEM,
				ErrorMessage: err.Error(),
			},
		}

		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.Build(
				http.StatusInternalServerError,
				responseMessage,
			),
		)
		return
	}

	var user = user.User{
		Username: userDto.Username,
		Email:    userDto.Email,
		Password: string(hashedPassword),
		DoB:      datatypes.Date(dob),
	}

	result, usecaseError := u.authUsecase.RegisterSvc(ctx, user)
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
		Message: "register success",
		Data:    result,
	}

	ctx.JSONP(
		http.StatusCreated,
		response.Build(
			http.StatusCreated,
			responseMessage,
		),
	)
}

// @Summary login
// @Description login user to get auth token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginBody body auth.Login true "login info"
// @Success 200 {object} auth.Token
// @Router /v1/auth/login [post]
func (u *AuthHdlImpl) LoginHdl(ctx *gin.Context) {
	// binding body
	reqBody := auth.Login{}
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

	// create token
	result, usecaseError := u.authUsecase.LoginSvc(ctx, reqBody)
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

	// response result if success
	responseMessage := response.Response{
		Message: "login success",
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

// @Summary refresh token
// @Description refresh token to extend token expiry, required auth with refresh_token
// @Tags auth
// @Security Bearer
// @Accept json
// @Produce json
// @Success 200 {object} auth.Token
// @Router /v1/auth/refresh-token [post]
func (u *AuthHdlImpl) RefreshTokenHdl(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	result, usecaseError := u.authUsecase.RefreshTokenSvc(ctx, userId)
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
		Message: "refresh token success",
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