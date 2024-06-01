package server

import (
	"net/http"
)

func CORSMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", config.WebConfig.WebURL)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		// OPTIONSメソッドのプリフライトリクエストに対して200を返す
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
			return
        }

		h.ServeHTTP(w, r)
	})
}