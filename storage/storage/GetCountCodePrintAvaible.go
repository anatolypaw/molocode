package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// Возвращает количество доступных для печати кодов
func (con *Storage) GetCountCodePrintAvaible(gtin string) (int64, error) {
	const op = "storage.GetCountCodePrintAvaible"

	filter := bson.M{"PrintInfo.Avaible": true}
	avaible, err := con.db.Collection(gtin).CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return avaible, err
}
