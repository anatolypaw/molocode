package storage

import (
	"context"
	"fmt"
	"storage/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Добавляет продукт в хранилище, возвращает все поля добавленного продукта
func (con *Storage) AddGood(good model.Good) (model.Good, error) {
	const op = "storage.mongodb.AddGood"

	err := good.Validate()
	if err != nil {
		return model.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	good.Created = time.Now()

	// Добавляем продукт в БД
	objID, err := con.db.Collection(collectionGoods).InsertOne(context.TODO(), good)
	if err != nil {
		return model.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	// Считываем с базы, что мы там записали
	filter := bson.D{{Key: "_id", Value: objID.InsertedID}}

	var res model.Good
	err = con.db.Collection(collectionGoods).FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return model.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	return res, nil
}
