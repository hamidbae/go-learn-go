package photo

import "context"

type PhotoRepo interface {
	InsertPhoto(ctx context.Context, photo *Photo) (err error)
	// GetById(ctx context.Context, photoId uint64) (photos []Photo, err error)
	// UpdateById(ctx context.Context, photoId uint64, photo Photo) (updatedPhoto Photo, err error)
	// DeleteById(ctx context.Context, photoId uint64) (err error)
}