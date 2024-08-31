package handlers

import (
	"blog-rest/internal/models"
	"blog-rest/internal/storage/memory"
	"encoding/json"
	"net/http"
)

type PostHandler struct {
	store *memory.BlogStorage
}

func NewPostHandler(store *memory.BlogStorage) *PostHandler {
	return &PostHandler{store: store}
}

func (h *PostHandler) HandlePostCreate(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	h.store.AddPost(post)

	w.WriteHeader(http.StatusCreated)
}
