package mongostore

import (
	"context"
	"fmt"
	"hub/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Добавляет код как есть
func (ths *MongoStore) AddCode(ctx context.Context, code entity.FullCode) error {
	const op = "mongostore.AddCode"

	// MAPPING
	mappedCode := FullCode_dto{
		Serial:       string(code.Serial),
		Crypto:       string(code.Crypto),
		SourceInfo:   code.SourceInfo,
		PrintInfo:    code.PrintInfo,
		ProducedInfo: code.ProducedInfo,
		UploadInfo:   code.UploadInfo,
	}

	codes := ths.db.Collection(string(code.Gtin))
	_, err := codes.InsertOne(ctx, mappedCode)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}

// TODO в случае изменения поля printinfo entity.Code,
// может перестать выполняться запрос
// Можно решить полнным маппингом структуры кода
func (ths *MongoStore) GetCountPrintAvaible(ctx context.Context, gtin string,
) (uint, error) {
	filter := bson.M{"printinfo.avaible": true}
	codes := ths.db.Collection(gtin)
	avaible, err := codes.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return uint(avaible), err
}

// Возвращает код для печати, увеличивает счетчик кодов
func (ths *MongoStore) GetCodeForPrint(
	ctx context.Context,
	gtin string,
	terminalName string,
) (entity.CodeForPrint, error) {

	// Получаем код, пригодный для печати, ставим в бд флаг,
	// что он больше не доступен для печати, что бы заблокировать
	// возможность получения этого кода в другом потоке
	filter := bson.M{"printinfo.avaible": true}
	update := bson.M{"$set": bson.M{"printinfo.avaible": false,
		"printinfo.terminalname": terminalName,
		"printinfo.uploadtime":   time.Now()}}

	var code FullCode_dto
	codes := ths.db.Collection(gtin)
	err := codes.FindOneAndUpdate(ctx, filter, update).Decode(&code)
	if err != nil {
		return entity.CodeForPrint{},
			fmt.Errorf("получение доступного кода для печати: %s", err)
	}

	// Получаем PrintId для этого кода, инкрементируем счетчик кодов
	filter = bson.M{"_id": "nextprintid"}
	update = bson.M{"$inc": bson.M{"value": 1}}
	opt := options.FindOneAndUpdate().SetUpsert(true)

	var printId Counters
	counters := ths.db.Collection(COLLECTION_COUNTERS)
	res := counters.FindOneAndUpdate(ctx, filter, update, opt)
	err = res.Decode(&printId)
	if err != nil {
		return entity.CodeForPrint{},
			fmt.Errorf("ошибка инкремента nextprintid %s", err)
	}

	// Присваиваем коду PrintID
	filter = bson.M{"_id": code.Serial}
	update = bson.M{"$set": bson.M{"printinfo.printid": printId.Value}}
	updResult, err := ths.db.Collection(gtin).UpdateOne(ctx, filter, update)
	if err != nil {
		return entity.CodeForPrint{},
			fmt.Errorf("присввоение коду printId: %s", err)
	}

	// Когда PrintID == 0, то обновления происходить не будет,
	// и будет ложная ошибка, по этому
	// ошибку для случая с ID = 0 пропускаем
	if printId.Value > 0 && updResult.ModifiedCount != 1 {
		return entity.CodeForPrint{},
			fmt.Errorf("ошибка установки printID для кода GTIN: %s serial: %s",
				gtin, code.Serial)
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

// Возвращает код
func (ths *MongoStore) GetCode(
	ctx context.Context,
	gtin string,
	serial string,
) (entity.FullCode, error) {
	const op = "mongo.GetCode"

	filter := bson.M{"_id": serial}
	codes := ths.db.Collection(gtin)
	reqResult := codes.FindOne(ctx, filter)

	var code FullCode_dto
	err := reqResult.Decode(&code)
	if err != nil {
		return entity.FullCode{}, fmt.Errorf("%s: %w", op, err)
	}
	if code.Serial == "" {
		return entity.FullCode{}, fmt.Errorf("%s: Продукт не найден", op)
	}

	// MAPPING
	mappedCode := entity.FullCode{
		Code: entity.Code{
			Gtin:   gtin,
			Serial: code.Serial,
			Crypto: code.Crypto,
		},
		SourceInfo:   code.SourceInfo,
		PrintInfo:    code.PrintInfo,
		ProducedInfo: code.ProducedInfo,
		UploadInfo:   code.UploadInfo,
	}

	return mappedCode, nil

}
