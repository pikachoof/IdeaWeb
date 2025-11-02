package repositories

import (
	"IdeaWeb/models"
	"fmt"
	"time"

	"database/sql"
)

type UserRepositoryInterface interface {
	Create(newUser *models.User) error // Will be reserved for creating users by admin, the admin sets the password,
	GetAll() ([]*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	SetRole(id uint, newRole *models.UserRole)
	Update(id uint, updatedUser *models.User) error
	DeleteByID(id uint) error
	Exists(id uint) (bool, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(newUser *models.User) error {
	query := `
	INSERT INTO Users (id, name, surname, email, password_hash, role)
	VALUES($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(query,
		newUser.ID,
		newUser.Name,
		newUser.Surname,
		newUser.Email,
		newUser.PasswordHash,
		newUser.Role,
	)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetAll() ([]*models.User, error) {
	query := `
		SELECT * FROM Users
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Email,
			&user.PasswordHash,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("fows iteration error: %w", err)
	}

	return users, nil
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	query := `
		SELECT id, name, surname, email, role FROM Users
		WHERE id = $1
	`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Role,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}

	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, name, surname, email, role FROM Users
		WHERE email = $1
	`

	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Email,
		&user.Role,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}

	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	return user, nil
}

func (r *UserRepository) SetRole(id uint, newRole models.UserRole) error {
	exists, err := r.Exists(id)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	if !exists {
		return fmt.Errorf("cannot set role for non-existing user")
	}

	query := `
		UPDATE users
		SET role = $1, updated_at = $2
		WHERE id = $3
	`

	_, err = r.db.Exec(query, newRole, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to update user role: %w", err)
	}

	return nil
}

func (r *UserRepository) Update(id uint, updatedUser *models.User) error {
	exists, err := r.Exists(id)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	if !exists {
		return fmt.Errorf("cannot update non-existing user")
	}

	query := `
		UPDATE users
		SET name = $1, surname = $2, email = $3, password_hash = $4, role = $5, updated_at = $6
		WHERE id = $7
	`

	_, err = r.db.Exec(query, updatedUser.Name, updatedUser.Surname, updatedUser.Email, updatedUser.PasswordHash, updatedUser.Role, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (r *UserRepository) DeleteByID(id uint) error {
	exists, err := r.Exists(id)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	if !exists {
		return fmt.Errorf("cannot delete non-existing user")
	}

	query := `DELETE FROM Users WHERE id = $1`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (r *UserRepository) Exists(id uint) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM Users WHERE id = $1)`

	var exists bool
	err := r.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if user exists: %w", err)
	}

	return exists, nil
}
