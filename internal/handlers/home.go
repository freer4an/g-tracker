package handlers

import (
	"context"
	"net/http"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	val, ok := r.Context().Value("key").(string)
	if !ok {
		w.Write([]byte("Hello, World!"))
		return
	}
	w.Write([]byte("Hello, World! " + val))
}

func (h *Handler) Context(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "key", "value")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
