package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	_ "github.com/joho/godotenv"
)

type JwtService interface {
	GenerateToken(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {}

type jwtCustomClaim struct {
	UserId string `json:"userId"`
	*jwt.StandardClaims
}

func (s *jwtService) GenerateToken(userId string) string {
	claims := &jwtCustomClaim{
		userId,
		&jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}
	return t
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("UNEXPECTED SIGNING METHOD %v", t.Header["alg"])	
		}
		return []byte(os.Getenv("JWT_SECRET")), nil 
	})
}

func NewJwtService() JwtService {
	return &jwtService{}
}
