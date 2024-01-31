package admin

import (
	"context"
	"molocode/internal/app/entity"
	"molocode/internal/app/repository"
	"time"
)

/*
Методы для управления хранением марок
- Добавление продукта
- Получение списка добавленных продуктов
- Изменение продукта
*/
type AdminUsecase struct {
	goodRepository repository.IGoodRepository
}

func New(goodRepository repository.IGoodRepository) AdminUsecase {
	return AdminUsecase{
		goodRepository: goodRepository,
	}
}

// Добавляет новый продукт
// Валидация gtin, desc
// Ошибка, если такой продукт с таким gtin уже существует
func (usecase *AdminUsecase) AddGood(ctx context.Context, good entity.Good,
) error {
	err := good.ValidateGtin()
	if err != nil {
		return err
	}

	err = good.ValidateDesc()
	if err != nil {
		return err
	}

	good.CreatedAt = time.Now()
	return usecase.goodRepository.Add(ctx, good)
}

func (ths *AdminUsecase) GetAllGoods(ctx context.Context,
) ([]entity.Good, error) {
	// TODO валидировать ответ хранилища
	// на корректность gtin
	return ths.goodRepository.GetAll(ctx)
}
