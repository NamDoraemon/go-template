package repositories

import (
	"gorm.io/gorm"
	"log"
	"urbox-viettel-go/db"
	"urbox-viettel-go/models"
)

var VoucherDetailRepo *VoucherDetailRepository

type VoucherDetailRepository struct {
	BaseRepository
}

func NewVoucherDetailRepository() *VoucherDetailRepository {
	if VoucherDetailRepo == nil {
		VoucherDetailRepo = &VoucherDetailRepository{}
	}
	return VoucherDetailRepo
}

func (repo *VoucherDetailRepository) setDB(db *db.ManageDB) *VoucherDetailRepository {
	repo.DB = db
	repo.Model = db.Client.Model(&models.ModelVoucherDetail{})
	return repo
}

func (repo *VoucherDetailRepository) build() *VoucherDetailRepository {
	return repo
}

func (repo *VoucherDetailRepository) FindById(id int) *models.ModelVoucherDetail {
	var result *models.ModelVoucherDetail
	repo.Model.Where("id = ?", id).Limit(1).Find(&result)
	return result
}

func (repo *VoucherDetailRepository) HoldCode(id int, quantity int) bool {
	if quantity == 0 {
		quantity = 1
	}

	updated := repo.Model.Where(
		"id = ?", id,
	).Where(
		"sales_limit >= ?", quantity,
	).Where(
		"status = ?", 2,
	).Where(
		"urbox_status = ?", 2,
	).Update("sales_limit", gorm.Expr("sales_limit + ?", quantity))

	if updated.Error != nil {
		log.Println(updated.Error)
		return false
	}

	if updated.RowsAffected == 0 {
		return false
	}

	return true
}
