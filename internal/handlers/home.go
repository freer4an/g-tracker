package handlers

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/freer4an/groupie-tracker/internal/models"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	data, err := models.GetHomeData()
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	// handle search param request
	search := r.URL.Query().Get("search")
	if search != "" {
		h.search(search, w, r)
		return
	}

	// handle filter param request
	filter := r.URL.Query().Get("filter")
	if filter != "" {
		h.filter(filter, w, err, http.StatusInternalServerError)
		return
	}

	// handle no params
	if err := temp.ExecuteTemplate(&buf, "index.html", data); err != nil {
		h.errorResponse(w, errors.Join(err, ErrTemplate), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

func (h *Handler) search(search string, w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) filter(filter string, w http.ResponseWriter, err error, code int) {

}
