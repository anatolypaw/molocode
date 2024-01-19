package usecase_exchange

import (
	"molocode/internal/app/entity"
	"molocode/internal/app/storage"
)

type Usecase struct {
	goodRepo storage.IGood
	codeRepo storage.ICode
}

func New(goodRepo storage.IGood, codeRepo storage.ICode) Usecase {
	return Usecase{
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
func (ths *Usecase) GetGoodsReqCodes() ([]CodeReq, error) {
	// - Получить продукты
	allGoods, err := ths.goodRepo.GetAllGoods()
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
		avaibleCount, err := ths.codeRepo.GetCountPrintAvaible(good.Gtin)
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
