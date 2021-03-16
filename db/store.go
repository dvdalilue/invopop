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
    basket *Basket
    products []*Product
    basketProducts []*Pair
}

func (ms *InMemoryStore) SetProducts(ps []*Product) {
    ms.products = ps
}

func (ms *InMemoryStore) Init() {
    ms.products = []*Product{}
    ms.basketProducts = []*Pair{}
}

// Interface methods

func (ms *InMemoryStore) CreateBasket(ctx context.Context) (*Basket, *Error) {
    if ms.basket != nil {
        return ms.basket, &Error{
            Code: http.StatusConflict,
            Message: "The basket is already created",
        }
    }

    ms.basket = &Basket{ID:1, Name:"default"}

    return ms.basket, nil
}

func (ms *InMemoryStore) GetBasket(ctx context.Context) (*Basket, *Error) {
    if ms.basket == nil {
        return nil, &Error{
            Code: http.StatusNotFound,
            Message: "There is no basket",
        }
    }

    return ms.basket, nil
}

func (ms *InMemoryStore) DeleteBasket(ctx context.Context) *Error {
    if ms.basket == nil {
        return &Error{
            Code: http.StatusNotFound,
            Message: "There is no basket to delete",
        }
    }

    ms.basket = nil
    ms.basketProducts = []*Pair{}

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
    productID int64,
) (*Basket, *Error) {
    basket, err := ms.GetBasket(ctx)

    if err != nil {
        return nil, err
    }

    product, err := ms.GetProduct(ctx, productID)

    if err != nil {
        return nil, err
    }

    ms.basketProducts = append(ms.basketProducts, &Pair{
        basketID: ms.basket.ID,
        productID: product.ID,
    })

    return basket, nil
}

func (ms *InMemoryStore) GetBasketProducts(
    ctx context.Context,
) ([]*Product, *Error) {
    basket, err := ms.GetBasket(ctx)

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
