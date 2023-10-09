package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("NHvbW`Yw3BgPec>=!{T-)D")

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Second * 30).Unix(), // Token expiration time
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func EncryptString(password string) (string, error) {
    // Generate a salt and hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }

    return string(hashedPassword), nil
}

func DecryptString(hashedPassword string, password string) error {
    // Compare the hashed password and the plaintext password
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err
}
