package codeloader

import (
	"1c/entity"
	"1c/internal/httpClient"
	"fmt"
	"log"
)

func Load(config *entity.Config) {
	// Запрашиваем из hub продукты, для которых нужно запросить коды в 1с и количество кодов
	res, err := httpClient.GetReqCodeCountFromHub(config.Hub.IP)
	if err != nil {
		log.Print(err)
	}

	for _, good := range res {
		fmt.Printf("Для %s %15s: %4v кодов\n", good.Gtin, good.Desc, good.RequiredCount)
	}

	for _, good := range res {
		fmt.Printf("Для %s %15s: %4v кодов\n", good.Gtin, good.Desc, good.RequiredCount)
	}
}
