package usecase_exchange

type iStorage interface {
}

type UseCase struct {
	store iStorage
}

func New(storeService iStorage) UseCase {
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
func (ths *UseCase) GetGoodsReqCodes() ([]CodeReq, error) {
	// - Получить продукты, для которых включено наполнение кодами
	goods, err := ths.store.GetGoodsForPrint()
	if err != nil {
		return nil, err
	}

	var codesReq []CodeReq
	for _, good := range goods {
		// - Для каждого продукта получить доступное количество кодов
		avaible, err := ths.store.CountPrintAvaibleCode(good.Gtin)
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
