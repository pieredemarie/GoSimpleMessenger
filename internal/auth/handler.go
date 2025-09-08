package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pieredemarie/GoSimpleMessenger/internal/storage"
)

func (h *Handler) RegisterHandler(c *gin.Context) {
	var req storage.User
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "bad request"})
		return
	}

	err := h.Storage.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"responce": "user succesfully created"})
}

func (h *Handler) LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "bad request"})
		return
	}

	token, err := h.Storage.Login(req.Email,req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Token": token,
	})
}

func (h *Handler) GetUserInfo(c *gin.Context) {
	
}