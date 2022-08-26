package repositories

import (
	"urbox-viettel-go/db"
	"urbox-viettel-go/models"
)

var OrderRepo *OrderRepository

type OrderRepository struct {
	BaseRepository
}

type OrderRepoBuilder struct {
	BaseRepository
}

func NewOrderRepository() *OrderRepository {
	if OrderRepo == nil {
		OrderRepo = &OrderRepository{}
	}
	return OrderRepo
}

func (repo *OrderRepository) setDB(db *db.ManageDB) *OrderRepository {
	repo.DB = db
	repo.Model = db.Client.Model(&models.ModelOrder{})
	return repo
}

func (repo *OrderRepository) build() *OrderRepository {
	return repo
}

func (repo *OrderRepository) First() {
	var a *models.ModelOrder
	repo.Model.Limit(1).Find(&a)
}
