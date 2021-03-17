package client

import (
    "fmt"
    "log"
    "strconv"
    "net/http"

    "github.com/dvdalilue/invopop/api/basket"
    "github.com/dvdalilue/invopop/api/product"
)

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
        fmt.Printf("ID: %d, Name: %s, Price: %.2f€\n",
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
