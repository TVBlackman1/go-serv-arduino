package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	pack "go-serv-arduino"
	"go-serv-arduino/pkg/handler"
	"go-serv-arduino/pkg/repository"
	"go-serv-arduino/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	db, err := repository.NewMysql(repository.Config{
		User: viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		DBName: viper.GetString("db.dbname"),
	})

	if err != nil {
		log.Fatalf("Failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(pack.Server)
	if err:= srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}