package users

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input InputRegister) (Users, error)
	LoginUser(input InputLogin) (Users, error)
	IsEmailAvailable(input ChekEmail) (bool, error)
	SaveAvatar(ID int, filelocation string) (Users, error)
	GetUserById(ID int) (Users, error)
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
func (s *service) IsEmailAvailable(input ChekEmail) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.Id == 0 {
		return true, nil
	}
	return false, nil
}
func (s *service) SaveAvatar(ID int, filelocation string) (Users, error) {
	//dapatkan user dengan id
	//upload avatar file name
	//simpan perubahan
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}
	user.Avatar_file_name = filelocation
	updateuser, err := s.repository.Update(user)
	if err != nil {
		return updateuser, err
	}
	return updateuser, nil
}
func (s *service) GetUserById(ID int) (Users, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("user not found")
	}
	return user, nil
}
