package storage

import "molocode/internal/app/entity"

type IGood interface {
	AddGood(entity.Good) error
	GetGood(string) (entity.Good, error)
	GetAllGoods() ([]entity.Good, error)
}

type ICode interface {
	GetCountPrintAvaible(gtin string) (uint, error)
}
