package repositories

import (
	"IdeaWeb/models"
	"database/sql"
	"fmt"
)

type SessionRepositoryInterface interface {
	Create(newSession *models.Session) error
	FindByID(id uint) (*models.Session, error)
	Update(id uint, updatedSession *models.Session) error
	DeleteByID(id uint) error
	Exists(id uint) (bool, error)
}

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{
		db: db,
	}
}

func (r *SessionRepository) Create(newSession *models.Session) error {
	query := `
		INSERT INTO Sessions (id, user_id, expires_at)
		VALUES($1, $2, $3)
	`

	_, err := r.db.Exec(query,
		newSession.ID,
		newSession.UserID,
		newSession.ExpiresAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	return nil
}

func (r *SessionRepository) FindByID(id uint) (*models.Session, error) {
	query := `
		SELECT id, user_id, expires_at
		FROM Sessions
		WHERE id = $1
	`

	session := &models.Session{}
	err := r.db.QueryRow(query, id).Scan(
		&session.ID,
		&session.UserID,
		&session.ExpiresAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("session not found")
	}

	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	return session, nil
}

func (r *SessionRepository) Update(id uint, updatedSession *models.Session) error {
	exists, err := r.Exists(id)
	if err != nil {
		return fmt.Errorf("failed to find session: %w", err)
	}

	if !exists {
		return fmt.Errorf("cannot update non-existing session")
	}

	query := `
		UPDATE users
		SET user_id = $2, expires_at = $3
		WHERE id = $1
	`

	_, err = r.db.Exec(query, id, updatedSession.UserID, updatedSession.ExpiresAt)
	if err != nil {
		return fmt.Errorf("failed to update session: %w", err)
	}

	return nil
}

func (r *SessionRepository) DeleteByID(id uint) error {
	exists, err := r.Exists(id)
	if err != nil {
		return fmt.Errorf("failed to find session: %w", err)
	}

	if !exists {
		return fmt.Errorf("cannot delete non-existing session")
	}

	query := `DELETE FROM Sessions WHERE id = $1`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}

	return nil
}

func (r *SessionRepository) Exists(id uint) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM Sessions WHERE id = $1)`

	var exists bool
	err := r.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if session exists: %w", err)
	}

	return exists, nil
}
