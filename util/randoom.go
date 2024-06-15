package util

import (
	"math/rand"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	println("Seeding random number generator...")
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandOwnersName() string {
	return RandString(6)
}

func RandAmount() int64 {
	return RandInt(0, 1000)
}

func RandCurrency() string {
	return AllCurrencies()[rand.Intn(len(AllCurrencies()))]
}

func RandEmail() string {
	return RandString(6) + "@gmail.com"
}
