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
