package mongo

import (
	"context"
	"fmt"
	"molocode/internal/entity"
)

func (hs *Store) AddCode(code entity.FullCode) error {
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

	_, err := hs.db.Collection(string(code.Gtin)).InsertOne(context.TODO(), mappedCode)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return err
}
