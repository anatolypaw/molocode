package mongostore

import (
	"context"
	"fmt"
	"molocode/internal/app/entity"

	"go.mongodb.org/mongo-driver/bson"
)

func (ths *MongoStore) AddCode(ctx context.Context, code entity.FullCode) error {
	const op = "hubstorage.AddCode"

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
func (ths *MongoStore) GetCountPrintAvaibleCode(ctx context.Context, gtin string) (uint, error) {
	filter := bson.M{"printinfo.avaible": true}
	avaible, err := ths.db.Collection(gtin).CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return uint(avaible), err
}
