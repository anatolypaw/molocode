package mongodb

import (
	"context"
	"fmt"
	"storage/model"
)

// Возвращает код к указанному gtin продукту для последующей печати
func (con *Storage) GetCodeForPrint(gtin, terminal string) (model.CodeForPrint, error) {
	const op = "storage.mongodb.GetCodeCodeForPrint"

	// Проверяем, существует ли такой продукт в БД
	goods, err := con.GetGoods(gtin)
	if err != nil {
		return model.CodeForPrint{}, fmt.Errorf("%s: %w", op, err)
	}

	if len(goods) != 1 {
		return model.CodeForPrint{}, fmt.Errorf("%s: %s", op, "Продукт не найден")
	}
	good := goods[0]

	// Проверяем, разрешена ли для этого продукта выдача кодов для печати
	if !good.SendForPrint {
		return model.CodeForPrint{}, fmt.Errorf("%s: %s", op, "Для этого продукта запрещена выдача кодов для нанесения")
	}

	//////////////
	var newCode model.Code

	_, err = con.db.Collection(gtin).InsertOne(context.TODO(), newCode)
	if err != nil {
		return model.CodeForPrint{}, err
	}

	return model.CodeForPrint{}, err
}
