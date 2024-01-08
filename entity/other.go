package entity

// Представляет продукт и количество кодов, которое нужно загрузить из 1с
type CodeReq struct {
	Gtin          string
	Desc          string // Имя продукта
	RequiredCount int64
}

// Информация о коде, передаваемая на терминал для нанесения
type CodeForPrint struct {
	Gtin    string
	Serial  string
	Crypto  string
	PrintId uint32
}
