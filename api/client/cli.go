package client

import (
    "os"
    "fmt"
    "time"
    "bufio"
    "strings"
    "net/http"
)

type Client struct {
    BaseURL    string
    HTTPClient *http.Client
}

func NewClient(baseURL string) *Client {
    return &Client{
        BaseURL: baseURL,
        HTTPClient: &http.Client{
            Timeout: time.Minute,
        },
    }
}

func (c *Client) Cli() {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Printf("(h for help) ? ")
        text, _ := reader.ReadString('\n')

        cmd := strings.Fields(text)

        if len(cmd) <= 0 { continue; }

        switch cmd[0] {
        case "h":
            fmt.Printf("Available commands: help, new, list, items, add, del, q\n")
        case "new":
            c.createBasket()
        case "list":
            c.listBaskets()
        case "items":
            c.listProducts()
        case "add":
            if len(cmd) <= 2 {
                fmt.Printf("Missing argument, 2 were expected (basket id and product id).\n")
                continue;
            }
            c.addProduct(cmd[1], cmd[2])
        case "del":
            if len(cmd) <= 1 {
                fmt.Printf("Missing argument, 1 was expected (basket id).\n")
                continue;
            }
            c.deleteBasket(cmd[1])
        case "q":
            os.Exit(0)
        default:
            fmt.Printf("Unknown command: %s\n", cmd[0])
        }
    }
}