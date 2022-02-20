package services

import (
	"log"
	"yesiamdonation/entities"
	"yesiamdonation/repositories"
	"yesiamdonation/services/request"

	"github.com/mashingan/smapping"
)

type AuthService interface {
	Register(user request.RegisterRequest) entities.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	user repositories.UserRepository
}

func NewAuthService(user repositories.UserRepository) AuthService {
	return &authService{
		user: user,
	}
}

func (service *authService) Register(user request.RegisterRequest) entities.User {
	userCreate := entities.User{}

	err := smapping.FillStruct(&userCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed map %v", err)
	}
	res := service.user.Register(userCreate)
	return res
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.user.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
