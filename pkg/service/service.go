package service

import (
	pack "go-serv-arduino"
	"go-serv-arduino/pkg/repository"
)

type Authorization interface {
	CreateUser(user pack.User) (int, error)
	GenerateToken(user pack.User) (string, error)
}

type News interface {
	GetNews() ([]pack.News, error)
}

type Service struct {
	Authorization
	News
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		News: NewNewsService(repository.News),
	}
}
