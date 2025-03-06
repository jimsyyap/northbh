package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimsyyap/northbh/tennis-club-backend/internal/middleware"
	"github.com/jimsyyap/northbh/tennis-club-backend/internal/models"
	"github.com/jimsyyap/northbh/tennis-club-backend/internal/repositories"
)

type AuthHandler struct {
	userRepo *repositories.UserRepository
}

func NewAuthHandler(userRepo *repositories.UserRepository) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var userRegister models.UserRegistration
	
	// Bind and validate the request body
	if err := c.ShouldBindJSON(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	
	// Check if user already exists
	exists, err := h.userRepo.CheckUserExists(userRegister.Email, userRegister.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check if user exists"})
		return
	}
	
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email or username already exists"})
		return
	}
	
	// Hash the password
	hashedPassword, err := middleware.HashPassword(userRegister.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	
	// Create the new user
	newUser := &models.User{
		Username:     userRegister.Username,
		Email:        userRegister.Email,
		PasswordHash: hashedPassword,
		Role:         "user", // Default role
	}
	
	if err := h.userRepo.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	
	// Generate JWT token
	token, err := middleware.GenerateToken(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	
	// Return user data (excluding password) and token
	c.JSON(http.StatusCreated, models.TokenResponse{
		Token: token,
		User:  *newUser,
	})
}

// Login handles user authentication
func (h *AuthHandler) Login(c *gin.Context) {
	var userLogin models.UserLogin
	
	// Bind and validate the request body
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	
	// Find the user by email
	user, err := h.userRepo.GetUserByEmail(userLogin.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}
	
	// Verify the password
	if !middleware.CheckPasswordHash(userLogin.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	
	// Generate JWT token
	token, err := middleware.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	
	// Return user data and token
	c.JSON(http.StatusOK, models.TokenResponse{
		Token: token,
		User:  *user,
	})
}

// GetCurrentUser returns the current authenticated user
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}
	
	user, err := h.userRepo.GetUserByID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}
	
	c.JSON(http.StatusOK, user)
}
