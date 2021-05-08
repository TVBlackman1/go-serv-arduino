package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	pack "go-serv-arduino"
)

type AuthMysql struct {
	db *sqlx.DB
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user pack.User) (int, error){
	query := fmt.Sprintf("INSERT INTO Users (Login, Password) VALUES ('%s', '%s')", user.Name, user.Password)
	r.db.Query(query)

	return 0, nil
}

func (r *AuthMysql) GetUser(username string) (pack.User, error) {
	var user pack.User
	query := fmt.Sprintf("SELECT * from Users WHERE Login='%s'", username)
	err := r.db.Get(&user, query)

	return user, err
}