package service

import (
	"molocode/internal/domain/entity"
)

type IRepository interface {
	AddGood(good entity.Good) error
	GetGood(gtin string) (entity.Good, error)
	GetAllGoods() ([]entity.Good, error)

	AddCode(entity.FullCode) error
	CountPrintAvaibleCode(gtin string) (uint, error)
}

type Store struct {
	repo IRepository
}

func NewStoreService(repo IRepository) *Store {
	return &Store{repo: repo}
}
