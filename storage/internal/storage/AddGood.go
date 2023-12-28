package storage

import (
	"context"
	"fmt"
	"storage/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Добавляет продукт в хранилище, возвращает все поля добавленного продукта
func (con *Storage) AddGood(good entity.Good) (entity.Good, error) {
	const op = "storage.AddGood"

	err := good.Validate()
	if err != nil {
		return entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	good.Created = time.Now()

	// Добавляем продукт в БД
	objID, err := con.db.Collection(collectionGoods).InsertOne(context.TODO(), good)
	if err != nil {
		return entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	// Считываем с базы, что мы там записали
	filter := bson.D{{Key: "_id", Value: objID.InsertedID}}

	var res entity.Good
	err = con.db.Collection(collectionGoods).FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	return res, nil
}
