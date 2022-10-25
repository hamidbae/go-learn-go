package comment

import (
	"context"
	"final-project/pkg/domain/response"
)

type CommentUsecase interface {
	AddCommentSvc(ctx context.Context, input AddCommentInput, userId uint64) (result CommentDto, err response.UsecaseError)
	// GetCommentByIdSvc(ctx context.Context, commentId uint64) (result Comment, err response.UsecaseError)
	UpdateCommentByIdSvc(ctx context.Context, commentId uint64, userId uint64, input UpdateCommentInput) (result CommentDto, err response.UsecaseError)
	DeleteCommentByIdSvc(ctx context.Context, userId uint64, commentId uint64) (err response.UsecaseError)
}