package storage

import "fmt"

// Продукт, gtin для каждого уникален. 14 символов
type Good struct {
	Gtin       string // gtin продукта
	Desc       string // описание продукта
	StoreCount int    // сколько хранить кодов
	Get        bool   // флаг, получать коды из 1с
	Upload     bool   // флаг, выгружать коды в 1с
	Awaible    bool   // флаг, выдавать ли кода на терминал
	ShelfLife  int    // срок годности продукта
}

// Добавляет продукт в хранилище
func (m *mongodb) AddGood(gtin string, desc string) error {
	const op = "storage.AddGood"

	g := Good{
		Gtin: gtin,
		Desc: desc,
	}

	_, err := m.db.Collection("goods").InsertOne(*m.ctx, g)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}