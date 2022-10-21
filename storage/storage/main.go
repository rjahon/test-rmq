package storage

import (
	"github.com/rjahon/labs-rmq/storage/storage/postgres"
	"github.com/rjahon/labs-rmq/storage/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Phone() repo.PhoneI
}

type StoragePg struct {
	phoneRepo repo.PhoneI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &StoragePg{
		phoneRepo: postgres.NewPhone(db),
	}
}

func (s *StoragePg) Phone() repo.PhoneI {
	return s.phoneRepo
}
