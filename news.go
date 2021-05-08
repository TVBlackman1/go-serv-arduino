package pack

type News struct {
	Id      int    `json:"id" db:"ID"`
	Title   string `json:"title" db:"Title"`
	Content string `json:"content" db:"Content"`
	Author  string `json:"author" db:"Author"`
}
