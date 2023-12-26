package storage

import (
	"context"
	"fmt"
	"storage/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Возвращает код к указанному gtin продукту для последующей печати
func (con *Storage) GetCodeForPrint(gtin, terminal string) (model.CodeForPrint, error) {
	const op = "storage.GetCodeCodeForPrint"

	// Проверка имени терминала
	if len(terminal) == 0 {
		return model.CodeForPrint{}, fmt.Errorf("%s: %s", op, "Не указано имя терминала")
	}

	// Проверяем, существует ли такой продукт в БД
	good, err := con.GetGood(gtin)
	if err != nil {
		return model.CodeForPrint{}, fmt.Errorf("%s: %w", op, err)
	}

	// Проверяем, разрешена ли для этого продукта выдача кодов для печати
	if !good.SendForPrint {
		return model.CodeForPrint{}, fmt.Errorf("%s: %s", op, "Для этого продукта запрещена выдача кодов для нанесения")
	}

	// Получаем код, пригодный для печати, ставим флаг в бд, что код получен, что бы заблокировать
	// возможность получения этого кода в другом потоке
	filter := bson.M{"printinfo.uploaded": false}
	update := bson.M{"$set": bson.M{"printinfo.uploaded": true, "printinfo.terminalname": terminal, "printinfo.uploadtime": time.Now()}}
	reqResult := con.db.Collection(gtin).FindOneAndUpdate(context.TODO(), filter, update)

	var code model.Code
	reqResult.Decode(&code)

	if code.Serial == "" {
		return model.CodeForPrint{}, fmt.Errorf("%s: Для GTIN %s нет кодов для печати", op, gtin)
	}

	// Получаем PrintID для этого кода, инкрементируем счетчик кодов
	filter = bson.M{"name": "NextPrintID"}
	update = bson.M{"$inc": bson.M{"value": 1}}
	opt := options.FindOneAndUpdate().SetUpsert(true)
	reqResult = con.db.Collection(collectionCounters).FindOneAndUpdate(context.TODO(), filter, update, opt)

	var printID model.Counters
	reqResult.Decode(&printID)

	// Присваиваем коду PrintID
	filter = bson.M{"_id": code.Serial}
	update = bson.M{"$set": bson.M{"printinfo.printid": printID.Value}}
	updResult, err := con.db.Collection(gtin).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return model.CodeForPrint{}, fmt.Errorf("%s: %w", op, err)
	}

	if updResult.ModifiedCount != 1 {
		return model.CodeForPrint{}, fmt.Errorf("%s: Ошибка установки PrintID для кода GTIN: %s serial: %s", op, gtin, code.Serial)
	}

	// Приводим к нужной структуре
	var result model.CodeForPrint
	result.Gtin = gtin
	result.Serial = code.Serial
	result.Crypto = code.Crypto
	result.PrintId = printID.Value

	return result, err
}
