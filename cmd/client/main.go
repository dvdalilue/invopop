package main

import (
    "os"
    "github.com/dvdalilue/invopop/api/client"
)

func main() {
    url := "http://localhost:8080"

    envUrl, ok := os.LookupEnv("INVOPOP_SERVER_URL")

    if ok {
        url = envUrl
    }

    client := client.NewClient(url)

    client.Cli()
}