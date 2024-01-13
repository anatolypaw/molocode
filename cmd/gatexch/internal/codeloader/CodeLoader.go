package codeloader

import (
	"fmt"
	"gatexch/internal/config"
	"gatexch/internal/httpClient"
	"log"
)

func Load(config *config.Config) {
	const op = "CodeLoader.Load"
	// Запрашиваем из hub продукты, для которых нужно запросить коды в 1с и количество кодов
	res, err := httpClient.GetReqCodeCountFromHub(config.Hub.IP)
	if err != nil {
		log.Print(err)
		return
	}

	for _, good := range res {
		log.Printf("%s: для %s %15s: нужно %4v кодов", op, good.Gtin, good.Desc, good.RequiredCount)
	}

	log.Println("Запрос кодов")
	for _, good := range res {
		log.Printf("%s: %s: %s", op, good.Gtin, good.Desc)
		marks, err := httpClient.GetNewCodesFromGate(config.MainExchangemarks.IP, good.Gtin, 2)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(marks)
	}
}
