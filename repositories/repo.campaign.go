package repositories

import (
	"go-template/db"
	"go-template/models"
)

var CampaignRepo *CampaignRepository

type CampaignRepository struct {
	BaseRepository
}

func NewCampaignRepository() *CampaignRepository {
	if CampaignRepo == nil {
		CampaignRepo = &CampaignRepository{}
	}
	return CampaignRepo
}

func (repo *CampaignRepository) setDB(db *db.ManageDB) *CampaignRepository {
	repo.DB = db
	repo.Model = db.Client.Model(&models.ModelCampaign{})
	return repo
}

func (repo *CampaignRepository) build() *CampaignRepository {
	return repo
}
