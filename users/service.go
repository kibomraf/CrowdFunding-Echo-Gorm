package users

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input InputRegister) (Users, error)
	LoginUser(input InputLogin) (Users, error)
}
type service struct {
	repository Repository
}

func UserService(repository Repository) *service {
	return &service{repository}
}
func (s *service) RegisterUser(input InputRegister) (Users, error) {
	user := Users{
		Name:       input.Name,
		Occupation: input.Occupation,
		Email:      input.Email,
		Role:       "user",
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password_hash = string(password)
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
func (s *service) LoginUser(input InputLogin) (Users, error) {
	email := input.Email
	password := input.Password
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("email not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}
