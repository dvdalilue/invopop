package basket

type Basket struct {

    // id
    ID int64 `json:"id"`

    // items
    Items []string `json:"items"`

    // total
    Total float64 `json:"total"`
}

type Baskets struct {

	// items
    Items []*Basket `json:"baskets"`
}


type AddBasketProduct struct {

    // product_id
    ProductID int64 `json:"product_id" binding:"required"`
}
