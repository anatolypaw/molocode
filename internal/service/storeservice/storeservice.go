package storeservice

import (
	"molocode/internal/entity"
)

type Store interface {
	AddGood(good entity.Good) error
	GetGood(entity.Gtin) (entity.Good, error)
	GetAllGoods() ([]entity.Good, error)

	AddCode(entity.FullCode) error
}

// Проверяет входные данные, работает с хранилищем
type storeService struct {
	store Store
}

func NewStoreService(storage Store) *storeService {
	return &storeService{store: storage}
}
