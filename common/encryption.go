package common

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func PasswordCry(password string) []byte {
	userPassword := "123456"
	newPassword, err := GeneratePassword(userPassword)
	if err != nil {
		fmt.Println("加密出错了")
		log.Fatal("加密出错了")
		return nil
	}
	return newPassword
}

// GeneratePassword 给密码就行加密操作
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// ValidatePassword 密码比对
func ValidatePassword(userPassword string, hashed string) (isOK bool, err error) {
	newPwd := PasswordCry(userPassword)
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(newPwd)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil
}
