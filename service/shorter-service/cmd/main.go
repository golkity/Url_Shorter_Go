package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"url-shortener/internal/custom_errors"
	http2 "url-shortener/internal/handler/http"
	"url-shortener/internal/handler/routes"
	"url-shortener/internal/kafka"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	port := os.Getenv("APP_PORT")
	dbURL := os.Getenv("POSTGRES_DSN")
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	kafkaBrokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	urlRepo := repository.NewPostgresRepository(dbpool)
	var userRepo repository.UserRepository = urlRepo

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go kafka.RunConsumer(ctx, log, kafkaBrokers, userRepo)

	urlService := service.NewURLService(urlRepo)
	urlHandler := http2.NewURLHandler(urlService)

	router := routes.NewRouter(urlHandler, jwtSecret)

	log.Info("Swagger UI available", "url", "http://localhost:"+port+"/swagger/index.html")
	log.Info("Starting server", "port", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Error(custom_errors.Error_StartServer.Error(), "error", err)
		os.Exit(1)
	}
}
