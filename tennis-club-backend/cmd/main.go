package main

import (
    "github.com/jimsyyap/northbh/tennis-club-backend/internal/handlers"
    "github.com/jimsyyap/northbh/tennis-club-backend/internal/middleware"
    "github.com/jimsyyap/northbh/tennis-club-backend/internal/repositories"
    "github.com/jimsyyap/northbh/tennis-club-backend/pkg"

    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize database
    pkg.InitDB("host=localhost port=5432 user=jim dbname=tennis_club sslmode=disable password=havetheexperience")

    // Initialize repositories
    blogRepo := repositories.NewBlogRepository(pkg.DB)
    userRepo := repositories.NewUserRepository(pkg.DB)

    // Initialize handlers
    blogHandler := handlers.NewBlogHandler(blogRepo)
    authHandler := handlers.NewAuthHandler(userRepo)

    // Set up Gin router
    r := gin.Default()

    // Public routes
    // Auth routes
    r.POST("/api/auth/register", authHandler.Register)
    r.POST("/api/auth/login", authHandler.Login)

    // Blog routes - public read access
    r.GET("/api/posts", blogHandler.GetAllPosts)       // Get all posts
    r.GET("/api/posts/:id", blogHandler.GetPostByID)   // Get a single post

    // Protected routes
    // User routes - require authentication
    userRoutes := r.Group("/api/user")
    userRoutes.Use(middleware.AuthMiddleware())
    {
        userRoutes.GET("/me", authHandler.GetCurrentUser)
    }

    // Protected blog routes - require authentication
    protectedBlogRoutes := r.Group("/api/protected")
    protectedBlogRoutes.Use(middleware.AuthMiddleware())
    {
        protectedBlogRoutes.POST("/posts", blogHandler.CreatePost)       // Create a new post
        protectedBlogRoutes.PUT("/posts/:id", blogHandler.UpdatePost)    // Update a post
        protectedBlogRoutes.DELETE("/posts/:id", blogHandler.DeletePost) // Delete a post
    }

    // Admin-only routes
    adminRoutes := r.Group("/api/admin")
    adminRoutes.Use(middleware.AdminMiddleware())
    {
        // Add admin-specific endpoints here
    }

    // Start server
    r.Run(":8080")
}
