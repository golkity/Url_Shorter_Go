package main

import (
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
	Postgres := os.Getenv("POSTGRES_DSN")

	repo, err := repository.NewPostgresRepository(Postgres)
	if err != nil {
		log.Fatalf("%s : %v", custom_errors.Error_InitRepository, err)
	}
	defer repo.Close()

	urlService := service.NewURLService(repo)
	urlHandler := handler.NewURLHandler(urlService)
	router := routes.NewRouter(urlHandler)

	log.Printf("Swwager: http://localhost:%s/swagger/index.html", Port)

	if err := http.ListenAndServe(":"+Port, router); err != nil {
		log.Fatalf("%s : %v", custom_errors.Error_StartServer, err)
	}
}
