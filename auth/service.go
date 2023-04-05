package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(tkn string) (*jwt.Token, error)
}
type jwtservice struct {
}

func AuthService() *jwtservice {
	return &jwtservice{}
}

// just for learning
var secretKey = []byte("CrowdFunding-Echo")

func (s *jwtservice) GenerateToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signtkn, err := tkn.SignedString(secretKey)
	if err != nil {
		return signtkn, err
	}
	return signtkn, nil
}
func (s *jwtservice) ValidateToken(token string) (*jwt.Token, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return tkn, err
	}
	return tkn, nil
}
