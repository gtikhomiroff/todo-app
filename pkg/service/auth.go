package service

import (
	"crypto/sha1"
	"fmt"
	app "github.com/gtikhomiroff/todo-app"
	"github.com/gtikhomiroff/todo-app/pkg/repository"
)

const salt = "qdfgelk3gkk34kgroeg00o0"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user app.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
