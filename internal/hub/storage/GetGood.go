package storage

import (
	"context"
	"fmt"
	"molocode/entity"

	"go.mongodb.org/mongo-driver/bson"
)

// Возвращает информацию о продукте. Если gtin пустой, возвращает все продукты
func (con *Storage) GetGood(gtin string) (entity.Good, error) {
	const op = "storage.GetGood"

	// Валидация gtin
	if err := entity.ValidateGtin(gtin); err != nil {
		return entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	filter := bson.M{"_id": gtin}

	reqResult := con.db.Collection(collectionGoods).FindOne(context.TODO(), filter)

	var result entity.Good
	err := reqResult.Decode(&result)
	if err != nil {
		return entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	if result.Gtin == "" {
		return entity.Good{}, fmt.Errorf("%s: Продукт не найден", op)
	}

	return result, nil
}
