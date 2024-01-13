package services

import "molocode/internal/store"

type HubService struct {
	Store store.Store
}

func NewHubService(s store.Store) *HubService {
	return &HubService{Store: s}
}
