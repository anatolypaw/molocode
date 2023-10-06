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
// Продукт, gtin для каждого уникален. 14 символов
type Good struct {
	Gtin       string // gtin продукта
	Desc       string // описание продукта
	StoreCount int    // сколько хранить кодов
	Get        bool   // флаг, получать коды из 1с
	Upload     bool   // флаг, выгружать коды в 1с
	Awaible    bool   // флаг, выдавать ли кода на терминал
	ShelfLife  int    // срок годности продукта
}

// Код
type Code struct {
	Gtin   string // gtin продукта
	Serial string // серийный номер КМ
	Crypto string // криптохвост
}