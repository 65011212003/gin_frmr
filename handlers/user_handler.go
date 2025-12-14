package handlers

import (
	"net/http"

	"gin_frmr/database"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"omitempty,email"`
}

// GetUsers returns all users
func GetUsers(c *gin.Context) {
	var users []database.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUser returns a single user by ID
func GetUser(c *gin.Context) {
	var user database.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := database.User{
		Name:  input.Name,
		Email: input.Email,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
	var user database.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&user).Updates(database.User{
		Name:  input.Name,
		Email: input.Email,
	})

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	var user database.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
