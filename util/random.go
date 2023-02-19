package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString will generate a random string
func RandomString(n int) string {
	var sb strings.Builder

	k := len(ALPHABET)

	for i := 0; i < n; i++ {
		c := ALPHABET[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail() string {
	email := fmt.Sprintf("%s@gmail.com", RandomString(30))
	return email
}
