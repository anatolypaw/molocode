package repository

import (
	"context"
	"molocode/internal/app/entity"
)

type IGoodRepository interface {
	Add(context.Context, entity.Good) error
	Get(context.Context, string) (entity.Good, error)
	GetAll(context.Context) ([]entity.Good, error)
}

type ICodeRepository interface {
	GetCountPrintAvaible(context.Context, string) (uint, error)
	AddCode(context.Context, entity.FullCode) error
	GetCodeForPrint(ctx context.Context, gtin string, terminal string) (entity.CodeForPrint, error)
}
