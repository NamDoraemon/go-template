package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
	"urbox-viettel-go/configs"
)

type ManageDB struct {
	Client *gorm.DB
}

func (db *ManageDB) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configs.GlobalConf.DBUser, configs.GlobalConf.DBPass, configs.GlobalConf.DBHost, configs.GlobalConf.DBPort, configs.GlobalConf.DBName)

	conn, err := gorm.Open(mysql.New(
		mysql.Config{
			DSN: dsn,
		}), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return err
	}
	manage, err := conn.DB()
	if err != nil {
		log.Print("sql select DB")
	}
	manage.SetConnMaxLifetime(30 * time.Minute)

	if err != nil {
		return err
	}

	tick := time.NewTicker(15 * time.Minute)
	go func(engine *gorm.DB) {
		for {
			select {
			case <-tick.C:
				err = manage.Ping()
				if err != nil {
					log.Println("sql can not ping")
				}
			}
		}
	}(conn)
	db.Client = conn
	return nil
}
