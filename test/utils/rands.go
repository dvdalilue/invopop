package utils

import (
    "time"
    "math/rand"

    "github.com/dvdalilue/invopop/db"
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

func RandomBasket() db.Basket {
    return db.Basket{
        ID: RandomInt(1, 1000),
    }
}

func RandomBaskets(n int64) []*db.Basket {
    baskets := make([]*db.Basket, n)

    var p db.Basket

    for i := int64(0); i < n; i++ {
        p = RandomBasket()
        baskets[i] = &p
    }

    return baskets
}

func RandomProduct() db.Product {
    return db.Product{
        ID:    RandomInt(1, 1000),
        Name:  RandomString(10),
        Price: RandomFloat(2.0, 42.0),
    }
}

func RandomProducts(n int64) []*db.Product {
    products := make([]*db.Product, n)

    var p db.Product

    for i := int64(0); i < n; i++ {
        p = RandomProduct()
        products[i] = &p
    }

    return products
}
