package storage

import (
	"fmt"
	"storage/model"
)

// Возвращает код к указанному gtin продукту для последующей печати
func (con *Storage) GetReqCodeCount() ([]model.CodeReq, error) {
	const op = "storage.GetReqCodeCount"

	// Получаем все продукты
	goods, err := con.GetGoods()
	if err != nil {
		return []model.CodeReq{}, fmt.Errorf("%s: %w", op, err)
	}

	// Для продуктов, у которых разрешено получение кодов, подсчитываем, сколько
	// кодов доступно для печати.
	// если доступных меньше, чем должно храниться, добавляем этот продукт
	// и требуемое количество кодов в реузльтат
	result := []model.CodeReq{}

	for _, good := range goods {
		// Не требуется пополнение кодами, пропускаем
		if !good.GetCodeForPrint {
			continue
		}

		printAvaible, err := con.GetPrintAvaibleCountCode(good.Gtin)
		if err != nil {
			return []model.CodeReq{}, fmt.Errorf("%s: %w", op, err)
		}

		req := model.CodeReq{
			Gtin:          good.Gtin,
			RequiredCount: int64(good.StoreCount) - printAvaible,
		}

		if !good.GetCodeForPrint {
			req.RequiredCount = 0
		}

		result = append(result, req)
	}
	return result, err
}
