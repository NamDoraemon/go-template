package app

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"go-template/configs"
	"go-template/db"
	"go-template/repositories"
	"log"
)

var App *GApp

type GApp struct {
	DB    *db.ManageDB
	Redis *redis.Client
}

func CreateNewApp() *GApp {
	App = &GApp{}
	App.LoadRedis()
	App.LoadDB()
	return App
}

func (app *GApp) LoadDB() {
	log.Println("Loading database client ...")
	app.DB = &db.ManageDB{}
	err := app.DB.Connect()
	if err != nil {
		log.Fatal(err)
	}
	repositories.LoadRepositories(app.DB)
	log.Println("Loaded database client ...")
}

func (app *GApp) LoadRedis() {
	log.Println("Loading redis client ...")
	address := fmt.Sprintf("%s:%s", configs.GlobalConf.RedisHost, configs.GlobalConf.RedisPort)
	app.Redis = redis.NewClient(&redis.Options{
		Addr: address,
		DB:   configs.GlobalConf.RedisDB,
	})
	log.Println("Loaded redis client ...")
}
