package db

import (
    "net/http"
    "context"
)

type Store interface {
    Querier
}

type Pair struct {
    basketID int64
    productID int64
}

type InMemoryStore struct {
    baskets []*Basket
    products []*Product
    basketProducts []*Pair

    basketSeq int64
}

func (ms *InMemoryStore) SetProducts(ps []*Product) {
    ms.products = ps
}

func (ms *InMemoryStore) Init() {
    ms.baskets = []*Basket{}
    ms.products = []*Product{}
    ms.basketProducts = []*Pair{}
    ms.basketSeq = 1
}

// Interface methods

func (ms *InMemoryStore) CreateBasket(ctx context.Context) *Basket {
    basketID := ms.basketSeq

    ms.basketSeq += 1

    basket := &Basket{ID: basketID, Name:"default"}

    ms.baskets = append(ms.baskets, basket)

    return basket
}

func (ms *InMemoryStore) GetBaskets(ctx context.Context) []*Basket {
    if ms.baskets == nil {
        return []*Basket{}
    }

    return ms.baskets
}

func (ms *InMemoryStore) getBasketIndex(
    ctx context.Context,
    id int64,
) (*Basket, int, *Error) {
    if len(ms.baskets) <= 0 {
        return nil, -1, &Error{
            Code: http.StatusNotFound,
            Message: "There are no baskets",
        }
    }

    for idx, value := range ms.baskets {
        if value.ID == id {
            return value, idx, nil
        }
    }

    return nil, -1, &Error{
        Code: http.StatusNotFound,
        Message: "Basket not found",
    }
}

func (ms *InMemoryStore) GetBasket(
    ctx context.Context,
    id int64,
) (*Basket, *Error) {
    basket, _, err := ms.getBasketIndex(ctx, id)

    return basket, err
}

func (ms *InMemoryStore) DeleteBasket(ctx context.Context, id int64) *Error {
    _, idx, err := ms.getBasketIndex(ctx, id)

    if err != nil {
        return err
    }

    ms.baskets = append(ms.baskets[:idx], ms.baskets[idx+1:]...)

    var newPairs []*Pair

    for _, pair := range ms.basketProducts {
        if pair.basketID != id {
            newPairs = append(newPairs, pair)
        }
    }

    ms.basketProducts = newPairs

    return nil
}

func (ms *InMemoryStore) GetProducts(ctx context.Context) []*Product {
    if ms.products == nil {
        return []*Product{}
    }

    return ms.products
}

func (ms *InMemoryStore) GetProduct(
    ctx context.Context,
    id int64,
) (*Product, *Error) {
    notFound := &Error{
        Code: http.StatusNotFound,
        Message: "Product not found",
    }

    if ms.products == nil {
        return nil, notFound
    }

    for _, value := range ms.products {
        if value.ID == id {
            return value, nil
        }
    }

    return nil, notFound
}

func (ms *InMemoryStore) AddBasketProduct(
    ctx context.Context,
    basketID int64,
    productID int64,
) (*Basket, *Error) {
    basket, err := ms.GetBasket(ctx, basketID)

    if err != nil {
        return nil, err
    }

    product, err := ms.GetProduct(ctx, productID)

    if err != nil {
        return nil, err
    }

    ms.basketProducts = append(ms.basketProducts, &Pair{
        basketID: basket.ID,
        productID: product.ID,
    })

    return basket, nil
}

func (ms *InMemoryStore) GetBasketProducts(
    ctx context.Context,
    basketID int64,
) ([]*Product, *Error) {
    basket, err := ms.GetBasket(ctx, basketID)

    if err != nil {
        return nil, err
    }

    var res []*Product

    for _, value := range ms.basketProducts {
        if value.basketID != basket.ID {
            continue
        }

        product, err := ms.GetProduct(ctx, value.productID)

        if err != nil {
            return nil, err
        }

        res = append(res, product)
    }

    return res, nil
}
