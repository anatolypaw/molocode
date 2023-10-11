// app						- Рабочая папка самого приложения
// └── www					- Файлы для сервера веб интерфейса
//	├── build			- Собранный феб интерфейс React
//	├── node_modules	- рабочие файлы React
//	├── public			- рабочие файлы React
//	└── src				- рабочие файлы React
// internal				- Исходники, используемые внутри приложения. Подключаются в main.go
// ├── storage				- Хранилище данных
// ├── ts					- terminal server. Сервер, работающий с терминалами маркировки
// │   └── v1				- api версии 1
// └── ws					- web server - Сервер веб интерфейса
//	└── wapi			- api веб сервера

package main

import (
	"bufio"
	"log"
	"molocode/internal/storage"
	"molocode/internal/structs"
	"molocode/internal/ts"
	"molocode/internal/ws"
	"net/http"
	"os"
	"time"
)

func main() {
	//Инициализируем базу данных
	storage, err := storage.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("storage ready")

	storage.AddUser(structs.User{Login: "admin", Password: "test", Role: "admin"})

	
	//Запускаем сервер веб интерфейса
	go func() {
		s := &http.Server{ 
			Addr:         ":80",
			Handler:      ws.Router(storage),
			IdleTimeout:  1 * time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		log.Printf("Сервер веб интерфейса %s", s.Addr)
		log.Fatal(s.ListenAndServe())
	}()

	//Запускаем сервер работы с терминалами
	go func() {
		s := &http.Server{
			Addr:         ":3000",
			Handler:      ts.Router(),
			IdleTimeout:  1 * time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		
		log.Printf("Сервер API для терминалов %s", s.Addr)
		log.Fatal(s.ListenAndServe())
	}()
	
	
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
