package mongodb

import (
	"context"
	"fmt"
	"storage/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Добавляет код к указанному по gtin продукту
func (con *Storage) AddCode(gtin string,
	serial string,
	crypto string,
	sourceName string) error {
	const op = "storage.mongodb.AddCode"

	// Валидация входных данных
	if err := model.ValidateGtin(gtin); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := model.ValidateSerial(serial); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := model.ValidateCrypto(crypto); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	var newCode model.Code
	newCode.Serial = serial
	newCode.Crypto = crypto
	newCode.SourceInfo.Name = sourceName
	newCode.SourceInfo.Time = time.Now()

	filter := bson.M{"_id": gtin}
	update := bson.M{"$push": bson.M{"codes": newCode}}

	result, err := con.db.Collection(collectionGoods).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	// Не найдено совпадений с указанным gtin
	if result.MatchedCount != 1 {
		return fmt.Errorf("%s: Не найден продукт с GTIN %s", op, gtin)
	}

	return err
}
