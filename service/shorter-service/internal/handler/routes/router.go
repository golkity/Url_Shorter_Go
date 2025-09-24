package routes

import (
	"url-shortener/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(urlHandler *handler.URLHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{shortUrl}", urlHandler.Redirect)

	r.Group(func(r chi.Router) {
		r.Get("/swagger/*", httpSwagger.WrapHandler)

		r.Route("/api/v1", func(r chi.Router) {
			r.Post("/shorten", urlHandler.CreateShortURL)
		})
	})

	return r
}
