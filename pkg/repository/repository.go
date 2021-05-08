package repository

import (
	"github.com/jmoiron/sqlx"
	pack "go-serv-arduino"
)

type Authorization interface {
	CreateUser(user pack.User) (int, error)
	GetUser(username string) (pack.User, error)
}

type News interface {
	GetNews() ([]pack.News, error)
}

type Repository struct {
	Authorization
	News
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		News: NewNewsMysql(db),
	}
}
