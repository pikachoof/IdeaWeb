package services

import (
	"IdeaWeb/models"
	"IdeaWeb/repositories"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SessionService struct {
	userRepo    *repositories.UserRepository
	sessionRepo *repositories.SessionRepository
}

func NewSessionService(userRepo *repositories.UserRepository, sessionRepo *repositories.SessionRepository) *SessionService {
	return &SessionService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

func (s *SessionService) Register(registerReq *models.RegisterRequest) error {
	_, err := s.userRepo.FindByEmail(registerReq.Email)
	if err == nil {
		return errors.New("user with this email already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:         registerReq.Name,
		Surname:      registerReq.Surname,
		Email:        registerReq.Email,
		PasswordHash: passwordHash,
	}

	return s.userRepo.Create(user)
}

func (s *SessionService) Login(loginReq *models.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(loginReq.Email)
	// err == nil, found
	// err != nil, not found
	if err != nil {
		return "", errors.New("login failed: " + err.Error())
	}

	log.Print(string(user.PasswordHash))
	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(loginReq.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	session := &models.Session{
		ID:        uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := s.sessionRepo.Create(session); err != nil {
		return "", err
	}

	return session.ID, nil
}

func (s *SessionService) ValidateSession(sessionID uint) (*models.User, error) {
	session, err := s.sessionRepo.FindByID(sessionID)
	if err != nil {
		return nil, errors.New("invalid session")
	}

	user, err := s.userRepo.FindByID(session.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
