package postgres

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/rjahon/labs-rmq/storage/models"
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

func (pr *phoneRepo) Get(id int) (*models.Phone, error) {
	var (
		resp   models.Phone
		number string
	)

	query := `SELECT phone FROM phone p WHERE p.id=$1`
	err := pr.db.QueryRow(query, id).Scan(&number)
	if err != nil {
		return nil, err
	}
	resp.ID = strconv.Itoa(id)
	resp.Phone = number

	return &resp, nil

}
