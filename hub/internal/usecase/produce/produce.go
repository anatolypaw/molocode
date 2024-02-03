package produce

import (
	"context"
	"errors"
	"fmt"
	"hub/internal/entity"
)

type iGoodRepo interface {
	Add(context.Context, entity.Good) error
	Get(context.Context, string) (entity.Good, error)
}

type iCodeRepo interface {
	AddCode(context.Context, entity.FullCode) error
	GetCode(
		ctx context.Context,
		gtin string,
		serial string,
	) (entity.FullCode, error)

	GetCodeForPrint(
		ctx context.Context,
		gtin string,
		terminal string,
	) (entity.CodeForPrint, error)
}

type ProduceUsecase struct {
	goodRepo iGoodRepo
	codeRepo iCodeRepo
}

func New(goodRepo iGoodRepo, codeRepo iCodeRepo) ProduceUsecase {
	return ProduceUsecase{
		goodRepo: goodRepo,
		codeRepo: codeRepo,
	}
}

// Возвращает код для печати
func (usecase *ProduceUsecase) GetCodeForPrint(
	ctx context.Context,
	gtin string,
	terminal string,
) (entity.CodeForPrint, error) {

	// - Проверить корректность gtin
	err := entity.ValidateGtin(gtin)
	if err != nil {
		return entity.CodeForPrint{}, err
	}

	// - Проверить, разрешено ли для этого продукта выдача кодов для нанесения
	good, err := usecase.goodRepo.Get(ctx, gtin)
	if err != nil {
		return entity.CodeForPrint{},
			fmt.Errorf("ошибка запроса продукта: %s", err)
	}

	if !good.AllowPrint {
		return entity.CodeForPrint{},
			errors.New("для этого продукта запрещено выдача кодов для нанесения")
	}

	// - Получить код для печати
	// - TODO Проверить корректность кода в ответе БД
	codeForPrint, err := usecase.codeRepo.GetCodeForPrint(ctx, gtin, terminal)
	if err != nil {
		return entity.CodeForPrint{}, err
	}

	return codeForPrint, nil
}

// Отмечает ранее напечатанный код произведенным
func (usecase *ProduceUsecase) ProducePrinted(
	ctx context.Context,
	gtin string,
	serial string,
	terminal string,
	prodDate string,
) error {

	// - Проверить корректность gtin
	err := entity.ValidateGtin(gtin)
	if err != nil {
		return err
	}

	// - Проверить корректность serial
	err = entity.ValidateSerial(serial)
	if err != nil {
		return err
	}

	// - Проверить корректность даты

	// - Проверить корректность имени терминала

	// - Проверить, разрешено ли производство для этого продукта
	good, err := usecase.goodRepo.Get(ctx, gtin)
	if err != nil {
		return fmt.Errorf("ошибка запроса продукта: %s", err)
	}

	if !good.AllowProduce {
		return errors.New("для этого продукта запрещено производство")
	}

	// - Проверки статуса кода
	code, err := usecase.codeRepo.GetCode(ctx, gtin, serial)

	// Проверить, был ли этот код отправлен на печать
	if code.PrintInfo.Avaible {
		return errors.New("этот код не был передан на печать")
	}

	// - Проверить, со
	panic(nil)

	return nil
}
