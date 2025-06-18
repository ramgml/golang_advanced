package auth

import (
	"fmt"
	"math/rand/v2"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Uid   string `json:"sessionUid" gorm:"uniqueIndex"`
	Phone string `json:"phone" gorm:"uniqueIndex"`
	Code  string
}

func NewSession(phone string) *Session {
	session := Session{
		Uid:   uuid.NewString(),
		Phone: phone,
	}
	session.GenerateCode(4)
	return &session
}

func (s *Session)GenerateCode(length int) {
	code := ""
	for range length {
		code = fmt.Sprintf("%s%v", code, rand.IntN(10))
	}
	s.Code = code
}
