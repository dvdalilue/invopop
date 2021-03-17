package basket

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/dvdalilue/invopop/db"
    "github.com/dvdalilue/invopop/api/common"
)

// Mapper function to translate a basket model into a friendlier
// DTO. Get the basket/product relations and creates a 'summary'
// object with the list of items and total to pay (with discounts)
func toBasketDto(
    c *gin.Context,
    s db.Store,
    b *db.Basket,
) (*Basket, *db.Error) {
    var prices = make(map[int64]float64)
    var quantities = make(map[int64]int)

    products, err := s.GetBasketProducts(c, b.ID)

    if err != nil {
        return nil, err
    }

    var items []string = []string{}
    var subTotal float64 = 0.0

    for _, p := range products {
        _, exists := quantities[p.ID]

        if exists {
            quantities[p.ID] += 1
        } else {
            prices[p.ID] = p.Price
            quantities[p.ID] = 1
        }

        items = append(items, p.Name)
        subTotal += p.Price
    }

    dm := NewDiscountManager(subTotal, quantities, prices)

    return &Basket{b.ID, items, dm.getTotal()}, nil
}

// Handler function to create a basket. It's assumed that this is
// always successful. The basket model object is mapped to a DTO
// hidding the Store model
func createBasket(s db.Store) func(*gin.Context) {
    handler := func(c *gin.Context) {
        basket := s.CreateBasket(c)

        res, err := toBasketDto(c, s, basket)

        if err != nil {
            c.JSON(err.Code, err.ToAPIResponse())
            return
        }

        c.JSON(http.StatusOK, res)
    }

    return handler
}

// Handler function to get all baskets. The basket model object is
// mapped to a DTO as before
func getBaskets(s db.Store) func(*gin.Context) {
    handler := func(c *gin.Context) {
        baskets := s.GetBaskets(c)

        var res []*Basket = []*Basket{}

        for _, bsk := range baskets {
            basket, err := toBasketDto(c, s, bsk)

            if err != nil {
                c.JSON(err.Code, err.ToAPIResponse())
                return
            }

            res = append(res, basket)
        }

        c.JSON(http.StatusOK, &Baskets{res})
    }

    return handler
}

// Handler function to get a single basket, the 'id' is a
// path parameter extracted in the 'PathIDMiddleware'
func getBasket(s db.Store) func(*gin.Context) {
    handler := func(c *gin.Context) {
        objID := c.MustGet("id").(int64)

        if objID < 0 {
            return
        }

        basket, err := s.GetBasket(c, objID)

        if err != nil {
            c.JSON(err.Code, err.ToAPIResponse())
            return
        }

        res, err := toBasketDto(c, s, basket)

        if err != nil {
            c.JSON(err.Code, err.ToAPIResponse())
            return
        }

        c.JSON(http.StatusOK, res)
    }

    return handler
}

// Handler function to add a product to a basket, the 'id' is a
// path parameter extracted in the 'PathIDMiddleware' and the
// product 'id' is received in the request's body which is checked
// with the 'AddBasketProduct' DTO
func addBasketProduct(s db.Store) func(*gin.Context) {
    handler := func(c *gin.Context) {
        var req AddBasketProduct

        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, common.APIResponse{
                Code: http.StatusBadRequest,
                Message: err.Error(),
            })
            return
        }

        objID := c.MustGet("id").(int64)

        if objID < 0 {
            return
        }

        basket, err := s.AddBasketProduct(c, objID, req.ProductID)

        if err != nil {
            c.JSON(err.Code, err.ToAPIResponse())
            return
        }

        res, err := toBasketDto(c, s, basket)

        if err != nil {
            c.JSON(err.Code, err.ToAPIResponse())
            return
        }

        c.JSON(http.StatusOK, res)
    }

    return handler
}

// Handler function to remove a basket, the 'id' is a path
// parameter extracted in the 'PathIDMiddleware'
func deleteBasket(s db.Store) func(*gin.Context) {
    handler := func(c *gin.Context) {
        objID := c.MustGet("id").(int64)

        if objID < 0 {
            return
        }

        err := s.DeleteBasket(c, objID)

        if err != nil {
            c.JSON(err.Code, err.ToAPIResponse())
            return
        }

        c.JSON(http.StatusNoContent, nil)
    }

    return handler
}

// Includes basket operations in a router based on the prefix
// parameter and pass the store to the handlers
func IncludeOperations(r *gin.Engine, s db.Store, prefix string) {
    basketAPI := r.Group(prefix)

    {
        basketAPI.POST("/", createBasket(s))
        basketAPI.GET("/", getBaskets(s))
        basketAPI.GET("/:id", getBasket(s))
        basketAPI.DELETE("/:id", deleteBasket(s))
        basketAPI.POST("/:id/product", addBasketProduct(s))
    }
}