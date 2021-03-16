package main

import (
    "os"
    "fmt"
    "github.com/dvdalilue/invopop/api"
    "github.com/dvdalilue/invopop/db"
)

func main() {
    mStore := &db.InMemoryStore{}
    mStore.Init()

    mStore.SetProducts([]*db.Product{
        &db.Product{
            ID: 1,
            Name: "PEN",
            Price: 5.00,
        },
        &db.Product{
            ID: 2,
            Name: "TSHIRT",
            Price: 20.00,
        },
        &db.Product{
            ID: 3,
            Name: "MUG",
            Price: 7.50,
        },
    })

    server := api.NewServer(mStore)

    port := ":8080"

    envPort, ok := os.LookupEnv("INVOPOP_SERVER_PORT")

    if ok {
        port = fmt.Sprintf(":%s", envPort)
    }

    // listen and serve on 0.0.0.0:8080
    server.Listen(port)
}