package handlers

import (
	"blog-rest/internal/helperts"
	"blog-rest/internal/models"
	"blog-rest/internal/storage/memory"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	store *memory.BlogStorage
}

func NewUserHandler(store *memory.BlogStorage) *UserHandler {
	return &UserHandler{store: store}
}

// список
func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	users := h.store.GetUsers()

	fromCreatedAt := r.URL.Query().Get("fromCreatedAt")
	toCreatedAt := r.URL.Query().Get("toCreatedAt")
	nameFilter := r.URL.Query()["name"]
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	// topPostsAmount := r.URL.Query().Get("topPostsAmount")

	var filterUsers []models.User
	for _, user := range users {

		if fromCreatedAt != "" {
			fromTime := helperts.ParseTime(fromCreatedAt)
			if user.CreatedAt.Before(fromTime) {
				continue
			}
		}

		if toCreatedAt != "" {
			toTime := helperts.ParseTime(fromCreatedAt)
			if user.CreatedAt.After(toTime) {
				continue
			}
		}

		if len(nameFilter) > 0 {
			nameMatched := false
			for _, name := range nameFilter {
				if strings.Contains(user.Name, name) {
					nameMatched = true
					break
				}
			}
			if !nameMatched {
				continue
			}
		}
		filterUsers = append(filterUsers, user)
	}

	// if topPostsAmount == "asc" || topPostsAmount == "desc" {

	// }

	start := offset
	end := offset + limit

	if start > len(filterUsers) {
		start = len(filterUsers)
	}

	if end > len(filterUsers) {
		end = len(filterUsers)
	}

	filterUsers = filterUsers[start:end]

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(filterUsers)
}

// create
func (h *UserHandler) HandleUserCreate(w http.ResponseWriter, r *http.Request) {

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.store.AddUser(user)

	w.WriteHeader(http.StatusCreated)
}

// update
func (h *UserHandler) HandleUserUpdate(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/user/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updateRequest struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if err := h.store.UpdateUser(id, updateRequest.Name); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// delete
func (h *UserHandler) HandleUserDelete(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Path[len("/user/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.store.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
