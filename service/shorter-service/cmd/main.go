package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
	"url-shortener/internal/custom_errors"
	"url-shortener/internal/handler"
	"url-shortener/internal/handler/routes"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
)

func main() {
	Port := os.Getenv("APP_PORT")
	dbURL := os.Getenv("POSTGRES_DSN")

	dbpool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	urlRepo := repository.NewPostgresRepository(dbpool)

	urlService := service.NewURLService(urlRepo)
	urlHandler := handler.NewURLHandler(urlService)
	router := routes.NewRouter(urlHandler)

	log.Printf("Swwager: http://localhost:%s/swagger/index.html", Port)

	if err := http.ListenAndServe(":"+Port, router); err != nil {
		log.Fatalf("%s : %v", custom_errors.Error_StartServer, err)
	}
}
