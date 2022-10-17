package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"go-app/internal/app"
	"go-app/internal/app/handler"
	"go-app/internal/domain/repository"
	appservice "go-app/internal/domain/service"
	"go-app/internal/http/service"
)

func main() {
	_, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	if err := app.InitConfig(); err != nil {
		logrus.Fatalf("error initialize configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	confDB := app.NewConfig(
		viper.GetString("db.user"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.dbname"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
	)

	db, err := confDB.NewPostgresDB()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			logrus.Println("ошибка закрытия базы данных: ", err)
		}
	}()

	repos := repository.NewRepository(db)
	services := appservice.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(service.Server)
	go func() {
		if err := srv.Run(viper.GetString("server_port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Println("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down %s", err.Error())
	}
}
