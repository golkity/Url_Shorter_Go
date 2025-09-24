package models

type CreateShortURLRequest struct {
	URL string `json:"url" validate:"required,url"`
}
type CreateShortURLResponse struct {
	ShortURL string `json:"short_url"`
}
