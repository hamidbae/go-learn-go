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

func (u *PhotoRepoImpl) GetById(ctx context.Context, photoId uint64) (result photo.Photo, err error) {
	log.Printf("%T - GetById is invoked]\n", u)
	defer log.Printf("%T - GetById executed\n", u)
	db := u.pgCln.GetClient()
	db.Model(&photo.Photo{}).Where("id", photoId).Find(&result)

	if err = db.Error; err != nil {
		log.Printf("error when inserting photo\n")
	}
	return result, err
}

func (u *PhotoRepoImpl) GetByUserId(ctx context.Context, userId uint64) (photos []photo.Photo, err error) {
	log.Printf("%T - GetByUserId is invoked]\n", u)
	defer log.Printf("%T - GetByUserId executed\n", u)
	db := u.pgCln.GetClient()
	db.Model(&photo.Photo{}).Where("user_id", userId).Find(&photos)

	if err = db.Error; err != nil {
		log.Printf("error when inserting photo\n")
	}
	return photos, err
}

func (u *PhotoRepoImpl) UpdatePhoto(ctx context.Context, photo photo.Photo, input photo.UpdatePhotoInput) (err error){
	log.Printf("%T - UpdatePhoto is invoked]\n", u)
	defer log.Printf("%T - UpdatePhoto executed\n", u)

	db := u.pgCln.GetClient()
	db.Model(&photo).Update("title", input.Title).Update("caption", input.Caption)
	if err = db.Error; err != nil {
		log.Printf("error when update photo\n")
	}
	return err
}

func (u *PhotoRepoImpl) DeletePhotoById(ctx context.Context, photoId uint64) (err error){
	log.Printf("%T - DeletePhotoById is invoked]\n", u)
	defer log.Printf("%T - DeletePhotoById executed\n", u)
	db := u.pgCln.GetClient()
	db.Delete(&photo.Photo{}, photoId)
	if err = db.Error; err != nil {
		log.Printf("error when delete photo with id %v\n", photoId)
	}
	return err
}