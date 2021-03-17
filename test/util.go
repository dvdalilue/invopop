package test

import (
    "time"
    "math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomInit() {
    rand.Seed(time.Now().UnixNano())
}

func RandomInt(min int64, max int64) int64 {
    return min + rand.Int63n(max - min + 1)
}

func RandomFloat(min float64, max float64) float64 {
    return min + rand.Float64() * (max - min)
}

func RandomString(n int) string {
    b := make([]rune, n)
    k := len(letters)

    for i := range b {
        b[i] = letters[rand.Intn(k)]
    }

    return string(b)
}