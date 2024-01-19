package mongostore

import (
	"context"
	"fmt"
	"molocode/internal/app/entity"

	"go.mongodb.org/mongo-driver/bson"
)

// Добавляет продукт в хранилище, возвращает все поля добавленного продукта
func (ths *MongoStore) AddGood(g entity.Good) error {
	const op = "mongo.AddGood"
	// MAPPING
	mappedGood := Good_dto{
		Gtin:            g.Gtin,
		Desc:            g.Desc,
		StoreCount:      g.StoreCount,
		GetCodeForPrint: g.GetCodeForPrint,
		AllowProduce:    g.AllowProduce,
		Upload:          g.Upload,
		CreatedAt:       g.CreatedAt,
	}
	insertResult, err := ths.db.Collection(collectionGoods).InsertOne(context.TODO(), mappedGood)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	fmt.Printf("%s: %#v", op, insertResult)

	return nil
}

func (ths *MongoStore) GetGood(gtin string) (entity.Good, error) {
	const op = "mongo.GetGood"

	filter := bson.M{"_id": gtin}
	reqResult := ths.db.Collection(collectionGoods).FindOne(context.TODO(), filter)

	var good_dto Good_dto
	err := reqResult.Decode(&good_dto)
	if err != nil {
		return entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}
	if good_dto.Gtin == "" {
		return entity.Good{}, fmt.Errorf("%s: Продукт не найден", op)
	}

	// MAPPING
	mappedGood := entity.Good{
		Gtin:            good_dto.Gtin,
		Desc:            good_dto.Desc,
		StoreCount:      good_dto.StoreCount,
		GetCodeForPrint: good_dto.GetCodeForPrint,
		AllowProduce:    good_dto.AllowProduce,
		Upload:          good_dto.Upload,
		CreatedAt:       good_dto.CreatedAt,
	}

	return mappedGood, nil
}

func (ths *MongoStore) GetAllGoods() ([]entity.Good, error) {
	const op = "mongo.GetAllGoods"

	filter := bson.M{}
	cursor, err := ths.db.Collection(collectionGoods).Find(context.TODO(), filter)
	if err != nil {
		return []entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	goods_dto := []Good_dto{}
	err = cursor.All(context.TODO(), &goods_dto)
	if err != nil {
		return []entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	// MAPPING
	mappedGoods := []entity.Good{}
	for _, good_dto := range goods_dto {
		mappedGoods = append(mappedGoods, entity.Good{
			Gtin:            good_dto.Gtin,
			Desc:            good_dto.Desc,
			StoreCount:      good_dto.StoreCount,
			GetCodeForPrint: good_dto.GetCodeForPrint,
			AllowProduce:    good_dto.AllowProduce,
			Upload:          good_dto.Upload,
			CreatedAt:       good_dto.CreatedAt,
		})
	}

	return mappedGoods, nil
}
