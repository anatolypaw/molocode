package storage

import (
	"fmt"
	"storage/internal/entity"
)

// Добавляет продукт в хранилище
func (con *Connection) AddGood(gtin string, desc string) error {
	const op = "storage.AddGood"

	g := entity.Good{Gtin: gtin, Description: desc}

	_, err := con.db.Collection("goods").InsertOne(*con.ctx, g)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
