package product

type Product struct {

    // id
    ID int64 `json:"id"`

    // name
    Name string `json:"name"`

    // price
    Price float64 `json:"price"`
}

type Products struct {

    // items
    Items []*Product `json:"items"`
}
