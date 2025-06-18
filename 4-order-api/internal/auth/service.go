package auth

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
	return &Session{}, nil
}
