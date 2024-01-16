package admin_usecase

import (
	"molocode/internal/domain/entity"
)

// Для этого сценария нужны данные о кодах
type hubService interface {
	AddGood(good entity.Good) error
	GetAllGoods() ([]entity.Good, error)
}

type AdminUsecase struct {
	hub hubService
}

// Добавляет продукт в базу.
func (u AdminUsecase) AddGood(good entity.Good) error {
	return u.hub.AddGood(good)
}

// Возвращает все имеющиеся продукты
func (u AdminUsecase) GetAllGoods() ([]entity.Good, error) {
	return u.hub.GetAllGoods()
}
