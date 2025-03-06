package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/jimsyyap/northbh/tennis-club-backend/internal/models"
    "github.com/jimsyyap/northbh/tennis-club-backend/internal/repositories"
)

type BlogHandler struct {
    repo *repositories.BlogRepository
}

func NewBlogHandler(repo *repositories.BlogRepository) *BlogHandler {
    return &BlogHandler{repo: repo}
}

// CreatePost creates a new blog post
func (h *BlogHandler) CreatePost(c *gin.Context) {
    var post models.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Assume user ID is extracted from JWT token (admin only)
    userID := 1 // Replace with actual user ID from context
    post.AuthorID = userID

    if err := h.repo.CreatePost(&post); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
        return
    }

    c.JSON(http.StatusCreated, post)
}

// GetAllPosts retrieves all blog posts
func (h *BlogHandler) GetAllPosts(c *gin.Context) {
    posts, err := h.repo.GetAllPosts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
        return
    }

    c.JSON(http.StatusOK, posts)
}

// GetPostByID retrieves a single blog post by ID
func (h *BlogHandler) GetPostByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    post, err := h.repo.GetPostByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    c.JSON(http.StatusOK, post)
}

// UpdatePost updates an existing blog post
func (h *BlogHandler) UpdatePost(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    var post models.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    post.ID = id
    if err := h.repo.UpdatePost(&post); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
        return
    }

    c.JSON(http.StatusOK, post)
}

// DeletePost deletes a blog post by ID
func (h *BlogHandler) DeletePost(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    if err := h.repo.DeletePost(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
