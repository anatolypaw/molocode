package usecase_admin

import (
	"log"
	"molocode/internal/domain/entity"
	service "molocode/internal/domain/service/store"
)

/*
Методы для управления сервисом хранения марок
- Добавление продукта
- Получение списка добавленных продуктов
- Изменение продукта
*/

type AdminUseCase struct {
	store *service.Store
}

func NewAdminUseCase(storeService *service.Store) AdminUseCase {
	return AdminUseCase{store: storeService}
}

// Добавляет новый продукт
func (usecase *AdminUseCase) AddGood(good entity.Good) error {
	log.Println("Добавление продукта")
	err := usecase.store.AddGood(good)
	return err
}

// Возвращает все имеющиеся продукты
func (usecase *AdminUseCase) GetAllGoods() ([]entity.Good, error) {
	log.Println("Запрос всех продуктов")
	return usecase.store.GetAllGoods()
}
