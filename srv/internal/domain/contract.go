package domain

type App interface {
	Create() error
	Update() error
	Delete() error
}
