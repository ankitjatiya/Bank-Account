package db

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	l := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(1, 10000)
}

func RandomCurrency() string {
	currencies := []string{"INR", "USD", "CAD", "EUR"}
	l := len(currencies)
	return currencies[rand.Intn(l)]
}
