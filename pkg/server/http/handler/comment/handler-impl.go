package comment

import (
	"errors"
	"final-project/pkg/domain/comment"
	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CommentHdlImpl struct {
	commentUsecase comment.CommentUsecase
}

func NewCommentHandler(commentUsecase comment.CommentUsecase) comment.CommentHandler {
	return &CommentHdlImpl{commentUsecase: commentUsecase}
}

// @Summary add comment to post
// @Description add comment, auth required
// @Tags comment
// @Security Bearer
// @Accept json
// @Produce json
// @Param comment body comment.AddCommentInput true "comment info"
// @Success 201 {object} comment.CommentDto
// @Router /v1/comment [post]
func (u *CommentHdlImpl) AddCommentHdl(ctx *gin.Context){
	userId := ctx.GetUint64("user_id")
	reqBody := comment.AddCommentInput{}
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

	result, usecaseError := u.commentUsecase.AddCommentSvc(ctx, reqBody, userId)
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
		Message: "add comment success",
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

// @Summary update comment
// @Description update comment, auth required
// @Tags comment
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "comment id"
// @Param comment body comment.UpdateCommentInput true "comment info"
// @Success 201 {object} comment.CommentDto
// @Router /v1/comment/{id} [put]
func (u *CommentHdlImpl) UpdateCommentHdl(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	paramId := ctx.Params.ByName("id")
	commentId, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil{
		err := errors.New("params should be a number")
		responseMessage := response.Response{
			Message: "get comment failed",
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

	reqBody := comment.UpdateCommentInput{}
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

	result, usecaseError := u.commentUsecase.UpdateCommentByIdSvc(ctx, commentId, userId, reqBody)
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
		Message: "update comment success",
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

// @Summary delete comment by id
// @Description delete comment, auth required
// @Tags comment
// @Security Bearer
// @Accept json
// @Produce json
// @Param id path int true "comment id"
// @Success 201
// @Router /v1/comment/{id} [delete]
func (u *CommentHdlImpl) DeleteCommentByIdHdl(ctx *gin.Context) {
	userId := ctx.GetUint64("user_id")
	paramId := ctx.Params.ByName("id")
	commentId, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil{
		err := errors.New("params should be a number")
		responseMessage := response.Response{
			Message: "delete comment failed",
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
	usecaseError := u.commentUsecase.DeleteCommentByIdSvc(ctx, userId, commentId)
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
		Message: "delete comment success",
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