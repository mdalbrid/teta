package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/mdalbrid/teta/srv/internal/app"
	"log"
)

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	conf := new(config)

	if err := app.ProcessConf("", conf); err != nil {
		panic(err)
	}
	log.Println("start ")

	//test := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable ",
	//	conf.User, conf.Password, conf.DBName, conf.Host, conf.Port)
	db, err := app.ConnectPostgreSQL(ctx, "")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Println("ошибка закрытия базы данных: ", err)
		}
	}()

	log.Println("end")

}
