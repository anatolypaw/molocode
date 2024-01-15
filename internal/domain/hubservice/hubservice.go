package hubservice

import (
	"molocode/internal/domain/entity"
)

type HubStore interface {
	AddGood(good entity.Good) error
	GetGood(entity.Gtin) (entity.Good, error)
	GetAllGoods() ([]entity.Good, error)

	AddCode(entity.FullCode) error
}

// Проверяет входные данные, работает с хранилищем
type hubService struct {
	storage HubStore
}

func New(storage HubStore) *hubService {
	return &hubService{storage: storage}
}
