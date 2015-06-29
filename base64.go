package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

func main() {
	hash := sha1.New()
	hash.Write([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	h := strings.Replace(base64.StdEncoding.EncodeToString(hash.Sum(nil)), "=", "", -1)
	fmt.Printf("%s", h)
}
