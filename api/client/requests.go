package client

import (
    "fmt"
    "log"
    "bytes"
    "net/http"
    "io"
    "io/ioutil"
    "encoding/json"
)

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
    body interface{},
    obj interface{},
) {
    payload := new(bytes.Buffer)
    json.NewEncoder(payload).Encode(body)

    c.sendRequest(fmt.Sprintf("basket/%s/product", id), method, payload, obj)
}

func (c *Client) sendProductsRequest(method string, obj interface{}) {
    c.sendRequest("product", method, nil, obj)
}
