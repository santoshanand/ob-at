package router

import (
	"encoding/json"
	"net/http"
)

func spaCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e := "Cached For Speed"
		w.Header().Set("Etag", e)
		w.Header().Set("Cache-Control", "max-age=86400") // 24 hour cache
		next.ServeHTTP(w, r)
	})
}

func (p *routes) writeJSON(w http.ResponseWriter, statusCode int, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")

	w.WriteHeader(statusCode)
	data, err := json.Marshal(value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("error to Marshal"))
		if err != nil {
			return
		}
	}
	_, errWrite := w.Write(data)
	if errWrite != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("error to write"))
		if err != nil {
			return
		}
	}
}
