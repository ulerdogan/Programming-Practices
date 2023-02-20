package utils

import (
	"math/rand"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var roles = []string{"guard", "forward", "center"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GetRandomRole() string {
	return roles[rand.Intn(3)]
}