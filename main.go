package main

import (
	"log"
	"net/http"

	"github.com/mohamadHarith/urlshortener/internal/handler"
	service "github.com/mohamadHarith/urlshortener/internal/service/urlshortener"
)

func main() {
	h := handler.New(service.NewURLShortener())

	http.HandleFunc("/shorten", h.ShortenURLHandler)
	http.HandleFunc("/r", h.ShortenURLHandler)

	log.Println("Server started at port 80")

	http.ListenAndServe(":80", nil)
}
