package service

import (
	"dansmultipro/recruitment/model"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(user model.User) string
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTClaim struct{
	jwt.StandardClaims
	Username string `json:"username"`
}

type jwtService struct{
	secretKey string
	issuer string
}

func NewJWTService() *jwtService{
	return &jwtService{
		issuer: "nvlhnn",
		secretKey: getSecretKey(),
	}
}


func (s *jwtService) GenerateToken(user model.User) string{
	claims := &JWTClaim{
		Username: user.Password,
		StandardClaims: jwt.StandardClaims{
			Issuer:    s.issuer,
			IssuedAt: time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	t,err := token.SignedString([]byte(s.secretKey))
	if err!=nil{
		panic(err)
	}

	return t
} 


func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}


func getSecretKey() string {
	secreKey := os.Getenv("JWT_SECRET")
	if secreKey== "" {
		secreKey = "nvlhnn"
	}
	return secreKey
}