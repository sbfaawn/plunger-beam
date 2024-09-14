package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordEncryptor interface {
	Encrypt(password string) (string, error)
	IsHashedPasswordMatch(hashedPassword string, password string) bool
}

type passwordEncryptorImpl struct {
}

func NewPasswordEncryptor() PasswordEncryptor {
	return &passwordEncryptorImpl{}
}

func (encryptor *passwordEncryptorImpl) Encrypt(password string) (string, error) {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hashedPasswordBytes), nil
}

func (encryptor *passwordEncryptorImpl) IsHashedPasswordMatch(hashedPassword string, password string) bool {
	byteHash := []byte(hashedPassword)
	byteOriginal := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHash, byteOriginal)
	if err != nil {
		fmt.Println("hashing matcher error : ", err)
		return false
	}

	return true
}
