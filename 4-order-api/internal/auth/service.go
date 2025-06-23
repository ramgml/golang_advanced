package auth

import (
	"errors"
)

type AuthService struct {
	SessionRepository SessionRepository
}

func NewAuthService(sessionRepositry *SessionRepository) *AuthService {
	return &AuthService{
		SessionRepository: *sessionRepositry,
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

func (s *AuthService) Verify(sessionUid string, code string) (*Session, error) {
	session, err := s.SessionRepository.GetByUid(sessionUid)
	if err != nil {
		return nil, err
	}
	if session.Code != code {
		return nil, errors.New("wrong code")
	}
	return session, nil
}
