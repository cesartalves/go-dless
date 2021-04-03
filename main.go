package main

import (
	"fmt"

	"github.com/cesartalves/go-dless/rand"
)

func main() {
	fmt.Println(rand.RandomSeq(1, []rune("ABC")))

	n := 100000.0
	trues, falses := 0.0, 0.0

	for i := 0.0; i < n; i++ {
		if rand.RandBool() {
			trues++
			continue
		}

		falses++
	}

	fmt.Println(trues/n, falses/n)
}
