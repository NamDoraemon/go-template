package main

import (
	"go-template/app"
	"go-template/cmd"
	"go-template/configs"
)

func main() {
	configs.Load()
	app.CreateNewApp()

	err := cmd.SetupCmd()
	if err != nil {
		panic(err)
	}
}
