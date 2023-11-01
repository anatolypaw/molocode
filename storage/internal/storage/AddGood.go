package storage

import (
	"fmt"
	"storage/internal/entity"
)

// Добавляет продукт в хранилище
func (con *Connection) AddGood(good entity.Good) error {
	const op = "storage.AddGood"

	err := good.ValidateGtin()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = con.db.Collection("goods").InsertOne(*con.ctx, good)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
