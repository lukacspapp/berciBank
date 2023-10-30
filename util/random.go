package util

import (
	"math/rand"
	"strings"
	"time"
)

const abc = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // return a random int between 0 - (max-min)
}

func RandomSting(n int) string {
	var sb strings.Builder
	k := len(abc)

	for i := 0; i < n; i++ {
		c := abc[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomSting(6)
}

func RandomAmount() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "GBP", "USD", "HUF"}

	n := len(currencies)
	return currencies[rand.Intn(n)]
}
