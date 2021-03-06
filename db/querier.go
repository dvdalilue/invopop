// Code generated by sqlc. DO NOT EDIT.

package db

import (
    "context"
)

type Querier interface {
	CreateBasket(ctx context.Context) *Basket
	GetBaskets(ctx context.Context) []*Basket
	GetBasket(ctx context.Context, basketID int64) (*Basket, *Error)
	DeleteBasket(ctx context.Context, basketID int64) *Error
	GetProducts(ctx context.Context) []*Product
	GetProduct(ctx context.Context, id int64) (*Product, *Error)
	AddBasketProduct(ctx context.Context, basketID int64, productID int64) (*Basket, *Error)
	GetBasketProducts(ctx context.Context, basketID int64) ([]*Product, *Error)
}
