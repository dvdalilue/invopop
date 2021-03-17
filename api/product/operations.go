package product

import (
	"net/http"
    "github.com/gin-gonic/gin"
    "github.com/dvdalilue/invopop/db"
)

func getProducts(s db.Store) func(*gin.Context) {
    handler := func(c *gin.Context) {
        products := s.GetProducts(c)

        res := ToProductsDto(products)

        c.JSON(http.StatusOK, res)
    }

    return handler
}

func IncludeOperations(r *gin.Engine, s db.Store, prefix string) {
    productAPI := r.Group(prefix)

    {
        productAPI.GET("/", getProducts(s))
    }
}