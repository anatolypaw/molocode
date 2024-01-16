package production_usecase

import (
	"molocode/internal/entity"
)

type storeService interface {
}

type usecase struct {
	store storeService
}

// Возвращает код для нанесения на упаковку
func (u usecase) GetCodeForPrint(gtin entity.Gtin) error {
	return nil
}

// Помечает нанесенный код произведенным
func (pu usecase) SetCodePrinted() error {
	return nil
}
