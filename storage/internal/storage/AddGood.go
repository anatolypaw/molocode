package storage

import (
	"context"
	"fmt"
	"storage/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Добавляет продукт в хранилище, проверяя корректность gtin
func (con *Connection) AddGood(good entity.Good) (entity.Good, error) {
	const op = "storage.AddGood"
	err := good.ValidateGtin()
	if err != nil {
		return entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	// Добавляем продукт в БД
	objID, err := con.db.Collection(goodCollection).InsertOne(*con.ctx, good)
	if err != nil {
		return entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	fmt.Println(objID.InsertedID)

	// Считываем с базы, что мы там записали, кроме массива кодов
	filter := bson.D{{Key: "_id", Value: objID.InsertedID}}
	opt := options.FindOne().SetProjection(bson.D{{Key: "codes", Value: 0}})

	var res entity.Good
	err = con.db.Collection(goodCollection).FindOne(context.TODO(), filter, opt).Decode(&res)
	if err != nil {
		return entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	return res, nil
}
