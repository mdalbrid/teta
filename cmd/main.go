package main

import (
	"context"
)

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	db, err := ConnectPostgreSQL(ctx, "")
}