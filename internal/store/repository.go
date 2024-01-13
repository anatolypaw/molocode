package store

import "molocode/internal/entities"

type GoodRepository interface {
	Create(good *entities.Good) error
	GetOne(gtin *entities.Gtin) (*entities.Good, error)
	GetAll() (*entities.Good, error)
}

type CodeRepository interface {
	AddForPrint() error
	GetForPrint(gtin *entities.Gtin) error
	Produce() error
}
