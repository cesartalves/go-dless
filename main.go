package main

import (
	"fmt"

	"github.com/cesartalves/go-dless/rand"
)

func main() {
	fmt.Println(rand.RandomSeq(1, []rune("ABC")))
}
