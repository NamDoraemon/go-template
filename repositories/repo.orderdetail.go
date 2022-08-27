package repositories

import (
	"go-template/db"
	"go-template/models"
)

var OrderDetailRepo *OrderDetailRepository

type OrderDetailRepository struct {
	BaseRepository
}

func NewOrderDetailRepo() *OrderDetailRepository {
	if OrderDetailRepo == nil {
		OrderDetailRepo = &OrderDetailRepository{}
	}
	return OrderDetailRepo
}

func (repo *OrderDetailRepository) setDB(db *db.ManageDB) *OrderDetailRepository {
	repo.DB = db
	repo.Model = db.Client.Model(&models.ModelOrderDetail{})
	return repo
}

func (repo *OrderDetailRepository) build() *OrderDetailRepository {
	return repo
}
