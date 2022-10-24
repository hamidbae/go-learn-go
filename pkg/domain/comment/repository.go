package comment

import "context"

type CommentRepo interface {
	GetById(ctx context.Context, commentId uint64) (comment Comment, err error)
	InsertComment(ctx context.Context, comment *Comment) (err error)
	UpdateComment(ctx context.Context, comment *Comment, input UpdateCommentInput) (err error)
	DeleteCommentById(ctx context.Context, commentId uint64) (err error)
}