package storage

import (
	"context"
	"fmt"
	"molocode/entity"

	"go.mongodb.org/mongo-driver/bson"
)

// Возвращает информацию о продукте. Если gtin пустой, возвращает все продукты
func (con *Storage) GetGoods() ([]entity.Good, error) {
	const op = "storage.GetGoods"

	filter := bson.M{}

	cursor, err := con.db.Collection(collectionGoods).Find(context.TODO(), filter)
	if err != nil {
		return []entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	result := []entity.Good{}

	err = cursor.All(context.TODO(), &result)
	if err != nil {
		return []entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
}
