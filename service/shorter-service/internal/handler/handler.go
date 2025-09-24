package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"url-shortener/internal/custom_errors"
	"url-shortener/internal/handler/models"
	"url-shortener/internal/service"

	"github.com/go-chi/chi/v5"
)

type URLHandler struct {
	service service.URLService
}

func NewURLHandler(s service.URLService) *URLHandler {
	return &URLHandler{service: s}
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req models.CreateShortURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL, err := h.service.CreateShortURL(r.Context(), req.URL)
	if err != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}

	fullShortURL := "http://" + r.Host + "/" + shortURL

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.CreateShortURLResponse{ShortURL: fullShortURL})
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortUrl")
	if shortURL == "" {
		http.Error(w, "Short URL code is missing", http.StatusBadRequest)
		return
	}

	fullURL, err := h.service.GetFullURL(r.Context(), shortURL)
	if err != nil {
		if errors.Is(err, custom_errors.Error_URLNotFound) {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fullURL, http.StatusMovedPermanently)
}
