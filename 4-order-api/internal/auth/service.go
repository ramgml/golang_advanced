package auth

import (
	"errors"
	"purple/4-order-api/internal/user"
)

type AuthService struct {
	SessionRepository *SessionRepository
	UserRepository *user.UserRepository
}

func NewAuthService(sessionRepositry *SessionRepository, userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		SessionRepository: sessionRepositry,
		UserRepository: userRepository,
	}
}

func (s *AuthService) Auth(phone string) (*Session, error) {
	session, err := s.SessionRepository.GetByPhone(phone)
	if err != nil {
		newSession := NewSession(phone)
		session, err = s.SessionRepository.Create(newSession)
	}
	return session, err
}

func (s *AuthService) Verify(sessionUid string, code string) (string, error) {
	session, err := s.SessionRepository.GetByUid(sessionUid)
	if err != nil {
		return "", err
	}
	if session.Code != code {
		return "", errors.New("wrong code")
	}
	user, err := s.login_or_register(session.Phone)
	if err != nil {
		return "", err
	}
	return user.Phone, nil
}

func (s *AuthService) login_or_register(phone string) (*user.User, error) {
	existedUser, _ := s.UserRepository.GetByPhone(phone)
	if existedUser != nil {
		return existedUser, nil
	}
	newUser := &user.User{
		Phone: phone,
	}
	_, err := s.UserRepository.Create(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
