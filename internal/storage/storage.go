package storage

import (
	"fmt"
	"molocode/internal/storage/mongodb"
)

// Интерфейсы
type Storage interface {

	// Выдает экзэмпляр хранилища
	New()

	/* Управление продуктами файл goods.go*/
	//Выводит список продуктов в хранилище
	GetGoods()

	//Добавляет продукт в хранилище
	AddGood()

	/* Управление пользователями файл users.go*/
	// Создать пользователя
	AddUser()

}


// Код
type Code struct {
	Gtin   string // gtin продукта
	Serial string // серийный номер КМ
	Crypto string // криптохвост
}

type S struct {
	mongodb *mongodb.Mongodb
}


/* Возвращает инициализированное хранилище */
func New(mongoPath string, mongodbName string) (*S, error) {
	const op = "storage.New"

	/* Подключение MongoDB */
	mdb, err := mongodb.NewMongodb(mongoPath, mongodbName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = mdb.InitCollection()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &S{mongodb: mdb}, nil
}

