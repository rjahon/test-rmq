package postgres

import (
	"errors"
	"math/rand"

	"github.com/jmoiron/sqlx"
	"github.com/rjahon/labs-rmq/storage/storage/repo"
)

type phoneRepo struct {
	db *sqlx.DB
}

func NewPhone(db *sqlx.DB) repo.PhoneI {
	return &phoneRepo{
		db: db,
	}
}

func (pr *phoneRepo) Get(id int) (*string, error) {
	i := rand.Intn(4)
	if i == 1 {
		return nil, errors.New("server error")
	}

	var phone string

	query := `SELECT phone FROM phone p WHERE p.id=$1`
	err := pr.db.QueryRow(query, id).Scan(&phone)
	if err != nil {
		return nil, err
	}

	return &phone, nil
}
