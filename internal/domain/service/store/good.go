package service

import (
	"molocode/internal/domain/entity"
	"time"
)

// Добавляет (создает) продукт. Ошбика, если такой уже есть
func (service *Store) AddGood(good entity.Good) error {
	err := entity.ValidateGtin(good.Gtin)
	if err != nil {
		return err
	}
	good.CreatedAt = time.Now()
	return service.repo.AddGood(good)
}

// Возвращает продукт по GTIN
func (service *Store) GetGood(gtin string) (entity.Good, error) {
	err := entity.ValidateGtin(gtin)
	if err != nil {
		return entity.Good{}, err
	}
	return service.repo.GetGood(gtin)
}

// Возвращает все продукты
func (service *Store) GetAllGoods() ([]entity.Good, error) {
	// TODO валидация ответа хранилища
	return service.repo.GetAllGoods()
}

// Возвращает продукты, доступные для печати
func (service *Store) GetGoodsForPrint() ([]entity.Good, error) {
	// Получаем все продукты
	goods, err := service.GetAllGoods()
	if err != nil {
		return []entity.Good{}, err
	}

	// Выбираем из них доступные для печати
	var goodsForPrint []entity.Good
	for _, good := range goods {
		if good.GetCodeForPrint {
			goodsForPrint = append(goodsForPrint, good)
		}
	}
	return goodsForPrint, nil
}
