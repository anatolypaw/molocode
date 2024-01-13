package model

type Code struct {
	Gtin   string `json:"gtin"`
	Serial string `json:"serial"`
	Crypto string `json:"crypto"`
}
