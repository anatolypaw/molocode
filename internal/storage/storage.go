package storage

// Интерфейсы
type Storage interface {

	//Закрывает подключение к хранилищу
	Close() error

	//Выводит список продуктов в хранилище
	GetGoods() ([]good, error)

	//Добавляет продукт в хранилище
	AddGood(gtin string, description string) error
}

// Структуры данных
type good struct {
	good_id     int
	gtin        string
	description string
}
