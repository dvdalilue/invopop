package product

import (
    "github.com/dvdalilue/invopop/db"
)

func ToProductDto(product *db.Product) *Product {
    return &Product{product.ID, product.Name, product.Price}
}

func ToProductsDto(products []*db.Product) *Products {
    var res []*Product = []*Product{}

    for _, p := range products {
        res = append(res, ToProductDto(p))
    }

    return &Products{res}
}