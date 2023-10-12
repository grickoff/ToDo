package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grickoff/todo/internal/auth/controller"
	"github.com/grickoff/todo/internal/auth/storage"
	"github.com/grickoff/todo/internal/auth/usecases"
	"github.com/grickoff/todo/internal/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	//TODO: Поднять докер с базой +++
	//TODO: Написать свою первую миграцию +++
	//TODO: Реализовать коннект к БД +++

	db := initRepository()
	repo := storage.NewAuthStorage(db)
	authService := usecases.NewAuthService(repo)
	authControllers := controller.NewAuthController(authService)

	fmt.Println(authControllers)
	//TODO: Реализовать старт сервера на gin +++
	r := gin.Default()
	r.Run(":3000")

	// r.GET("/", func(c *gin.Context))

	//TODO: Реализовать контроллер + Юзкейс + стор +++
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initRepository() *sqlx.DB {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	return db
}
