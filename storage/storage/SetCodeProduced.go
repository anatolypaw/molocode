package storage

import (
	"context"
	"fmt"
	"storage/model"

	"go.mongodb.org/mongo-driver/bson"
)

// Устанавливает информацию о производстве кода
// Добавляет код, если его нет.
func (con *Storage) SetCodeProduced(gtin, serial, crypto, sourceName string, discard bool) (model.Code, error) {
	const op = "storage.mongodb.SetCodeProduced"

	// Валидируем данные о коде
	if err := model.ValidateSerial(serial); err != nil {
		return model.Code{}, fmt.Errorf("%s: %w", op, err)
	}

	if err := model.ValidateCrypto(crypto); err != nil {
		return model.Code{}, fmt.Errorf("%s: %w", op, err)
	}

	// Проверяем, существует ли такой продукт в БД
	good, err := con.GetGood(gtin)
	if err != nil {
		return model.Code{}, fmt.Errorf("%s: %w", op, err)
	}

	// Проверяем, разрешено ли производство этого продукта
	if !good.AllowProduce {
		return model.Code{}, fmt.Errorf("%s: %s", op, "Для этого продукта запрещено производство")
	}

	// Получаем информацию о коде из бд.
	filter := bson.M{"_id": serial}
	reqResult := con.db.Collection(gtin).FindOne(context.TODO(), filter)

	var code model.Code
	reqResult.Decode(&code)

	if code.Serial == "" {
		return model.Code{}, fmt.Errorf("%s: код не найден", op)
	}

	return code, nil
}
