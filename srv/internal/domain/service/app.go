package service

import "go-app/internal/domain/repository"

type appService struct {
	repo repository.App
}

func NewAppService(repo repository.App) *appService {
	return &appService{repo: repo}
}

func (a *appService) Create() error {
	//TODO implement me
	panic("implement me")
}

func (a *appService) Update() error {
	//TODO implement me
	panic("implement me")
}

func (a *appService) Delete() error {
	//TODO implement me
	panic("implement me")
}
