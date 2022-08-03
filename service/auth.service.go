package service

import (
	"dansmultipro/recruitment/model"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService interface {
	VerifyCredential(email string, password string)( bool, model.User)
}

type loginService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) LoginService {
	return &loginService{db}
}


func (s *loginService) VerifyCredential(email string, password string) (bool, model.User) {

	var user model.User
	res := s.db.Where("username = ?", email).First(&user)
	if res.Error == nil {		
		comparedPassword := comparePassword(user.Password, []byte(password))
		if comparedPassword {
			return true, user
		}
	}
	return false, user


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
