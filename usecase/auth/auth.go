package auth

import "ecommerce/repository/auth"

type AuthUseCaseInterface interface {
	Login(email string, password string) (string, error)
}

type AuthUseCase struct {
	authRepository auth.AuthRepositoryInterface
}

func NewAuthUseCase(authRepo auth.AuthRepositoryInterface) AuthUseCaseInterface {
	return &AuthUseCase{
		authRepository: authRepo,
	}
}

func (auc *AuthUseCase) Login(email string, password string) (string, error) {
	token, err := auc.authRepository.Login(email, password)
	return token, err
}
