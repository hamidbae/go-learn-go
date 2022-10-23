package photo

import (
	"context"
	"final-project/config/postgres"
	"final-project/pkg/domain/photo"
	"log"
)

type PhotoRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewPhotoRepo(pgCln postgres.PostgresClient) photo.PhotoRepo {
	return &PhotoRepoImpl{pgCln: pgCln}
}

func (u *PhotoRepoImpl) InsertPhoto(ctx context.Context, insertedPhoto *photo.Photo) (err error) {
	log.Printf("%T - InsertPhoto is invoked]\n", u)
	defer log.Printf("%T - InsertPhoto executed\n", u)
	db := u.pgCln.GetClient()
	db.Model(&photo.Photo{}).Create(&insertedPhoto)

	if err = db.Error; err != nil {
		log.Printf("error when inserting photo\n")
	}
	return err
}