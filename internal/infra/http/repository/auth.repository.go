package repository

import (
	"context"
	"errors"
	"trivium/internal/domain/entity"
)

type AuthRepository struct{}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (f *AuthRepository) ValidateToken(token string) (bool, error) {
	if token == "" {
		return false, errors.New("token não pode ser vazio")
	}

	_, err := GetAuthClient().VerifyIDToken(context.Background(), token)
	if err != nil {
		return false, errors.New("token inválido: " + err.Error())
	}

	return true, nil
}

func (f *AuthRepository) ConvertTokenInUser(token string) (*entity.User, error) {
	if token == "" {
		return nil, errors.New("token não pode ser vazio")
	}

	decodedToken, err := GetAuthClient().VerifyIDToken(context.Background(), token)
	if err != nil {
		return nil, errors.New("erro ao decodificar token: " + err.Error())
	}

	userRecord, err := GetAuthClient().GetUser(context.Background(), decodedToken.UID)
	if err != nil {
		return nil, errors.New("erro ao obter informações do usuário: " + err.Error())
	}

	user := &entity.User{
		Email:     userRecord.Email,
		Name:      userRecord.DisplayName,
		PhotoPath: userRecord.PhotoURL,
	}

	return user, nil
}
