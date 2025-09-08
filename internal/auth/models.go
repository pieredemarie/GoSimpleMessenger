package auth

import "github.com/pieredemarie/GoSimpleMessenger/internal/storage"

type Handler struct {
	Storage storage.AuthStorage
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
