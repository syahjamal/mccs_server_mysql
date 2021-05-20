package repository

import (
	"log"

	"github.com/syahjamal/mccs_server_mysql/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository adalah kontrak apa yang akan userRepository yang akan lakukan
type UserRepository interface {
	InsertUser(user model.MccsUser) model.MccsUser
	UpdateUser(user model.MccsUser) model.MccsUser
	VerifyCredential(nik string, password string) interface{}
	IsDuplicateNIK(nik string) (tx *gorm.DB)
	FindByNIK(nik string) model.MccsUser
	ProfileUser(userID string) model.MccsUser
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user model.MccsUser) model.MccsUser {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user model.MccsUser) model.MccsUser {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(nik string, password string) interface{} {
	var user model.MccsUser
	res := db.connection.Where("nik = ?", nik).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateNIK(nik string) (tx *gorm.DB) {
	var user model.MccsUser
	return db.connection.Where("nik = ?", nik).Take(&user)
}

func (db *userConnection) FindByNIK(nik string) model.MccsUser {
	var user model.MccsUser
	db.connection.Where("nik = ?", nik).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) model.MccsUser {
	var user model.MccsUser
	db.connection.Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
