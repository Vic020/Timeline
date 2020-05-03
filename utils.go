package main

import (
	"math/rand"
	"time"
)

func NewRandString() func(int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return func(l int) string {
		bytes := make([]byte, l)
		for i := 0; i < l; i++ {
			bytes[i] = byte(r.Intn(26) + 65)
		}
		return string(bytes)
	}
}
