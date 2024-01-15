package codeloader

import (
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

	// Выводим список продуктов и требуемое количество кодов
	for _, good := range res {
		log.Printf("%s: для %s %15s: нужно %4v кодов", op, good.Gtin, good.Desc, good.RequiredCount)
	}

	// Запрашиваем коды в шлюзе и результат помещаем в marks
	log.Println("Запрос кодов")
	for _, good := range res {
		if good.RequiredCount <= 0 {
			log.Printf("%s: %s: %s запрос кодов не требуется", op, good.Gtin, good.Desc)
			continue
		}

		log.Printf("%s: %s: %s", op, good.Gtin, good.Desc)
		marks, err := httpClient.GetNewCodesFromGate(config.MainExchangemarks.IP, good.Gtin, 100)
		if err != nil {
			log.Printf("%s: %s", op, err)
			continue
		}
		log.Printf("%s: для %s %s получено %v КМ", op, good.Gtin, good.Desc, len(marks))
		//log.Println(marks)

		// Передаем в хаб полученные коды
		okCount := 0
		for _, code := range marks {
			err := httpClient.UploadCodeToHub(config.Hub.IP, "gatexch", code)
			if err != nil {
				log.Printf("%s: %s", op, err)
				continue
			}
			okCount++
		}
		log.Printf("%s: %s %s Передано в hub %v кодов", op, good.Gtin, good.Desc, okCount)

	}
}
