package pack

type User struct {
	Id       int    `json:"id" db:"ID"`
	Name     string `json:"name" db:"Login"`
	Password string `json:"password" db:"Password"`
}
