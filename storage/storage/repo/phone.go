package repo

import "github.com/rjahon/labs-rmq/storage/models"

type PhoneI interface {
	Get(id int) (*models.Phone, error)
}
