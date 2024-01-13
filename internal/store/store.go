package store

type Store interface {
	Good() GoodRepository
	Code() CodeRepository
}
