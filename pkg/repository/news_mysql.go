package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	pack "go-serv-arduino"
	"log"
)

type NewsMysql struct {
	db *sqlx.DB
}

func NewNewsMysql(db *sqlx.DB) *NewsMysql {
	return &NewsMysql{db: db}
}

func (r *NewsMysql) GetNews() ([]pack.News, error) {
	var news []pack.News
	query := fmt.Sprintf("select * from News")
	rows, _ := r.db.Query(query)
	for rows.Next() {
		var singleNews pack.News

		var tmpAuthor sql.NullString
		err := rows.Scan(&singleNews.Id, &singleNews.Title, &singleNews.Content, &tmpAuthor)
		if tmpAuthor.Valid {
			singleNews.Author = tmpAuthor.String
		}
		if err != nil {
			return nil, nil
		}

		news = append(news, singleNews)
	}

	err := rows.Err()
	if err != nil {
		log.Print("Error2")
		return nil, nil
	}

	return news, nil
}
