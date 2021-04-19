package service

import (
	"github.com/mashingan/smapping"
	"helloGinAndGorm/dto"
	"helloGinAndGorm/entity"
	"helloGinAndGorm/repository"
	"log"
)

// UserService is a contact.....
type UserService interface {
	Update(user dto.UserUpdateDto) entity.User
	Profile(userId string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService {
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDto) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	return service.userRepository.UpdateUser(userToUpdate)
}

func (service *userService) Profile(userId string) entity.User {
	return service.userRepository.ProfileUser(userId)
}

