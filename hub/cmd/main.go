package main

import (
	"hub/internal/repository/mongostore"
	"hub/internal/usecase/admin"
	"hub/internal/usecase/exchange"
	"hub/internal/usecase/produce"
	"hub/internal/view/http/mymiddleware"
	v1 "hub/internal/view/http/v1"
	"log/slog"
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
	admUsecase := admin.New(mstore)
	exchUsecase := exchange.New(mstore, mstore)
	prodUsecase := produce.New(mstore, mstore)

	/* Инициализация http сервера */
	router := chi.NewRouter()

	// Логгер slog встраивается в context
	// на каждый request создается уникальный req_id и встраивается в context
	// он выводится в лог для всего дерева вызовов
	router.Use(mymiddleware.Logger(logger))

	// Admin
	router.Post("/v1/admin/addGood", v1.AddGood(admUsecase))
	router.Get("/v1/admin/getAllGoods", v1.GetAllGoods(admUsecase))

	// Exchange
	router.Get("/v1/exchange/getGoodsReqCodes", v1.GetGoodsReqCodes(exchUsecase))
	router.Post("/v1/exchange/addCodeForPrint", v1.AddCodeForPrint(exchUsecase))

	// Produce
	router.Get("/v1/produce/getCodeForPrint", v1.GetCodeForPrint(prodUsecase))
	router.Get("/v1/produce/producePrinted", v1.ProducePrinted(prodUsecase))

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
