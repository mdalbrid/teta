package app

import (
	"context"
	"database/sql"
	"fmt"
)

func ConnectDatabase(ctx context.Context, driverName string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return db, fmt.Errorf("соединение с базой данных не сформированно: %w", err)
	}

	if err := db.Ping(); err != nil {
		return db, fmt.Errorf("соединение с базой данных не установлено: %w", err)
	}

	go func() {
		<-ctx.Done()

		_ = db.Close()
	}()

	return db, nil
}

func ConnectPostgreSQL(ctx context.Context, dsn string) (*sql.DB, error) {
	return ConnectDatabase(ctx, "postgres", dsn)
}
