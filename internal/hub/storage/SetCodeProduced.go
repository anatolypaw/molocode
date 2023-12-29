package storage

import (
	"context"
	"fmt"
	"molocode/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Устанавливает информацию о производстве кода
// Добавляет код, если его нет.
func (con *Storage) SetCodeProduced(gtin, serial, crypto, terminal, proddate string, discard bool) error {
	const op = "storage.SetCodeProduced"

	// Валидируем данные о коде
	if err := entity.ValidateSerialCrypto(serial, crypto); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Проверка корректности формата даты
	// TODO

	// Проверяем, существует ли такой продукт в БД
	good, err := con.GetGood(gtin)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Проверяем, разрешено ли производство этого продукта
	if !good.AllowProduce {
		return fmt.Errorf("%s: %s", op, "Для этого продукта не разрешено производство")
	}

	// Получаем информацию о коде из бд.
	filter := bson.M{"_id": serial}
	reqResult := con.db.Collection(gtin).FindOne(context.TODO(), filter)
	var code entity.Code
	reqResult.Decode(&code)

	if code.Serial == "" {
		return fmt.Errorf("%s: код GTIN: %s SN: %s не найден", op, gtin, serial)
	}

	// Если код уже есть в бд, проверяем, отбракован ли или произведен он
	if len(code.ProducedInfo) > 0 {
		if code.ProducedInfo[len(code.ProducedInfo)-1].Discard == discard {
			if discard {
				return fmt.Errorf("%s: код уже отбракован", op)
			} else {
				return fmt.Errorf("%s: код уже был произведен", op)
			}
		}
	}

	prodInfo := []entity.ProducedInfo{
		{
			Discard:  discard,
			ProdDate: proddate,
			Terminal: terminal,
			Time:     time.Now(),
		},
	}

	// Добавляем к коду информацию о производстве
	filter = bson.M{"_id": serial}
	update := bson.M{"$addToSet": bson.M{"ProducedInfo": bson.M{"$each": prodInfo}}}
	updResult, err := con.db.Collection(gtin).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if updResult.ModifiedCount != 1 {
		return fmt.Errorf("%s: Ошибка добавления информации о производстве GTIN: %s serial: %s", op, gtin, serial)
	}

	return nil
}
