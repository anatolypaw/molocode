package storeservice

import (
	"molocode/internal/entity"
	"time"
)

// Добавляет (создает) продукт. Ошбика, если такой уже есть
func (hs *storeService) AddGood(good entity.Good) error {
	err := good.Gtin.Validate()
	if err != nil {
		return err
	}
	good.CreatedAt = time.Now()
	return hs.store.AddGood(good)
}

// Возвращает продукт по GTIN
func (hs *storeService) GetGood(gtin entity.Gtin) (entity.Good, error) {
	err := gtin.Validate()
	if err != nil {
		return entity.Good{}, err
	}
	return hs.store.GetGood(gtin)
}

// Возвращает все продукты
func (hs *storeService) GetAllGoods() ([]entity.Good, error) {
	return hs.store.GetAllGoods()
}
