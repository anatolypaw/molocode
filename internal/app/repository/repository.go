package repository

import (
	"context"
	"molocode/internal/app/entity"
)

type IGoodRepository interface {
	AddGood(context.Context, entity.Good) error
	GetGood(context.Context, string) (entity.Good, error)
	GetAllGoods(context.Context) ([]entity.Good, error)
}

type ICodeRepository interface {
	GetCountPrintAvaible(context.Context, string) (uint, error)
}
