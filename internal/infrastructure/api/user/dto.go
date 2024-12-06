package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"test/internal/domain/entity"
)

type UserDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateUserDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (dto *CreateUserDTO) Validate() error {
	if dto.ID == "" || dto.Name == "" {
		return errors.New("id and name are required")
	}
	return nil
}

func toDTO(user *entity.User) *UserDTO {
	return &UserDTO{
		ID:   user.ID,
		Name: user.Name,
	}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
