package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(random(16))
	fmt.Println(random(16))
	fmt.Println(random(16))
}

func random(length int) string {
	const base = 36
	size := big.NewInt(base)
	n := make([]byte, length)
	for i, _ := range n {
		c, _ := rand.Int(rand.Reader, size)
		n[i] = strconv.FormatInt(c.Int64(), base)[0]
	}
	return string(n)
}
