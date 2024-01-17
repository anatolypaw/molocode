package usecase_exchange

import (
	service "molocode/internal/domain/service/store"
)

type UseCase struct {
	store *service.Store
}

func New(storeService *service.Store) UseCase {
	return UseCase{store: storeService}
}

// Отвечает за добавление новых кодов для печати
// и выдачу кодов на выгрузку
type CodeReq struct {
	Gtin     string
	Desc     string
	Required uint
}

// Возвращает список продуктов, требующих наполнения кодами для печати
// и количество требуемых кодов
func (usecase *UseCase) GetGoodsReqCodes() ([]CodeReq, error) {
	// - Получить продукты, для которых включено наполнение кодами
	goods, err := usecase.store.GetGoodsForPrint()
	if err != nil {
		return nil, err
	}

	var codesReq []CodeReq
	for _, good := range goods {
		// - Для каждого продукта получить доступное количество кодов
		avaible, err := usecase.store.CountPrintAvaibleCode(good.Gtin)
		if err != nil {
			return nil, err
		}

		requiredCount := good.StoreCount - avaible
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
