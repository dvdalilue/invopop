package basket

func twoOneDiscount(n int, price float64) float64 {
    return float64(n / 2 + n % 2) * price
}

func bulkDiscount(n int, price float64) float64 {
    var finalPrice = price

    if n > 2 {
        finalPrice = price * 0.75
    }

    return float64(n) * finalPrice
}

// The idea behind this struct is to define a list of discounts
// that any product can have and compute the total based the
// subtotal, quantities and prices. This is more like a
// 'CheckoutManager'
type DiscountManager struct {
    subTotal float64
    quantities map[int64]int
    prices map[int64]float64
    discountMap map[int64][]func(int, float64) float64
}

func NewDiscountManager(
    subTotal float64,
    quantities map[int64]int,
    prices map[int64]float64,
) *DiscountManager {
    return &DiscountManager{
        quantities: quantities,
        prices: prices,
        subTotal: subTotal,
        discountMap: map[int64][]func(int, float64) float64{
            1: {twoOneDiscount},
            2: {bulkDiscount},
        },
    }
}

// Applies all the discount and gets the final total
func (dm *DiscountManager) getTotal() float64 {
    var total float64 = dm.subTotal
    var original float64
    var reduced float64

    for id, discounts := range dm.discountMap {
        quantity, inBasket := dm.quantities[id]

        if !inBasket { continue; }

        price := dm.prices[id]

        for _, discountFunc := range discounts {
            original = float64(quantity) * price
            reduced = discountFunc(quantity, price)

            total -= original - reduced
        }
    }

    return total
}