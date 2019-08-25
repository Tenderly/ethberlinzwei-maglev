package helper

import (
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("0123456789abcde")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
