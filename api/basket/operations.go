package basket

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/dvdalilue/invopop/db"
    "github.com/dvdalilue/invopop/api/common"
)

func toBasketDto(
    c *gin.Context,
    s db.Store,
    b *db.Basket,
) (*Basket, *db.Error) {
    products, err := s.GetBasketProducts(c, b.ID)

    if err != nil {
        return nil, err
    }

    var items []string = []string{}
    var total float64 = 0.0

    for _, p := range products {
        items = append(items, p.Name)
        total += p.Price
    }

    return &Basket{b.ID, items, total}, nil
}

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