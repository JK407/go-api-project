package utils_test

import (
	"go_project/internal/utils"
	"testing"
)

func Test_EncodePassword(t *testing.T) {
	password := "123456"
	encodePWD, err := utils.EncodePassword(password)
	if err != nil {
		t.Error(err)
	}
	t.Log(encodePWD)
}

func Test_ValidatePassword(t *testing.T) {
	password := "123456"
	encodePWD, err := utils.EncodePassword(password)
	if err != nil {
		t.Error(err)
	}
	if !utils.ValidatePassword(encodePWD, password) {
		t.Error("密码验证失败")
	}
}
