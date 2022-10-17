package service

import (
	"go-app/internal/domain"
	"go-app/internal/domain/repository"
)

type Service struct {
	domain.App
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		App: NewAppService(repos.App),
	}
}
