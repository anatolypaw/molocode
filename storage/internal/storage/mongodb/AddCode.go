package mongodb

import (
	"fmt"
	"storage/internal/domain/models"
	"time"
)

// Добавляет код к продукту
func (con *Storage) AddCode(gtin string, code models.Code) (models.Code, error) {
	const op = "storage.mongodb.AddCode"

	// Валидация входных данных
	err := models.ValidateGtin(gtin)
	if err != nil {
		return models.Code{}, fmt.Errorf("%s: %w", op, err)
	}

	err = code.Validate()
	if err != nil {
		return models.Code{}, fmt.Errorf("%s: %w", op, err)
	}

	code.Loaded.Time = time.Now()
	/*
		// Добавляем продукт в БД
		objID, err := con.db.Collection(goodCollection).InsertOne(*con.ctx, code)
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
	*/
	var res models.Code
	return res, nil
}
