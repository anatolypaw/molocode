package mongo

import (
	"fmt"
	"molocode/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
)

// Добавляет продукт в хранилище, возвращает все поля добавленного продукта
func (s *Store) AddGood(g entity.Good) error {
	const op = "hubstore.AddGood"
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
	_, err := s.db.Collection(collectionGoods).InsertOne(s.ctx, mappedGood)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Store) GetGood(gtin string) (entity.Good, error) {
	const op = "mongo.GetGood"

	filter := bson.M{"_id": gtin}
	reqResult := s.db.Collection(collectionGoods).FindOne(s.ctx, filter)

	var good entity.Good
	err := reqResult.Decode(&good)
	if err != nil {
		return entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}
	if good.Gtin == "" {
		return entity.Good{}, fmt.Errorf("%s: Продукт не найден", op)
	}

	// MAPPING
	mappedGood := entity.Good{
		Gtin:            string(good.Gtin),
		Desc:            good.Desc,
		StoreCount:      good.StoreCount,
		GetCodeForPrint: good.GetCodeForPrint,
		AllowProduce:    good.AllowProduce,
		Upload:          good.Upload,
		CreatedAt:       good.CreatedAt,
	}

	return mappedGood, nil
}

func (s *Store) GetAllGoods() ([]entity.Good, error) {
	const op = "mongo.GetAllGoods"

	filter := bson.M{}
	cursor, err := s.db.Collection(collectionGoods).Find(s.ctx, filter)
	if err != nil {
		return []entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	goods := []Good_dto{}
	err = cursor.All(s.ctx, &goods)
	if err != nil {
		return []entity.Good{}, fmt.Errorf("%s: %w", op, err)
	}

	// MAPPING
	mappedGoods := []entity.Good{}
	for _, good := range goods {
		mappedGoods = append(mappedGoods, entity.Good{
			Gtin:            string(good.Gtin),
			Desc:            good.Desc,
			StoreCount:      good.StoreCount,
			GetCodeForPrint: good.GetCodeForPrint,
			AllowProduce:    good.AllowProduce,
			Upload:          good.Upload,
			CreatedAt:       good.CreatedAt,
		})
	}

	return mappedGoods, nil
}
