package main

import (
	"molocode/internal/app/repository/mongostore"
	"molocode/internal/app/usecase/usecase_admin"
	"molocode/internal/view/http/mymiddleware"
	v1 "molocode/internal/view/http/v1"
	"net/http"
	"os"
	"time"

	"log/slog"

	"github.com/go-chi/chi"
	"github.com/lmittmann/tint"
)

func main() {
	// Настройка логгера
	logger := slog.New(tint.NewHandler(os.Stdout, nil))
	logger.Debug("Включены DEBUG сообщения")
	logger.Info("Включены INFO сообщения")
	logger.Warn("Включены WARN сообщения")
	logger.Error("Включены ERROR сообщения")

	mstore, err := mongostore.New("mongodb://localhost:27017/", "molocode")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	admUseCase := usecase_admin.New(mstore)
	router := chi.NewRouter()

	// Логгер встраивается в контекст.
	// Встраивает в контекст reques_id
	router.Use(mymiddleware.Logger(logger))

	router.Post("/v1/addGood", v1.AddGood(&admUseCase))
	router.Get("/v1/getAllGoods", v1.GetAllGoods(&admUseCase))

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
