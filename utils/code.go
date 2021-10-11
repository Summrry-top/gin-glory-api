package utils

import (
	"math/rand"
	"time"
)

// code字符
const CodeChars = "0123456789"

//
func CreateCode(size int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, size)
	for i := 0; i < size; i++ {
		result[i] = CodeChars[rand.Intn(len(CodeChars))]
	}
	return string(result)
}
