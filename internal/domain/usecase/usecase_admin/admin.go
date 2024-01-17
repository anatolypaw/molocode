package usecase_admin

import (
	"log"
	"molocode/internal/domain/entity"
	"time"
)

/*
Методы для управления сервисом хранения марок
- Добавление продукта
- Получение списка добавленных продуктов
- Изменение продукта
*/
type iStorage interface {
	AddGood(entity.Good) error
	GetAllGoods() ([]entity.Good, error)
}

type AdminUseCase struct {
	store iStorage
}

func NewAdminUseCase(storeService iStorage) AdminUseCase {
	return AdminUseCase{store: storeService}
}

// Добавляет новый продукт
func (ths *AdminUseCase) AddGood(good entity.Good) error {
	log.Println("Добавление продукта")

	err := entity.ValidateGtin(good.Gtin)
	if err != nil {
		return err
	}
	good.CreatedAt = time.Now()

	return err
}

func (ths *AdminUseCase) GetAllGoods() ([]entity.Good, error) {
	// TODO валидировать ответ хранилища.
	// на корректность gtin
	return ths.store.GetAllGoods()
}
