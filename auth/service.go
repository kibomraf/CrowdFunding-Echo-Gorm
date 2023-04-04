package auth

import "github.com/golang-jwt/jwt/v5"

type Service interface {
	GenerateToken(userID int) (string, error)
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
	return signtkn, err
}
