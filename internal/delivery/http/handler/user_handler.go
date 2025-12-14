package handler

import (
	"errors"
	"net/http"
	"strconv"

	"gin_frmr/internal/domain"
	"gin_frmr/internal/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: uc,
	}
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"omitempty,email"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userUseCase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to fetch users",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    users,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	user, err := h.userUseCase.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, Response{
				Success: false,
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to fetch user",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    user,
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user, err := h.userUseCase.CreateUser(req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user, err := h.userUseCase.UpdateUser(uint(id), req.Name, req.Email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, Response{
				Success: false,
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to update user",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "User updated successfully",
		Data:    user,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid user ID",
		})
		return
	}

	if err := h.userUseCase.DeleteUser(uint(id)); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, Response{
				Success: false,
				Message: "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to delete user",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "User deleted successfully",
	})
}
