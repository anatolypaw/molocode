package storage

import (
	"fmt"
)

// Продукт, gtin для каждого уникален. 14 символов
type Good struct {
	Gtin       string `bson:"gtin,omitempty"`        // gtin продукта
	Desc       string `bson:"desc,omitempty"`        // описание продукта
	StoreCount int    `bson:"store_count,omitempty"` // сколько хранить кодов
	Get        bool   `bson:"get,omitempty"`         // флаг, получать коды из 1с
	Upload     bool   `bson:"upload,omitempty"`      // флаг, выгружать коды в 1с
	Avaible    bool   `bson:"avaible,omitempty"`     // флаг, выдавать ли кода на терминал
	ShelfLife  int    `bson:"shelf_life,omitempty"`  // срок годности продукта
}

// Добавляет продукт в хранилище
func (con *Connection) AddGood(gtin string, desc string) error {
	const op = "storage.AddGood"

	g := Good{Gtin: gtin, Desc: desc}

	_, err := con.db.Collection("goods").InsertOne(*con.ctx, g)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
