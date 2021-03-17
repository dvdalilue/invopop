package product

import (
	"net/http"
    "github.com/gin-gonic/gin"
    "github.com/dvdalilue/invopop/db"
)

// Handler function to get all the products. It's assumed that this
// is always successful. The products are mapped to a DTO
func getProducts(s db.Store) func(*gin.Context) {
    handler := func(c *gin.Context) {
        products := s.GetProducts(c)

        res := ToProductsDto(products)

        c.JSON(http.StatusOK, res)
    }

    return handler
}

// Includes product operations in a router based on the prefix
// parameter and pass the store to the handlers
func IncludeOperations(r *gin.Engine, s db.Store, prefix string) {
    productAPI := r.Group(prefix)

    {
        productAPI.GET("/", getProducts(s))
    }
}