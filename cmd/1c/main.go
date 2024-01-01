package main

import (
	"fmt"
	"log"
	"molocode/cmd/1c/internal/config"
	"molocode/cmd/1c/internal/httpClient"
)

func main() {
	log.Println("Запуск сервиса обема с 1c ")

	// Загрузка конфигурационных данных из файла YAML.
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v\n", err)
	}

	fmt.Printf("%#v", config)

	// Запрашиваем в storage продукты, для которых нужно запросить коды в 1с и количество кодов

	for i := 0; i < 1000; i++ {
		res, err := httpClient.GetReqCodeCountFromStorage(config.Storage.Host)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(res)
	}
}
