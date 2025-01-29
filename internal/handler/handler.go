package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	service "github.com/mohamadHarith/urlshortener/internal/service/urlshortener"
)

type Handler struct {
	urlShortener *service.URLShortener
}

func New(urlShortener *service.URLShortener) *Handler {
	return &Handler{
		urlShortener: urlShortener,
	}
}

type Response struct {
	Message string
	Status  int32
}

func writeResp(w http.ResponseWriter, resp any) {
	j, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(j))
}
