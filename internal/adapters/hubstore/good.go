package hubstore

import (
	"context"
	"fmt"
	"molocode/internal/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
)

// Добавляет продукт в хранилище, возвращает все поля добавленного продукта
func (hs *hubStorage) AddGood(g entity.Good) error {
	const op = "hubstorage.AddGood"
	// MAPPING
	mappedGood := Good_dto{
		Gtin:            string(g.Gtin),
		Desc:            g.Desc,
		StoreCount:      g.StoreCount,
		GetCodeForPrint: g.GetCodeForPrint,
		AllowProduce:    g.AllowProduce,
		Upload:          g.Upload,
		CreateAt:        g.CreatedAt,
	}
	_, err := hs.db.Collection(collectionGoods).InsertOne(context.TODO(), mappedGood)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (hs *hubStorage) GetGood(gtin entity.Gtin) (entity.Good, error) {
	const op = "hubservice.GetGood"

	filter := bson.M{"_id": gtin}
	reqResult := hs.db.Collection(collectionGoods).FindOne(context.TODO(), filter)

	var result entity.Good
	err := reqResult.Decode(&result)
	if err != nil {
		return entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}
	if result.Gtin == "" {
		return entity.Good{}, fmt.Errorf("%s: Продукт не найден", op)
	}
	return result, nil
}

func (hs *hubStorage) GetAllGoods() ([]entity.Good, error) {
	const op = "hubservice.GetAllGoods"

	filter := bson.M{}
	cursor, err := hs.db.Collection(collectionGoods).Find(context.TODO(), filter)
	if err != nil {
		return []entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}
	result := []entity.Good{}
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		return []entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}
