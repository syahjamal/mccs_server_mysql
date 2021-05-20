package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/syahjamal/mccs_server_mysql/dto"
	"github.com/syahjamal/mccs_server_mysql/model"
	"github.com/syahjamal/mccs_server_mysql/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(nik string, password string) interface{}
	CreateUser(user dto.RegisterDTO) model.MccsUser
	FindByNIK(nik string) model.MccsUser
	IsDuplicateNIK(nik string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(nik string, password string) interface{} {
	res := service.userRepository.VerifyCredential(nik, password)
	if v, ok := res.(model.MccsUser); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.NIK == nik && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) model.MccsUser {
	userToCreate := model.MccsUser{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByNIK(nik string) model.MccsUser {
	return service.userRepository.FindByNIK(nik)
}

func (service *authService) IsDuplicateNIK(nik string) bool {
	res := service.userRepository.IsDuplicateNIK(nik)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
