package rand

import (
	"math/rand"
	"time"
)

func RandomSeq(n int, chars []rune) string {
	rand.Seed(time.Now().UnixNano())

	charLen := len(chars)

	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Intn(charLen)]
	}
	return string(b)
}
