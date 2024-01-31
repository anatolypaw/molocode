package mongostore

import (
	"context"
	"fmt"
	"molocode/internal/app/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Good_dto struct {
	Gtin            string `bson:"_id"`
	Desc            string
	StoreCount      uint
	GetCodeForPrint bool
	AllowProduce    bool
	AllowPrint      bool
	Upload          bool
	CreatedAt       time.Time
}

// Добавляет продукт в хранилище, возвращает все поля добавленного продукта
func (ths *MongoStore) Add(ctx context.Context, good entity.Good) error {
	const op = "mongo.Add"
	// MAPPING
	mappedGood := Good_dto{
		Gtin:            good.Gtin,
		Desc:            good.Desc,
		StoreCount:      good.StoreCount,
		GetCodeForPrint: good.GetCodeForPrint,
		AllowProduce:    good.AllowProduce,
		AllowPrint:      good.AllowPrint,
		Upload:          good.Upload,
		CreatedAt:       good.CreatedAt,
	}
	_, err := ths.db.Collection(collectionGoods).InsertOne(context.TODO(), mappedGood)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (ths *MongoStore) Get(ctx context.Context, gtin string) (entity.Good, error) {
	const op = "mongo.Get"

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
		AllowPrint:      good_dto.AllowPrint,
		Upload:          good_dto.Upload,
		CreatedAt:       good_dto.CreatedAt,
	}

	return mappedGood, nil
}

func (ths *MongoStore) GetAll(ctx context.Context) ([]entity.Good, error) {
	const op = "mongo.GetAll"

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
			AllowPrint:      good_dto.AllowPrint,
			Upload:          good_dto.Upload,
			CreatedAt:       good_dto.CreatedAt,
		})
	}

	return mappedGoods, nil
}
