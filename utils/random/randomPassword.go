package random

import (
	"bytes"
	"math/rand"
	"time"
)

func GetRandomPassword() string {
	const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&*"
	rand.NewSource(time.Now().UnixNano())
	var s bytes.Buffer
	for i := 0; i < 14; i++ {
		s.WriteByte(charSet[rand.Int63()%int64(len(charSet))])
	}
	return s.String()
}
