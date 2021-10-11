package utils

import "golang.org/x/crypto/bcrypt"

// 加密
func CreatePw(password string) string {
	hashPw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPw)
}

// 加密验证 一致true 不一致false
func CheckPw(hashPw string, password string) bool {
	pwErr := bcrypt.CompareHashAndPassword([]byte(hashPw), []byte(password))
	return pwErr == nil
}
