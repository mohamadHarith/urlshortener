package handler

import (
	"log"
	"net/http"
	"strings"
)

func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	log.Println(path)
	segments := strings.Split(path, "/")
	log.Println(segments)
	if len(segments) < 2 {
		writeResp(w, Response{Message: "invalid url key", Status: http.StatusBadRequest})
		return
	}

	key := segments[1]
	log.Println(key)

	url, err := h.urlShortener.Get(key)
	if err != nil {
		writeResp(w, Response{Message: err.Error(), Status: http.StatusInternalServerError})
		return
	}

	log.Println("url=> ", url)

	http.Redirect(w, r, url, http.StatusFound)
}
