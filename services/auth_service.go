package services

import (
	"learning-freemarket-api/models"
	"learning-freemarket-api/repositories"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SignUp(email string, password string) error
	Login(email string, password string) (*string, error)
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(repository repositories.IAuthRepository) IAuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) SignUp(email string, password string) error {
	// 平文をデータベースに登録するとセキュリティリスクになるので、ハッシュ化を行う
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}

func (s *AuthService) Login(email string, password string) (*string, error) {
	foundUser, err := s.repository.FindUser(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &foundUser.Email, nil
}
