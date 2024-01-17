package mongo

import (
	"context"
	"fmt"
	"molocode/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
)

func (ths *Store) AddCode(code entity.FullCode) error {
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

	_, err := ths.db.Collection(string(code.Gtin)).InsertOne(ths.ctx, mappedCode)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}

// TODO в случае изменения entity.Code, может перестать выполняться запрос
func (ths *Store) CountPrintAvaibleCode(gtin string) (uint, error) {
	filter := bson.M{"printinfo.avaible": true}
	avaible, err := ths.db.Collection(gtin).CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return uint(avaible), err
}
