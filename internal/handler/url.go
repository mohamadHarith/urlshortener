package handler

import (
	"encoding/json"
	"io"
	"net/http"
)

type ShortenURLRequest struct {
	URL string `json:"url"`
	Key string `json:"key"`
}

func (h *Handler) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeResp(w, Response{Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	var req ShortenURLRequest
	if err := json.Unmarshal(body, &req); err != nil {
		writeResp(w, Response{Message: err.Error(), Status: http.StatusBadRequest})
		return
	}

	h.urlShortener.Store(req.Key, req.URL)

	w.Header().Set("content-type", "application/json")
	writeResp(w, Response{Message: "successful", Status: http.StatusOK})
}
