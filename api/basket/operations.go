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
    products, err := s.GetBasketProducts(c)

    if err != nil {
        return nil, err
    }

    var items []string = []string{}
    var total float64 = 0.0

    for _, p := range products {
        items = append(items, p.Name)
        total += p.Price
    }

    return &Basket{items, total}, nil
}

func createBasket(s db.Store) func(*gin.Context) {
    handler := func(c *gin.Context) {
        basket, err := s.CreateBasket(c)

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

func getBasket(s db.Store) func(*gin.Context) {
    handler := func(c *gin.Context) {
        basket, err := s.GetBasket(c)

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

        basket, err := s.AddBasketProduct(c, req.ProductID)

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
        err := s.DeleteBasket(c)

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
        basketAPI.GET("/", getBasket(s))
        basketAPI.DELETE("/", deleteBasket(s))
        basketAPI.POST("/product", addBasketProduct(s))
    }
}