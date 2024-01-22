package main

import (
	"log/slog"
	"molocode/internal/app/repository/mongostore"
	"molocode/internal/app/usecase/usecase_admin"
	"molocode/internal/app/usecase/usecase_exchange"
	"molocode/internal/view/http/mymiddleware"
	v1 "molocode/internal/view/http/v1"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/lmittmann/tint"
)

func main() {
	/* Настройка логгера */
	logger := slog.New(tint.NewHandler(os.Stdout, nil))
	//logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Включены DEBUG сообщения")
	logger.Info("Включены INFO сообщения")
	logger.Warn("Включены WARN сообщения")
	logger.Error("Включены ERROR сообщения")

	/* Подключение к базе данных */
	mstore, err := mongostore.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	/* Инициализация usecase */
	admUsecase := usecase_admin.New(mstore)
	exchUsecase := usecase_exchange.New(mstore, mstore)

	/* Инициализация http сервера */
	router := chi.NewRouter()

	// Логгер slog встраивается в context
	// Здесь же на каждый request создается уникальный req_id и встраивается в context
	// он выводится для всего дерева вызовов
	router.Use(mymiddleware.Logger(logger))

	router.Post("/v1/addGood", v1.AddGood(&admUsecase))
	router.Get("/v1/getAllGoods", v1.GetAllGoods(&admUsecase))

	router.Get("/v1/getGoodsReqCodes", v1.GetGoodsReqCodes(&exchUsecase))
	router.Post("/v1/addCodeForPrint", v1.AddCodeForPrint(&exchUsecase))

	s := &http.Server{
		Addr:         "localhost:3000",
		Handler:      router,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	logger.Info("Server run on", "addres", s.Addr)
	logger.Error(s.ListenAndServe().Error())
}
