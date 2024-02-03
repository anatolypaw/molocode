package repository

import (
	"context"
	"hub/internal/entity"
)

type IGoodRepository interface {
	Add(context.Context, entity.Good) error
	Get(context.Context, string) (entity.Good, error)
	GetAll(context.Context) ([]entity.Good, error)
}

type ICodeRepository interface {
	AddCode(context.Context, entity.FullCode) error
	GetCode(
		ctx context.Context,
		gtin string,
		serial string,
	) (entity.Code, error)

	GetCountPrintAvaible(context.Context, string) (uint, error)
	GetCodeForPrint(
		ctx context.Context,
		gtin string,
		terminal string,
	) (entity.CodeForPrint, error)
}
