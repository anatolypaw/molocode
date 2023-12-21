package mongodb

import (
	"context"
	"fmt"
	"storage/model"
	"time"
)

// Добавляет код к указанному по gtin продукту для последующей печати
// Добавляет код, только если в свойствах этого продукта разрешено получение кодов для нанесения
func (con *Storage) AddCodeForPrint(gtin, serial, crypto, sourceName string) error {
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

	// Проверяем, существует ли такой продукт в БД
	goods, err := con.GetGoods(gtin)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if len(goods) != 1 {
		return fmt.Errorf("%s: %s", op, "Продукт не найден")
	}
	good := goods[0]

	// Проверяем, разрешено ли получение кодов для этого продукта
	if !good.AcceptForPrint {
		return fmt.Errorf("%s: %s", op, "Для этого продукта запрещено получение кодов для нанесения")
	}

	var newCode model.Code
	newCode.Serial = serial
	newCode.Crypto = crypto
	newCode.SourceInfo.Name = sourceName
	newCode.SourceInfo.Time = time.Now()

	_, err = con.db.Collection(gtin).InsertOne(context.TODO(), newCode)
	if err != nil {
		return err
	}

	return err
}