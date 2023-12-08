package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	k := len(alphabets)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return randomString(6)
}

func RandomBalance() int64 {
	return randInt(0, 1000)
}

func RandomCurrency() string {
	var currency = []string{"USD", "CAD", "EUR"}
	n := len(currency)

	return currency[rand.Intn(n)]
}
