package main

import (
	"urbox-viettel-go/app"
	"urbox-viettel-go/cmd"
	"urbox-viettel-go/configs"
)

func main() {
	configs.Load()
	app.CreateNewApp()

	err := cmd.SetupCmd()
	if err != nil {
		panic(err)
	}
}
