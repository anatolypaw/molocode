package service

import (
	"molocode/internal/domain/entity"
	"time"
)

type HubStorage interface {
	AddGood(good entity.Good) error
	GetAllGoods() ([]entity.Good, error)
}

type hubService struct {
	storage HubStorage
}

func NewHubService(storage HubStorage) *hubService {
	return &hubService{storage: storage}
}

func (hs *hubService) AddGood(good entity.Good) error {
	err := good.Gtin.Validate()
	if err != nil {
		return err
	}
	good.CreatedAt = time.Now()
	return hs.storage.AddGood(good)
}

func (hs *hubService) GetAllGoods() ([]entity.Good, error) {
	return hs.storage.GetAllGoods()
}
