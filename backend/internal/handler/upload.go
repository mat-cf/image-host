package handler

import (
	"fmt"
	"net/http"

	"github.com/mat-cf/image-host/internal/service"
)

type ImageHandler struct {
	service service.ImageService
}

func NewImageHandler(s service.ImageService) *ImageHandler {
	return &ImageHandler{service: s}
}

func (h *ImageHandler) Upload(w http.ResponseWriter, r *http.Request) {
	// limits to 5mb
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
			http.Error(w, "file too big", http.StatusBadRequest)
			return
	}

	// reads requisition file
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
			http.Error(w, "error reading file", http.StatusBadRequest)
			return
	}
	defer file.Close()

	url, err := h.service.Upload(r.Context(), file, fileHeader)
	if err != nil {
			http.Error(w, "error saving img", http.StatusInternalServerError)
			return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"url": "%s"}`, url)
}