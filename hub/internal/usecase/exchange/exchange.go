package exchange

import (
	"context"
	"errors"
	"hub/internal/entity"
	"time"
)

type iGoodRepo interface {
	Get(context.Context, string) (entity.Good, error)
	GetAll(context.Context) ([]entity.Good, error)
}

type iCodeRepo interface {
	AddCode(context.Context, entity.FullCode) error
	GetCountPrintAvaible(context.Context, string) (uint, error)
}

type ExchangeUsecase struct {
	goodRepo iGoodRepo
	codeRepo iCodeRepo
}

func New(goodRepo iGoodRepo, codeRepo iCodeRepo) ExchangeUsecase {
	return ExchangeUsecase{
		goodRepo: goodRepo,
		codeRepo: codeRepo,
	}
}

// Возвращаемая кейсом структура
type CodeReq struct {
	Gtin     string
	Desc     string
	Required uint
}

// Возвращает список продуктов, требующих наполнения кодами для печати
// и количество требуемых кодов
func (eu *ExchangeUsecase) GetGoodsReqCodes(ctx context.Context,
) ([]CodeReq, error) {
	// - Получить продукты
	allGoods, err := eu.goodRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// - Выбрать те, для которых включено наполнения кодами
	var goodsAvaibleForPrint []entity.Good
	for _, good := range allGoods {
		if good.GetCodeForPrint {
			goodsAvaibleForPrint = append(goodsAvaibleForPrint, good)
		}
	}

	// - Для каждого продукта получить доступное количество кодов
	var codesReq []CodeReq
	for _, good := range goodsAvaibleForPrint {
		avaibleCount, err := eu.codeRepo.GetCountPrintAvaible(ctx, good.Gtin)
		if err != nil {
			return nil, err
		}

		requiredCount := good.StoreCount - avaibleCount
		if requiredCount > 0 {
			codesReq = append(codesReq, CodeReq{
				Gtin:     good.Gtin,
				Desc:     good.Desc,
				Required: requiredCount,
			})
		}

	}

	// - Вернуть продукт, описание и недостающее количество кодов
	return codesReq, nil
}

// Добавляет код для печати
func (usecase *ExchangeUsecase) AddCodeForPrint(
	ctx context.Context,
	code entity.Code,
	source string,
) error {
	// - Проверить корректность кода
	err := code.Validate()
	if err != nil {
		return err
	}

	// - Проверить, разрешено ли для этого продукта добавление кодов
	good, err := usecase.goodRepo.Get(ctx, code.Gtin)
	if err != nil {
		return err
	}

	if !good.GetCodeForPrint {
		return errors.New("для этотого продукта запрещено получение кодов")
	}

	// - Добавить код для печати
	fullCode := entity.FullCode{
		Code: code,
		SourceInfo: entity.SourceInfo{
			Name: source,
			Time: time.Now(),
		},
		PrintInfo: entity.PrintInfo{
			Avaible: true,
		},
	}

	err = usecase.codeRepo.AddCode(ctx, fullCode)
	if err != nil {
		return err
	}

	return nil
}
