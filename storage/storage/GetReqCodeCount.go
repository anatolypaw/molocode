package storage

import (
	"fmt"
	"storage/model"
)

// Возвращает код к указанному gtin продукту для последующей печати
func (con *Storage) GetReqCodeCount() ([]model.Good, error) {
	const op = "storage.GetReqCodeCount"

	goods, err := con.GetGoods()
	if err != nil {
		return []model.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	return goods, err
}
