package main

import (
    "os"
    "fmt"
    "github.com/dvdalilue/invopop/api"
    "github.com/dvdalilue/invopop/db"
)

func main() {
    mStore := db.NewInMemoryStore()

    server := api.NewServer(mStore)

    port := ":8080"

    envPort, ok := os.LookupEnv("INVOPOP_SERVER_PORT")

    if ok {
        port = fmt.Sprintf(":%s", envPort)
    }

    // listen and serve on 0.0.0.0:8080
    server.Listen(port)
}