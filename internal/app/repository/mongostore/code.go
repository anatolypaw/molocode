package mongostore

import (
	"context"
	"fmt"
	"molocode/internal/app/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Добавляет код как есть
func (ths *MongoStore) AddCode(ctx context.Context, code entity.FullCode) error {
	const op = "mongostore.AddCode"

	// MAPPING
	mappedCode := Code_dto{
		Serial:       string(code.Serial),
		Crypto:       string(code.Crypto),
		SourceInfo:   code.SourceInfo,
		PrintInfo:    code.PrintInfo,
		ProducedInfo: code.ProducedInfo,
		UploadInfo:   code.UploadInfo,
	}

	_, err := ths.db.Collection(string(code.Gtin)).InsertOne(context.TODO(), mappedCode)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}

// TODO в случае изменения поля printinfo entity.Code, может перестать выполняться запрос
// Можно решить полнным маппингом структуры кода
func (ths *MongoStore) GetCountPrintAvaible(ctx context.Context, gtin string) (uint, error) {
	filter := bson.M{"printinfo.avaible": true}
	avaible, err := ths.db.Collection(gtin).CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return uint(avaible), err
}

// Возвращает код для печати, увеличивает счетчик кодов
func (ths *MongoStore) GetCodeForPrint(ctx context.Context, gtin string, terminalName string) (entity.CodeForPrint, error) {

	// Получаем код, пригодный для печати, ставим в бд флаг, что он больше не доступен для печати, что бы заблокировать
	// возможность получения этого кода в другом потоке
	filter := bson.M{"printinfo.avaible": true}
	update := bson.M{"$set": bson.M{"printinfo.avaible": false, "printinfo.terminalname": terminalName, "printinfo.uploadtime": time.Now()}}

	var code Code_dto
	err := ths.db.Collection(gtin).FindOneAndUpdate(ctx, filter, update).Decode(&code)
	if err != nil {
		return entity.CodeForPrint{}, fmt.Errorf("получение доступного кода для печати: %s", err)
	}

	// Получаем PrintId для этого кода, инкрементируем счетчик кодов
	filter = bson.M{"_id": "nextprintid"}
	update = bson.M{"$inc": bson.M{"value": 1}}
	opt := options.FindOneAndUpdate().SetUpsert(true)

	var printId Counters
	err = ths.db.Collection(collectionCounters).FindOneAndUpdate(ctx, filter, update, opt).Decode(&printId)
	if err != nil {
		return entity.CodeForPrint{}, fmt.Errorf("ошибка инкремента nextprintid %s", err)
	}

	// Присваиваем коду PrintID
	filter = bson.M{"_id": code.Serial}
	update = bson.M{"$set": bson.M{"printinfo.printid": printId.Value}}
	updResult, err := ths.db.Collection(gtin).UpdateOne(ctx, filter, update)
	if err != nil {
		return entity.CodeForPrint{}, fmt.Errorf("присввоение коду printId: %s", err)
	}

	// Когда PrintID == 0, то обновления происходить не будет, и будет ложная ошибка, по этому
	// ошибку для случая с ID = 0 пропускаем
	if printId.Value > 0 && updResult.ModifiedCount != 1 {
		return entity.CodeForPrint{}, fmt.Errorf("ошибка установки printID для кода GTIN: %s serial: %s", gtin, code.Serial)
	}

	// Приводим к нужной структуре
	codeForPrint := entity.CodeForPrint{
		Code: entity.Code{
			Gtin:   gtin,
			Serial: code.Serial,
			Crypto: code.Crypto,
		},
		PrintId: printId.Value,
	}

	return codeForPrint, nil
}
