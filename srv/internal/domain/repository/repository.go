package repository

import "database/sql"

type Repository struct {
	App
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		App: NewAppPostgres(db),
	}
}
