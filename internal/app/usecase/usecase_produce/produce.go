package usecase_produce

import (
	"context"
	"errors"
	"fmt"
	"molocode/internal/app/entity"
	"molocode/internal/app/repository"
)

type ProduceUsecase struct {
	goodRepository repository.IGoodRepository
	codeRepository repository.ICodeRepository
}

func New(goodRepository repository.IGoodRepository, codeRepository repository.ICodeRepository) ProduceUsecase {
	return ProduceUsecase{
		goodRepository: goodRepository,
		codeRepository: codeRepository,
	}
}

// Возвращает код для печати
func (usecase *ProduceUsecase) GetCodeForPrint(ctx context.Context, gtin string, terminal string) (entity.CodeForPrint, error) {
	// - Проверить корректность gtin
	err := entity.ValidateGtin(gtin)
	if err != nil {
		return entity.CodeForPrint{}, err
	}

	// - Проверить, разрешено ли для этого продукта выдача кодов для нанесения
	good, err := usecase.goodRepository.GetGood(ctx, gtin)
	if err != nil {
		return entity.CodeForPrint{}, fmt.Errorf("ошибка запроса продукта: %s", err)
	}

	if !good.AllowPrint {
		return entity.CodeForPrint{}, errors.New("для этотого продукта запрещено выдача кодов для нанесения")
	}

	// - Получить код для печати
	// - TODO Проверить корректность кода в ответе БД
	codeForPrint, err := usecase.codeRepository.GetCodeForPrint(ctx, gtin, terminal)
	if err != nil {
		return entity.CodeForPrint{}, err
	}

	return codeForPrint, nil
}
