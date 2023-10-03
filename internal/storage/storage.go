package storage

// Интерфейсы
type Storage interface {

	//Закрывает подключение к хранилищу
	Close() error

	//Выводит список продуктов в хранилище
	GetGoods() ([]Good, error)

	//Добавляет продукт в хранилище
	AddGood(gtin string, description string) error
}

// Структуры данных
type Good struct {
	gtin        string
	description string
}
