package user

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	RegisterAdmin(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
	VerificationUser(input VerificationInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	rand.Seed(time.Now().Unix())
	charSet := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
	length := 10
	var output strings.Builder
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	user.Code = output.String()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"
	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
func (s *service) RegisterAdmin(input RegisterUserInput) (User, error) {
	admin := User{}
	admin.Name = input.Name
	admin.Email = input.Email
	admin.Occupation = input.Occupation

	rand.Seed(time.Now().Unix())
	charSet := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
	length := 10
	var output strings.Builder
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	admin.Code = output.String()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return admin, err
	}
	admin.PasswordHash = string(passwordHash)
	admin.Role = "admin"
	newAdmin, err := s.repository.Save(admin)

	if err != nil {
		return newAdmin, err
	}
	return newAdmin, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}

	if user.Verif == 0 {
		return user, errors.New("verifikasi dahulu")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if user.Verif == 0 {
		return user, errors.New("verifikasi dahulu")
	}
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if user.Verif == 0 {
		return user, errors.New("verifikasi dahulu")
	}
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found")
	}

	return user, nil
}

func (s *service) VerificationUser(input VerificationInput) (User, error) {
	user, err := s.repository.FindByCode(input.Code)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found")
	}

	if input.Code != user.Code {
		return user, errors.New("invalid")
	}

	user.Verif = 1
	newVerif, err := s.repository.Update(user)

	return newVerif, nil
}
