package comment

import (
	"context"
	"errors"
	"final-project/pkg/domain/comment"
	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"final-project/pkg/domain/user"
	"log"
	"net/http"
)

type CommentUsecaseImpl struct {
	commentRepo comment.CommentRepo
	userRepo user.UserRepo
}

func NewCommentUsecase(commentRepo comment.CommentRepo, userRepo user.UserRepo) comment.CommentUsecase {
	return &CommentUsecaseImpl{commentRepo: commentRepo, userRepo: userRepo}
}

func (u *CommentUsecaseImpl) AddCommentSvc(ctx context.Context, input comment.AddCommentInput, userId uint64) (result comment.CommentDto, usecaseError response.UsecaseError){
	comment := comment.Comment{
		Message: input.Message,
		PhotoId: input.PhotoId,
		UserId: userId,
	}

	err := u.commentRepo.InsertComment(ctx, &comment)
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

	result.ID = comment.ID
	result.Message = comment.Message
	result.PhotoId = comment.PhotoId
	result.UserId = comment.UserId
	result.CreatedAt = comment.CreatedAt
	result.UpdatedAt = comment.UpdatedAt

	return result, usecaseError
}

// func (u *CommentUsecaseImpl) GetCommentByIdSvc(ctx context.Context, commentId uint64) (result comment.Comment, usecaseError response.UsecaseError){
	
// 	result, err := u.commentRepo.GetById(ctx, commentId)
// 	if err != nil{
// 		log.Printf("error when getting comment:%v\n", err.Error())
// 		err = errors.New("internal server error")
// 		usecaseError = response.UsecaseError{
// 			HttpCode:  http.StatusInternalServerError,
// 			Message:   "get comment failed",
// 			ErrorType: errortype.INTERNAL_SERVER_ERROR,
// 			Error:     err,
// 		}
// 		return result, usecaseError
// 	}

// 	if(result.ID == 0){
// 		err := errors.New("comment not found")
// 		usecaseError = response.UsecaseError{
// 			HttpCode:  http.StatusOK,
// 			Message:   "get comment failed",
// 			ErrorType: errortype.INVALID_INPUT,
// 			Error:     err,
// 		}
// 		return result, usecaseError
// 	}

// 	return result, usecaseError
// }

func (u *CommentUsecaseImpl) UpdateCommentByIdSvc(ctx context.Context, commentId uint64, userId uint64, input comment.UpdateCommentInput) (result comment.CommentDto, usecaseError response.UsecaseError){
	comment, err := u.commentRepo.GetById(ctx, commentId)
	if(userId != comment.UserId){
		log.Printf("unauthorized user\n")
		err = errors.New("unauthorized")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusUnauthorized,
			Message:   "update comment failed",
			ErrorType: errortype.INVALID_SCOPE,
			Error:     err,
		}
		return result, usecaseError
	}
	if(comment.ID == 0){
		err := errors.New("comment not found")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusOK,
			Message:   "get comment failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return result, usecaseError
	}

	err = u.commentRepo.UpdateComment(ctx, &comment, input)
	if(err != nil){
		log.Printf("error when update comment:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "update comment failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return result, usecaseError
	}

	result.ID = comment.ID
	result.Message = comment.Message
	result.PhotoId = comment.PhotoId
	result.UserId = comment.UserId
	result.CreatedAt = comment.CreatedAt
	result.UpdatedAt = comment.UpdatedAt

	return result, usecaseError
}

func (u *CommentUsecaseImpl) DeleteCommentByIdSvc(ctx context.Context, userId uint64, commentId uint64) (usecaseError response.UsecaseError){
	comment, err := u.commentRepo.GetById(ctx, commentId)
	if(err != nil){
		log.Printf("error when deleting comment:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "delete comment failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return usecaseError
	}
	if(comment.ID == 0){
		err := errors.New("comment not found")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusOK,
			Message:   "get comment failed",
			ErrorType: errortype.INVALID_INPUT,
			Error:     err,
		}
		return usecaseError
	}
	if(userId != comment.UserId){
		log.Printf("unauthorized user\n")
		err = errors.New("unauthorized")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusUnauthorized,
			Message:   "delete comment failed",
			ErrorType: errortype.INVALID_SCOPE,
			Error:     err,
		}
		return usecaseError
	}

	err = u.commentRepo.DeleteCommentById(ctx, commentId)
	if(err != nil){
		log.Printf("error when deleting comment:%v\n", err.Error())
		err = errors.New("internal server error")
		usecaseError = response.UsecaseError{
			HttpCode:  http.StatusInternalServerError,
			Message:   "delete comment failed",
			ErrorType: errortype.INTERNAL_SERVER_ERROR,
			Error:     err,
		}
		return usecaseError
	}

	return usecaseError
}



