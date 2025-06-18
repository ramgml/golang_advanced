package auth

import "purple/4-order-api/pkg/db"

type SessionRepository struct {
	Database *db.Db
}

func NewSessionRepository(database *db.Db) *SessionRepository {
	return &SessionRepository{
		Database: database,
	}
}

func (repo *SessionRepository) Create(session *Session) (*Session, error) {
	result := repo.Database.DB.Create(session)
	if result.Error != nil {
		return nil, result.Error
	}
	return session, nil
}

func (repo *SessionRepository) GetByUid(sessionUid string) (*Session, error) {
	var session Session
	result := repo.Database.DB.First(&session, "uid = ?", sessionUid)	
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}

func (repo *SessionRepository) GetByPhone(phone string) (*Session, error) {
	var session Session
	result := repo.Database.DB.First(&session, "phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}
