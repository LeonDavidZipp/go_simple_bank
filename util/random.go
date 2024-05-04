package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomStr(len int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < len; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomStr(RandomInt(6, 12))
}

func RandomBalance() int {
	return RandomInt(0, 50000)
}

func RandomCurrency() string {
	return [3]string{"USD", "EUR", "CAD"}[RandomInt(0, 2)]
}