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

	// Получаем код, пригодный для печати, меняем его доступность, что бы заблокировать
	// возможность получения этого кода в другом потоке
	filter := bson.M{"PrintInfo.Avaible": true}
	update := bson.M{"$set": bson.M{"PrintInfo.Avaible": false, "PrintInfo.TerminalName": terminal, "PrintInfo.UploadTime": time.Now()}}
	reqResult := con.db.Collection(gtin).FindOneAndUpdate(context.TODO(), filter, update)

	var code model.Code
	reqResult.Decode(&code)

	if code.Serial == "" {
		return model.CodeForPrint{}, fmt.Errorf("%s: Для GTIN %s нет кодов для печати", op, gtin)
	}

	// Получаем PrintID для этого кода, инкрементируем счетчик кодов
	filter = bson.M{"_id": "NextPrintID"}
	update = bson.M{"$inc": bson.M{"Value": 1}}
	opt := options.FindOneAndUpdate().SetUpsert(true)
	reqResult = con.db.Collection(collectionCounters).FindOneAndUpdate(context.TODO(), filter, update, opt)

	var printID model.Counters
	reqResult.Decode(&printID)

	// Присваиваем коду PrintID
	filter = bson.M{"_id": code.Serial}
	update = bson.M{"$set": bson.M{"PrintInfo.PrintID": printID.Value}}
	updResult, err := con.db.Collection(gtin).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return model.CodeForPrint{}, fmt.Errorf("%s: %w", op, err)
	}

	// Когда PrintID == 0, то обновления происходить не будет, и будет ложная ошибка, по этому
	// учитываем этот момент
	if printID.Value > 0 && updResult.ModifiedCount != 1 {
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
