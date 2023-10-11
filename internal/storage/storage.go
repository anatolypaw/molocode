package storage

import (
	"fmt"
	"molocode/internal/storage/mongodb"
	"molocode/internal/structs"
)

// Интерфейсы
type IStorage interface {

	// Выдает экзэмпляр хранилища
	New()

	/* Управление продуктами файл goods.go*/
	//Выводит список продуктов в хранилище
	GetGoods()

	//Добавляет продукт в хранилище
	AddGood()

	/* Управление пользователями файл users.go*/
	AddUser(structs.User) error
	DeleteUser(login string) error
	EditUserRole(login string, role string) error
	GetUser(login string) ()

}


// Код
type Code struct {
	Gtin   string // gtin продукта
	Serial string // серийный номер КМ
	Crypto string // криптохвост
}

type Storage struct {
	mongodb *mongodb.Mongodb
}


/* Возвращает инициализированное хранилище */
func New(mongoPath string, mongodbName string) (*Storage, error) {
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

	return &Storage{mongodb: mdb}, nil
}

