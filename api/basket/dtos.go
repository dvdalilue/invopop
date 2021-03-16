package basket

type Basket struct {

    // items
    Items []string `json:"items"`

    // total
    Total float64 `json:"total"`
}


type AddBasketProduct struct {

    // product_id
    ProductID int64 `json:"product_id" binding:"required"`
}
