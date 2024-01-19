package usecase_admin

import (
	"context"
	"log"
	"molocode/internal/app/entity"
	"molocode/internal/app/storage"
	"time"
)

/*
Методы для управления хранением марок
- Добавление продукта
- Получение списка добавленных продуктов
- Изменение продукта
*/
type AdminUsecase struct {
	goodRepo storage.IGood
}

func New(goodRepo storage.IGood) AdminUsecase {
	return AdminUsecase{goodRepo: goodRepo}
}

// Добавляет новый продукт
// Валидация gtin, desc
// Ошибка, если такой продукт с таким gtin уже существует
func (ths *AdminUsecase) AddGood(ctx context.Context, good entity.Good) error {
	log.Println("Добавление продукта")

	err := good.ValidateGtin()
	if err != nil {
		return err
	}

	err = good.ValidateDesc()
	if err != nil {
		return err
	}

	good.CreatedAt = time.Now()
	return err
}

func (ths *AdminUsecase) GetAllGoods(ctx context.Context) ([]entity.Good, error) {
	// TODO валидировать ответ хранилища
	// на корректность gtin
	return ths.goodRepo.GetAllGoods()
}
