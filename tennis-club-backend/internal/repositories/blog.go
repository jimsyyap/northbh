package repositories

import (
    "time"

    "github.com/jmoiron/sqlx"
    "github.com/jimsyyap/northbh/tennis-club-backend/internal/models"
)

type BlogRepository struct {
    DB *sqlx.DB
}

func NewBlogRepository(db *sqlx.DB) *BlogRepository {
    return &BlogRepository{DB: db}
}

// CreatePost inserts a new blog post into the database
func (r *BlogRepository) CreatePost(post *models.Post) error {
    query := `
        INSERT INTO posts (title, content, author_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`
    post.CreatedAt = time.Now()
    post.UpdatedAt = time.Now()

    return r.DB.QueryRow(query, post.Title, post.Content, post.AuthorID, post.CreatedAt, post.UpdatedAt).Scan(&post.ID)
}

// GetAllPosts retrieves all blog posts
func (r *BlogRepository) GetAllPosts() ([]models.Post, error) {
    query := `SELECT id, title, content, author_id, created_at, updated_at FROM posts ORDER BY created_at DESC`
    var posts []models.Post
    err := r.DB.Select(&posts, query)
    return posts, err
}

// GetPostByID retrieves a single blog post by ID
func (r *BlogRepository) GetPostByID(id int) (*models.Post, error) {
    query := `SELECT id, title, content, author_id, created_at, updated_at FROM posts WHERE id = $1`
    var post models.Post
    err := r.DB.Get(&post, query, id)
    return &post, err
}

// UpdatePost updates an existing blog post
func (r *BlogRepository) UpdatePost(post *models.Post) error {
    query := `
        UPDATE posts
        SET title = $1, content = $2, updated_at = $3
        WHERE id = $4`
    post.UpdatedAt = time.Now()
    _, err := r.DB.Exec(query, post.Title, post.Content, post.UpdatedAt, post.ID)
    return err
}

// DeletePost deletes a blog post by ID
func (r *BlogRepository) DeletePost(id int) error {
    query := `DELETE FROM posts WHERE id = $1`
    _, err := r.DB.Exec(query, id)
    return err
}
