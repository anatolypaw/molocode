package storeservice

import (
	"molocode/internal/entity"
	"time"
)

// Добавляет (создает) продукт. Ошбика, если такой уже есть
func (hs *Service) AddGood(good entity.Good) error {
	err := entity.ValidateGtin(good.Gtin)
	if err != nil {
		return err
	}
	good.CreatedAt = time.Now()
	return hs.store.AddGood(good)
}

// Возвращает продукт по GTIN
func (hs *Service) GetGood(gtin string) (entity.Good, error) {
	err := entity.ValidateGtin(gtin)
	if err != nil {
		return entity.Good{}, err
	}
	return hs.store.GetGood(gtin)
}

// Возвращает все продукты
func (hs *Service) GetAllGoods() ([]entity.Good, error) {
	return hs.store.GetAllGoods()
}
