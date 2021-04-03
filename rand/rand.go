package rand

import (
	"math/rand"
	"time"
)

// RandomSeq returns a sequence of length n using from the chars characters randomly
func RandomSeq(n int, chars []rune) string {
	rand.Seed(time.Now().UnixNano())

	charLen := len(chars)

	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(charLen)]
	}
	return string(b)
}

// RandBool returns a random boolean value
func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
