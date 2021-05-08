package service

import (
	pack "go-serv-arduino"
	"go-serv-arduino/pkg/repository"
)

type NewsService struct {
	repo repository.News
}

func NewNewsService(repo repository.News) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) GetNews() ([]pack.News, error) {
	return s.repo.GetNews()
}