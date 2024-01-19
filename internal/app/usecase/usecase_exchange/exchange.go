package usecase_exchange

import (
	"context"
	"molocode/internal/app/entity"
	"molocode/internal/app/repository"
)

type ExchangeUsecase struct {
	goodRepository repository.IGoodRepository
	codeRepository repository.ICodeRepository
}

func New(goodRepository repository.IGoodRepository, codeRepository repository.ICodeRepository) ExchangeUsecase {
	return ExchangeUsecase{
		goodRepository: goodRepository,
		codeRepository: codeRepository,
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
func (eu *ExchangeUsecase) GetGoodsReqCodes(ctx context.Context) ([]CodeReq, error) {
	// - Получить продукты
	allGoods, err := eu.goodRepository.GetAllGoods(ctx)
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
		avaibleCount, err := eu.codeRepository.GetCountPrintAvaible(ctx, good.Gtin)
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
