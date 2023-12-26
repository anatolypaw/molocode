package model

// Представляет продукт и количество кодов, которое нужно загрузить из 1с
type CodeReq struct {
	Gtin  string
	Count int
}
