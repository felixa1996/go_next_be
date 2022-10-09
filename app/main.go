package main

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	"github.com/felixa1996/go_next_be/app/common"
	"github.com/felixa1996/go_next_be/app/infra/database"
)

func main() {
	// todo move config into single function
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Infof("Service %s on DEBUG mode", viper.GetString("APP_NAME"))
	}

	dbManager := database.Manager{}
	err = dbManager.InitDB(viper.GetString("MONGODB_URI"))
	if err != nil {
		panic(err)
	}

	common.Init(dbManager)

	common.Application.Echo.Logger.Fatal(common.Application.Echo.Start(":" + viper.GetString("PORT")))
}
