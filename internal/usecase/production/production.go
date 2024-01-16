package production_usecase

type storeService interface {
}

type usecase struct {
	store storeService
}

// Возвращает код для нанесения на упаковку
func (u usecase) GetCodeForPrint(gtin string) error {
	return nil
}

// Помечает нанесенный код произведенным
func (pu usecase) SetCodePrinted() error {
	return nil
}
