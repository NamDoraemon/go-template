package repositories

import (
	"urbox-viettel-go/db"
	"urbox-viettel-go/models"
)

var LogRepo *LogRepository

type LogRepository struct {
	BaseRepository
}

func NewLogRepository() *LogRepository {
	if LogRepo == nil {
		LogRepo = &LogRepository{}
	}
	return LogRepo
}

func (repo *LogRepository) setDB(db *db.ManageDB) *LogRepository {
	repo.DB = db
	repo.Model = db.Client.Model(&models.ModelLog{})
	return repo
}

func (repo *LogRepository) build() *LogRepository {
	return repo
}
