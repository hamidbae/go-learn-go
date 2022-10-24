package socialmedia

import (
	"context"
	"final-project/config/postgres"
	socialMedia "final-project/pkg/domain/socialmedia"
	"log"
)

type SocialMediaRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewSocialMediaRepo(pgCln postgres.PostgresClient) socialMedia.SocialMediaRepo {
	return &SocialMediaRepoImpl{pgCln: pgCln}
}

func (u *SocialMediaRepoImpl) InsertSocialMedia(ctx context.Context, input *socialMedia.SocialMedia) (err error) {
	log.Printf("%T - InsertSocialMedia is invoked]\n", u)
	defer log.Printf("%T - InsertSocialMedia executed\n", u)
	db := u.pgCln.GetClient()
	db.Model(&socialMedia.SocialMedia{}).Create(&input)

	if err = db.Error; err != nil {
		log.Printf("error when inserting socialMedia\n")
		return err
	}
	return err
}

func (u *SocialMediaRepoImpl) GetById(ctx context.Context, socialMediaId uint64) (result socialMedia.SocialMedia, err error) {
	log.Printf("%T - GetById is invoked]\n", u)
	defer log.Printf("%T - GetById executed\n", u)
	db := u.pgCln.GetClient()
	db.Model(&socialMedia.SocialMedia{}).Where("id", socialMediaId).Find(&result)

	if err = db.Error; err != nil {
		log.Printf("error when finding social media\n")
		return result, err
	}
	return result, err
}

func (u *SocialMediaRepoImpl) GetByUserId(ctx context.Context, userId uint64) (result []socialMedia.SocialMedia, err error) {
	log.Printf("%T - GetByUserId is invoked]\n", u)
	defer log.Printf("%T - GetByUserId executed\n", u)
	db := u.pgCln.GetClient()
	db.Model(&socialMedia.SocialMedia{}).Where("user_id", userId).Find(&result)

	if err = db.Error; err != nil {
		log.Printf("error when finding social media\n")
		return result, err
	}
	return result, err
}

func (u *SocialMediaRepoImpl) UpdateSocialMedia(ctx context.Context, socialMedia *socialMedia.SocialMedia, input socialMedia.UpdateSocialMediaInput) (err error){
	log.Printf("%T - UpdateSocialMedia is invoked]\n", u)
	defer log.Printf("%T - UpdateSocialMedia executed\n", u)

	db := u.pgCln.GetClient()
	db.Model(&socialMedia).Update("url", input.URL)
	if err = db.Error; err != nil {
		log.Printf("error when update social media\n")
		return err
	}
	return err
}

func (u *SocialMediaRepoImpl) DeleteSocialMediaById(ctx context.Context, socialMediaId uint64) (err error){
	log.Printf("%T - DeleteSocialMediaById is invoked]\n", u)
	defer log.Printf("%T - DeleteSocialMediaById executed\n", u)
	db := u.pgCln.GetClient()
	db.Delete(&socialMedia.SocialMedia{}, socialMediaId)
	if err = db.Error; err != nil {
		log.Printf("error when delete social media with id %v\n", socialMediaId)
		return err
	}
	return err
}