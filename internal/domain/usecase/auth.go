package usecase

import (
	"fmt"
	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
	"trivium/internal/presentation/dto"
)

type AuthUseCase struct {
	firebaseRepo repositorier.AuthRepositorier
}

func NewAuthUseCase(firebaseRepo repositorier.AuthRepositorier) *AuthUseCase {
	return &AuthUseCase{
		firebaseRepo: firebaseRepo,
	}
}

func (a *AuthUseCase) Auth(auth *dto.Auth) (*entity.User, error) {
	_, err := a.firebaseRepo.ConvertTokenInUser(auth.Token)
	if err != nil {
		return nil, err
	}

	cUser, err := entity.NewUser("Carlos", "Yj6tZ@example.com", "photoPath")
	if err != nil {
		return nil, err
	}

	fmt.Print(cUser)

	return cUser, nil
}
