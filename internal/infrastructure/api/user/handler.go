package user

import (
	"encoding/json"
	"net/http"
	"test/internal/application/service/user"
	"test/internal/domain/entity"
)

type Handler struct {
	userService *user.UserService
}

func NewHandler(userService *user.UserService) *Handler {
	return &Handler{userService: userService}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, "missing user id", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		writeError(w, "user not found", http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, toDTO(user))
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	cmd := user.CreateUserCommand{
		ID:   input.ID,
		Name: input.Name,
	}

	if err := cmd.Validate(); err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.userService.SaveUser(&entity.User{ID: input.ID, Name: input.Name}); err != nil {
		writeError(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, toDTO(&entity.User{ID: input.ID, Name: input.Name}))
}
