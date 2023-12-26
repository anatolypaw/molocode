package storage

import (
	"context"
	"fmt"
	"storage/model"

	"go.mongodb.org/mongo-driver/bson"
)

// Возвращает информацию о продукте. Если gtin пустой, возвращает все продукты
func (con *Storage) GetGood(gtin string) (model.Good, error) {
	const op = "storage.GetGood"

	// Валидация gtin
	if err := model.ValidateGtin(gtin); err != nil {
		return model.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	filter := bson.M{"_id": gtin}

	reqResult := con.db.Collection(collectionGoods).FindOne(context.TODO(), filter)

	var result model.Good
	err := reqResult.Decode(&result)
	if err != nil {
		return model.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	if result.Gtin == "" {
		return model.Good{}, fmt.Errorf("%s: Продукт не найден", op)
	}

	return result, nil
}
