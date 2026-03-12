package handler

import (
	"net/http"
	"user-service/internal/store"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	DB *store.Postgres
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserHandler(db *store.Postgres) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.DB.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to list users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.DB.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if req.Name == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name and email are required",
		})
		return
	}

	user, err := h.DB.CreateUser(req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}