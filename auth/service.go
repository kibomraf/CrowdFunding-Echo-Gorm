package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"

	"crowdfunding/helper"
	"crowdfunding/users"
)

//...

type AuthService interface {
	GenerateToken(userId int) (string, error)
}
type jwtservice struct {
	secretkey string
}

func NewJWTservice(secretkey string) *jwtservice {
	return &jwtservice{secretkey}
}
func (s *jwtservice) GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretkey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func (s *jwtservice) ValidationToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("error")
		}
		return []byte(s.secretkey), nil
	})
}
func (s *jwtservice) JWTMiddleware(auth AuthService, userService users.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if !strings.Contains(authHeader, "Bearer") {
				response := helper.APIResponse("unathorized1", echo.ErrUnauthorized.Code, "error", nil)
				return c.JSON(echo.ErrUnauthorized.Code, response)
			}
			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
			token, err := s.ValidationToken(tokenString)
			if err != nil {
				response := helper.APIResponse("unathorized2", echo.ErrUnauthorized.Code, "error", token)
				return c.JSON(echo.ErrUnauthorized.Code, response)
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				response := helper.APIResponse("unathorized3", echo.ErrUnauthorized.Code, "error", nil)
				return c.JSON(echo.ErrUnauthorized.Code, response)
			}
			userId, ok := claims["user_id"].(float64)
			if !ok {
				response := helper.APIResponse("unathorized4", echo.ErrUnauthorized.Code, "error", nil)
				return c.JSON(echo.ErrUnauthorized.Code, response)
			}
			user, err := userService.GetUserById(int(userId))
			if err != nil {
				response := helper.APIResponse("unathorized5", echo.ErrUnauthorized.Code, "error", nil)
				return c.JSON(echo.ErrUnauthorized.Code, response)
			}
			c.Set("user", user)
			return next(c)
		}
	}
}
