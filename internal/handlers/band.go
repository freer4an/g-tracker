package handlers

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/freer4an/groupie-tracker/internal/models"
)

func (h *Handler) BandHandler(w http.ResponseWriter, req *http.Request) {
	buf := bytes.Buffer{}
	id, ok := req.Context().Value("id").(int)
	if !ok {
		h.errorResponse(w, errors.New("Invalid id"), http.StatusBadRequest)
		return
	}
	data, err := models.GetBandData(id)
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}
	if err := temp.ExecuteTemplate(&buf, "artist.html", data); err != nil {
		h.errorResponse(w, errors.Join(err, ErrTemplate), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}
