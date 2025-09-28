package main

import (
	"auth-service/internal/crypto"
	handlr "auth-service/internal/handler/http"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/internal/tokens"
	"context"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"

	"auth-service/internal/kafka"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	log.Info("starting auth-service")

	dbURL := getEnv("DATABASE_URL", "")
	jwtSecret := getEnv("JWT_SECRET", "")
	httpAddress := getEnv("HTTP_ADDRESS", ":8080")
	bcryptCost, _ := strconv.Atoi(getEnv("BCRYPT_COST", "12"))
	accessTTL, _ := time.ParseDuration(getEnv("ACCESS_TOKEN_TTL", "15m"))
	refreshTTL, _ := time.ParseDuration(getEnv("REFRESH_TOKEN_TTL", "30d"))
	kafkaBrokers := strings.Split(getEnv("KAFKA_BROKERS", "localhost:9092"), ",")

	kafkaProducer := kafka.NewProducer(kafkaBrokers, "users.created", log)
	defer kafkaProducer.Close()

	if dbURL == "" || jwtSecret == "" {
		log.Error("DATABASE_URL and JWT_SECRET must be set")
		os.Exit(1)
	}

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	log.Info("database connection established")

	authRepo := repository.NewAuthRepository(dbpool)
	hasher := crypto.NewHasher(bcryptCost)
	tokenManager := tokens.NewManager(
		[]byte(jwtSecret),
		int64(accessTTL.Seconds()),
		int64(refreshTTL.Seconds()),
	)
	authService := service.NewAuthService(log, authRepo, hasher, tokenManager, kafkaProducer)
	authHandler := handlr.NewAuthHandler(log, authService)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)

	log.Info("server starting", slog.String("address", httpAddress))
	if err := http.ListenAndServe(httpAddress, r); err != nil {
		log.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
