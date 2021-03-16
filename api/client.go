package api

import (
    "os"
    "fmt"
    "log"
    "time"
    "bufio"
    "bytes"
    "strings"
    "strconv"
    "net/http"
    "io"
    "io/ioutil"
    "encoding/json"
    "github.com/dvdalilue/invopop/api/basket"
    "github.com/dvdalilue/invopop/api/product"
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

func printBasket(bsk *basket.Basket) {
    fmt.Printf(
        "- ID: %d\n  Items: %s\n  Total: %.2f\n\n",
        bsk.ID,
        strings.Trim(fmt.Sprint(bsk.Items), "[]"),
        bsk.Total,
    )
}

func (c *Client) sendRequest(
    path string,
    method string,
    body io.Reader,
    obj interface{},
) {
    req, err := http.NewRequest(
        method,
        fmt.Sprintf("%s/%s", c.BaseURL, path),
        body,
    )

    if err != nil {
        log.Fatalf(err.Error())
        return
    }

    res, err := c.HTTPClient.Do(req)

    if err != nil {
        log.Fatalf(err.Error())
        return
    }

    responseBytes, err := ioutil.ReadAll(res.Body)

    if err != nil {
        log.Fatalf(err.Error())
        return
    }

    if len(responseBytes) <= 0 {
        return
    }

    if err := json.Unmarshal(responseBytes, obj); err != nil {
        log.Fatalf("error deserializing data")
    }
}

func (c *Client) sendBasketsRequest(method string, obj interface{}) {
    c.sendRequest("basket", method, nil, obj)
}

func (c *Client) sendBasketRequest(
    id string,
    method string,
    obj interface{},
) {
    c.sendRequest(fmt.Sprintf("basket/%s", id), method, nil, obj)
}

func (c *Client) sendBasketProductRequest(
    id string,
    method string,
    body *basket.AddBasketProduct,
    obj interface{},
) {
    payload := new(bytes.Buffer)
    json.NewEncoder(payload).Encode(body)

    c.sendRequest(fmt.Sprintf("basket/%s/product", id), method, payload, obj)
}

func (c *Client) sendProductsRequest(method string, obj interface{}) {
    c.sendRequest("product", method, nil, obj)
}

func (c *Client) createBasket() {
    var bsk basket.Basket

    c.sendBasketsRequest(http.MethodPost, &bsk)

    printBasket(&bsk)
}

func (c *Client) listBaskets() {
    var bs basket.Baskets

    c.sendBasketsRequest(http.MethodGet, &bs)

    for _, b := range bs.Items {
        printBasket(b)
    }
}

func (c *Client) deleteBasket(basketID string) {
    c.sendBasketRequest(basketID, http.MethodDelete, struct{}{})
}

func (c *Client) listProducts() {
    var ps product.Products

    c.sendProductsRequest(http.MethodGet, &ps)

    for _, p := range ps.Items {
        fmt.Printf("ID: %d, Name: %s, Price: %.2f\n",
            p.ID,
            p.Name,
            p.Price,
        )
    }
}

func (c *Client) addProduct(basketID string, productID string) {
    var bsk basket.Basket

    pId, err := strconv.Atoi(productID)

    if err != nil {
        log.Fatalf(err.Error())
        return
    }

    body := &basket.AddBasketProduct{ProductID: int64(pId)}

    c.sendBasketProductRequest(basketID, http.MethodPost, body, &bsk)

    printBasket(&bsk)
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