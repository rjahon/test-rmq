package models

type Phone struct {
	ID    string `json:"id" db:"id"`
	Phone string `json:"phone"`
}
