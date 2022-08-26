package repositories

import (
	"gorm.io/gorm"
	"urbox-viettel-go/db"
)

type BaseRepository struct {
	DB      *db.ManageDB
	Model   *gorm.DB
	TblName string
}

func LoadRepositories(db *db.ManageDB) {
	NewOrderRepository().setDB(db).build()
	NewCampaignRepository().setDB(db).build()
	NewOrderDetailRepo().setDB(db).build()
	NewVoucherDetailRepository().setDB(db).build()
	NewLogRepository().setDB(db).build()
	NewGiftCodeRepository().setDB(db).build()
}
