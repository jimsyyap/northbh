package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jimsyyap/northbh/tennis-club-backend/internal/models"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (username, email, password_hash, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`
	
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	
	// If role is empty, set default role as "user"
	if user.Role == "" {
		user.Role = "user"
	}

	return r.DB.QueryRow(
		query,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
}

// GetUserByEmail retrieves a user by their email address
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = $1`
	err := r.DB.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID retrieves a user by their ID
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.DB.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CheckUserExists checks if a user with the given email or username already exists
func (r *UserRepository) CheckUserExists(email, username string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE email = $1 OR username = $2`
	err := r.DB.Get(&count, query, email, username)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
