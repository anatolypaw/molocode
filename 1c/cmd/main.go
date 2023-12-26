package main

import (
	"1c/internal/config"
	"fmt"
	"log"
)

func main() {
	log.Println("Запуск сервиса обема с 1c ")

	// Загрузка конфигурационных данных из файла YAML.
	configFile := "config.yaml"
	config, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v\n", err)
	}

	fmt.Printf("%#v", config)

	// Запрашиваем продукты, для которых нужно запросить коды в 1с и количество кодов
}
