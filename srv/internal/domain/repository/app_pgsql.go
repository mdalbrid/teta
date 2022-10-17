package repository

import "database/sql"

type appQuery struct {
	db *sql.DB
}

func NewAppPostgres(db *sql.DB) *appQuery {
	return &appQuery{db: db}
}

func (a *appQuery) Create() error {
	//TODO implement me
	panic("implement me")
}

func (a *appQuery) Update() error {
	//TODO implement me
	panic("implement me")
}

func (a *appQuery) Delete() error {
	//TODO implement me
	panic("implement me")
}
