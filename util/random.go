package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomStr(strlen int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < strlen; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomStr(int(RandomInt(6, 12)))
}

func RandomBalance() int64 {
	return RandomInt(0, 50000)
}

func RandomCurrency() string {
	return [3]string{"USD", "EUR", "CAD"}[int(RandomInt(0, 2))]
}

func RandomAmount() int64 {
	return RandomInt(0, 50000)
}

func RandomID() int64 {
	return RandomInt(0, 50000)
}
