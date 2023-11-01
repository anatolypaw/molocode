package storage

import (
	"fmt"
	"storage/internal/entity"
)

// Добавляет продукт в хранилище
func (con *Connection) EditGood(good *entity.Good) error {
	const op = "storage.EditGood"

	_, err := con.db.Collection("goods").InsertOne(*con.ctx, good)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
