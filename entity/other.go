package entity

// Представляет продукт и количество кодов, которое нужно загрузить из 1с
type CodeReq struct {
	Gtin          string `bson:"gtin" json:"gtin"`
	Desc          string `bson:"desc" json:"desc"` // Имя продукта
	RequiredCount int64  `bson:"required_count" json:"required_count"`
}

// Информация о коде, передаваемая на терминал для нанесения
type CodeForPrint struct {
	Gtin    string `json:"gtin"`
	Serial  string `json:"serial"`
	Crypto  string `json:"crypto"`
	PrintID uint32 `json:"print_id"`
}
