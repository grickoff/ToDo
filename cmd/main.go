package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Life is hell")

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	//TODO: Поднять докер с базой
	//TODO: Написать свою первую миграцию
	//TODO: Реализовать коннект к БД
	//TODO: Реализовать старт сервера на gin
	//TODO: Реализовать контроллер + Юзкейс + стор
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
