package main

import (
	"log"
	"time"

	"gatexch/internal/codeloader"
	"gatexch/internal/config"
)

func main() {
	log.Println("Запуск сервиса обема с 1c ")

	// Загрузка конфигурационных данных из файла YAML.
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v\n", err)
	}

	for {
		log.Print("Запрос количества требуемых кодов")
		codeloader.Load(cfg)
		time.Sleep(10 * time.Second)
	}

}
