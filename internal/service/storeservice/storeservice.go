package storeservice

import (
	"molocode/internal/entity"
)

type IStore interface {
	AddGood(good entity.Good) error
	GetGood(gtin string) (entity.Good, error)
	GetAllGoods() ([]entity.Good, error)

	AddCode(entity.FullCode) error
}

// Проверяет входные данные, работает с хранилищем
type Service struct {
	store IStore
}

func NewStoreService(storage IStore) *Service {
	return &Service{store: storage}
}
