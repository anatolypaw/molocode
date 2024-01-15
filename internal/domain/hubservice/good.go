package hubservice

import (
	"molocode/internal/domain/entity"
	"time"
)

// Добавляет (создает) продукт. Ошбика, если такой уже есть
func (hs *hubService) AddGood(good entity.Good) error {
	err := good.Gtin.Validate()
	if err != nil {
		return err
	}
	good.CreatedAt = time.Now()
	return hs.storage.AddGood(good)
}

// Возвращает продукт по GTIN
func (hs *hubService) GetGood(gtin entity.Gtin) (entity.Good, error) {
	err := gtin.Validate()
	if err != nil {
		return entity.Good{}, err
	}
	return hs.storage.GetGood(gtin)
}

// Возвращает все продукты
func (hs *hubService) GetAllGoods() ([]entity.Good, error) {
	return hs.storage.GetAllGoods()
}
