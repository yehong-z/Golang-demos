package service

type IdGenerator interface {
	GetId() (int64, error)
}
