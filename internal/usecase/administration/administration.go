package administration_usecase

import (
	"log"
	"molocode/internal/entity"
)

type storeService interface {
	AddGood(entity.Good) error
}

type usecase struct {
	store storeService
}

// Добавляет новый продукт
func (u usecase) AddGood(good entity.Good) error {
	log.Println("Добавление продукта")
	err := u.store.AddGood(good)
	return err
}

// Возвращает все имеющиеся продукты
func (u usecase) GetAllGoods() ([]entity.Good, error) {
	return []entity.Good{}, nil
}
