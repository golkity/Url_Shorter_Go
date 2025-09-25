package routes

import (
	"url-shortener/internal/handler/http"
	"url-shortener/internal/middleware"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(urlHandler *http.URLHandler, jwtSecret []byte) *chi.Mux {
	r := chi.NewRouter()

	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	r.Get("/{shortUrl}", urlHandler.Redirect)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Auth(jwtSecret))
		r.Post("/shorten", urlHandler.CreateShortURL)
	})

	return r
}
