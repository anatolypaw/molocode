package production_usecase

import "molocode/internal/domain/entity"

// Для этого сценария нужны данные о кодах
type hubService interface {
	GetCodeForPrint(entity.Gtin)
}

type ProductionUsecase struct {
	hub hubService
}

// Возвращает код для нанесения на упаковку
func (u ProductionUsecase) GetCodeForPrint(gtin entity.Gtin) error {

	u.hub.GetCodeForPrint(gtin)
	return nil
}

// Помечает нанесенный код произведенным
func (pu ProductionUsecase) SetCodePrinted() error {
	return nil
}
