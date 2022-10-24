package photo

import "context"

type PhotoRepo interface {
	InsertPhoto(ctx context.Context, photo *Photo) (err error)
	GetById(ctx context.Context, photoId uint64) (photo Photo, err error)
	GetByUserId(ctx context.Context, userId uint64) (photos []Photo, err error)
	UpdatePhoto(ctx context.Context, photo Photo, input UpdatePhotoInput) (err error)
	DeletePhotoById(ctx context.Context, photoId uint64) (err error)
}