package repository

type App interface {
	Create() error
	Update() error
	Delete() error
}
