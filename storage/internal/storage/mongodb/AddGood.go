package mongodb

import (
	"context"
	"fmt"
	"storage/internal/domain/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Добавляет продукт в хранилище, проверяя корректность gtin
func (con *Storage) AddGood(good models.Good) (models.Good, error) {
	const op = "storage.mongodb.AddGood"

	good.Created = time.Now()

	if good.Codes != nil {
		return models.Good{}, fmt.Errorf("%s: Недопустимо добавление кодов", op)
	}

	err := good.ValidateGtin()
	if err != nil {
		return models.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	// Добавляем продукт в БД
	objID, err := con.db.Collection(goodCollection).InsertOne(*con.ctx, good)
	if err != nil {
		return models.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	// Считываем с базы, что мы там записали, кроме массива кодов
	filter := bson.D{{Key: "_id", Value: objID.InsertedID}}
	opt := options.FindOne().SetProjection(bson.D{{Key: "codes", Value: 0}})

	var res models.Good
	err = con.db.Collection(goodCollection).FindOne(context.TODO(), filter, opt).Decode(&res)
	if err != nil {
		return models.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	return res, nil
}
