package client

import (
    "fmt"
    "strings"

    "github.com/dvdalilue/invopop/api/basket"
)

func printBasket(bsk *basket.Basket) {
    fmt.Printf(
        "- ID: %d\n  Items: %s\n  Total: %.2f\n\n",
        bsk.ID,
        strings.Trim(fmt.Sprint(bsk.Items), "[]"),
        bsk.Total,
    )
}