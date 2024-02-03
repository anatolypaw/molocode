package admin

import (
	"context"
	"hub/internal/entity"
	"time"
)

/*
Методы для управления хранением марок
- Добавление продукта
- Получение списка добавленных продуктов
- Изменение продукта
*/
type iGoodRepo interface {
	Add(context.Context, entity.Good) error
	Get(context.Context, string) (entity.Good, error)
	GetAll(context.Context) ([]entity.Good, error)
}

type AdminUsecase struct {
	goodRepo iGoodRepo
}

func New(goodRepo iGoodRepo) AdminUsecase {
	return AdminUsecase{
		goodRepo: goodRepo,
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
	return usecase.goodRepo.Add(ctx, good)
}

func (ths *AdminUsecase) GetAllGoods(ctx context.Context,
) ([]entity.Good, error) {
	// TODO валидировать ответ хранилища
	// на корректность gtin
	return ths.goodRepo.GetAll(ctx)
}
