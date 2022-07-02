package usecase

import (
	"app/app/model"
	"github.com/google/uuid"
	"time"
)

func (u *UseCase) SignIn(credentials model.Credentials) (*model.Session, error) {

	token := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	return &model.Session{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}
