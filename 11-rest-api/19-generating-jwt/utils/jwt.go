package utils

//Pacote de token de autenticação - golang jwt
import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

// Configurando token
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ //tipo de criptografia
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), //expira em 2 horas
	})

	return token.SignedString([]byte(secretKey))
}
