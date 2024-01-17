package usecase_production

import service "molocode/internal/domain/service/store"

type UseCase struct {
	store *service.Store
}

// Возвращает код для нанесения на упаковку
func (usecase UseCase) GetCodeForPrint(gtin string) error {
	return nil
}

// Помечает нанесенный код произведенным
func (usecase UseCase) SetCodePrinted() error {
	return nil
}
