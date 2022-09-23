package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmopqrstvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomNumber(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func GenerateRandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		letter := alphabet[rand.Intn(k)]
		sb.WriteByte(letter)
	}
	return sb.String()
}

func GenerateRandomOwner() string {
	return GenerateRandomString(6)
}

func GenerateRandomMoney() int64 {
	return GenerateRandomNumber(0, 1000)
}

func GenerateRandomCurrency() string {
	currencies := []string{"USD", "EUR", "INR"}
	return currencies[rand.Intn(len(currencies))]
}
