package comment

import (
	"context"
	"final-project/config/postgres"
	"final-project/pkg/domain/comment"
	"log"
)

type CommentRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewCommentRepo(pgCln postgres.PostgresClient) comment.CommentRepo {
	return &CommentRepoImpl{pgCln: pgCln}
}

func (u *CommentRepoImpl) InsertComment(ctx context.Context, insertedComment *comment.Comment) (err error) {
	log.Printf("%T - InsertComment is invoked]\n", u)
	defer log.Printf("%T - InsertComment executed\n", u)
	db := u.pgCln.GetClient()
	db.Model(&comment.Comment{}).Create(&insertedComment)

	if err = db.Error; err != nil {
		log.Printf("error when inserting comment\n")
		return err
	}
	return err
}

func (u *CommentRepoImpl) GetById(ctx context.Context, commentId uint64) (result comment.Comment, err error) {
	log.Printf("%T - GetById is invoked]\n", u)
	defer log.Printf("%T - GetById executed\n", u)
	db := u.pgCln.GetClient()
	db.Model(&comment.Comment{}).Where("id", commentId).Find(&result)

	if err = db.Error; err != nil {
		log.Printf("error when get comment\n")
		return result, err
	}
	return result, err
}

func (u *CommentRepoImpl) UpdateComment(ctx context.Context, comment *comment.Comment, input comment.UpdateCommentInput) (err error){
	log.Printf("%T - UpdateComment is invoked]\n", u)
	defer log.Printf("%T - UpdateComment executed\n", u)

	db := u.pgCln.GetClient()
	db.Model(&comment).Update("message", input.Message)
	if err = db.Error; err != nil {
		log.Printf("error when update comment\n")
		return err
	}
	return err
}

func (u *CommentRepoImpl) DeleteCommentById(ctx context.Context, commentId uint64) (err error){
	log.Printf("%T - DeleteCommentById is invoked]\n", u)
	defer log.Printf("%T - DeleteCommentById executed\n", u)
	db := u.pgCln.GetClient()
	db.Delete(&comment.Comment{}, commentId)
	if err = db.Error; err != nil {
		log.Printf("error when delete comment with id %v\n", commentId)
		return err
	}
	return err
}