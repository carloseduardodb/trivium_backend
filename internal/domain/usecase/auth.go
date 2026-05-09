package usecase

import (
	"trivium/internal/domain/entity"
	"trivium/internal/domain/repositorier"
	"trivium/internal/presentation/dto"
)

type AuthUseCase struct {
	firebaseRepo repositorier.AuthRepositorier
	userRepo     repositorier.UserRepositorier
}

func NewAuthUseCase(firebaseRepo repositorier.AuthRepositorier, userRepo repositorier.UserRepositorier) *AuthUseCase {
	return &AuthUseCase{
		firebaseRepo: firebaseRepo,
		userRepo:     userRepo,
	}
}

func (a *AuthUseCase) Auth(auth *dto.Auth) (*entity.User, error) {
	firebaseUser, err := a.firebaseRepo.ConvertTokenInUser(auth.Token)
	if err != nil {
		return nil, err
	}

	existingUser, err := a.userRepo.FindByEmail(firebaseUser.Email)
	if err == nil && existingUser != nil {
		return existingUser, nil
	}

	newUser, err := entity.NewUser(firebaseUser.Name, firebaseUser.Email, firebaseUser.PhotoPath)
	if err != nil {
		return nil, err
	}

	savedUser, err := a.userRepo.Save(*newUser)
	if err != nil {
		return nil, err
	}

	return &savedUser, nil
}
