package mongodb

import (
	"context"
	"fmt"
	"storage/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Добавляет продукт в хранилище, возвращает все поля добавленного продукта
func (con *Storage) AddGood(good model.Good) (model.Good, error) {
	const op = "storage.mongodb.AddGood"

	err := good.Validate()
	if err != nil {
		return model.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	good.Created = time.Now()
	good.Codes = []model.Code{}

	// Добавляем продукт в БД
	objID, err := con.db.Collection(collectionGoods).InsertOne(context.TODO(), good)
	if err != nil {
		return model.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	// Считываем с базы, что мы там записали, кроме массива кодов
	filter := bson.D{{Key: "_id", Value: objID.InsertedID}}
	opt := options.FindOne().SetProjection(bson.D{{Key: "codes", Value: 0}})

	var res model.Good
	err = con.db.Collection(collectionGoods).FindOne(context.TODO(), filter, opt).Decode(&res)
	if err != nil {
		return model.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	return res, nil
}
