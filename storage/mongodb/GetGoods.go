package mongodb

import (
	"context"
	"fmt"
	"storage/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Возвращает информацию о продукте. Если gtin пустой, возвращает все продукты
func (con *Storage) GetGoods(gtin string) ([]model.Good, error) {
	const op = "storage.mongodb.GetGoods"

	var filter primitive.M

	if gtin == "" {
		filter = bson.M{}
	} else {
		// Валидация gtin
		if err := model.ValidateGtin(gtin); err != nil {
			return []model.Good{}, fmt.Errorf("%s: %w", op, err)
		}
		filter = bson.M{"_id": gtin}
	}

	cursor, err := con.db.Collection(collectionGoods).Find(context.TODO(), filter)
	if err != nil {
		return []model.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	var result []model.Good

	err = cursor.All(context.TODO(), &result)
	if err != nil {
		return []model.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
}
