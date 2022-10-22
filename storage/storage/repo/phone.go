package repo

type PhoneI interface {
	Get(id int) (*string, error)
}
