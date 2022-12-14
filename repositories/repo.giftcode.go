package repositories

import (
	"go-template/db"
	"go-template/models"
)

var GiftCodeRepo *GiftCodeRepository

type GiftCodeRepository struct {
	BaseRepository
}

func NewGiftCodeRepository() *GiftCodeRepository {
	if GiftCodeRepo == nil {
		GiftCodeRepo = &GiftCodeRepository{}
	}
	return GiftCodeRepo
}

func (repo *GiftCodeRepository) setDB(db *db.ManageDB) *GiftCodeRepository {
	repo.DB = db
	repo.Model = db.Client.Model(&models.ModelGiftCode{})
	return repo
}

func (repo *GiftCodeRepository) build() *GiftCodeRepository {
	return repo
}
