package common

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/hamasfaa/project-evermos/configuration"
	"github.com/hamasfaa/project-evermos/exception"
)

func GenerateToken(noTelp string, isAdmin bool, userID int, config configuration.Config) string {
	jwtSecret := config.Get("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(config.Get("JWT_EXPIRE_MINUTES_COUNT"))
	exception.PanicLogging(err)

	claims := jwt.MapClaims{
		"userID":   userID,
		"noTelp":   noTelp,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned
}
