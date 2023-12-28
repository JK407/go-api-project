package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// EncodePassword 加密密码
func EncodePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println("EncodePasswordERR:", err)
		return "", err
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	return encodePWD, nil
}

func ValidatePassword(passwordFromData, InputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordFromData), []byte(InputPassword))
	if err != nil {
		return false
	}
	return true

}
