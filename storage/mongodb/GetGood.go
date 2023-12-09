package mongodb

import (
	"context"
	"fmt"
	"storage/model"

	"go.mongodb.org/mongo-driver/bson"
)

// Возвращает все продукты и информацию по ним
func (con *Storage) GetGood() ([]model.Good, error) {
	const op = "storage.mongodb.GetGood"

	filter := bson.D{{}}
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
