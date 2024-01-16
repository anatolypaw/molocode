package admin_usecase

import (
	"log"
	"molocode/internal/entity"
	"molocode/internal/service/storeservice"
)

type Usecase struct {
	store *storeservice.Service
}

func New(storeService *storeservice.Service) Usecase {
	return Usecase{store: storeService}
}

// Добавляет новый продукт
func (u Usecase) AddGood(good entity.Good) error {
	log.Println("Добавление продукта")
	err := u.store.AddGood(good)
	return err
}

// Возвращает все имеющиеся продукты
func (u Usecase) GetAllGoods() ([]entity.Good, error) {
	log.Println("Запрос всех продуктов")
	return u.store.GetAllGoods()
}
