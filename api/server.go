package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/dvdalilue/invopop/api/basket"
    "github.com/dvdalilue/invopop/api/product"
    "github.com/dvdalilue/invopop/api/common"
    "github.com/dvdalilue/invopop/db"
)

type Server struct {

    // DB
    store db.Store
    router *gin.Engine
}

func ObjectIDMiddleware(c *gin.Context) {
    var obj common.Object

    if err := c.ShouldBindUri(&obj); err != nil {
        c.JSON(http.StatusBadRequest, common.APIResponse{
            Code: http.StatusBadRequest,
            Message: err.Error(),
        })
        c.Set("id", int64(-1))
        c.Next()
    }

    c.Set("id", obj.ID)
    c.Next()
}

func NewServer(store db.Store) *Server {
    router := gin.Default()

    router.Use(ObjectIDMiddleware)

    router.NoRoute(func(c *gin.Context) {
        c.JSON(404, common.APIResponse{
            Code: 404,
            Message: "Page not found",
        })
    })

    router.GET("/healthz", func(c *gin.Context) {
        c.JSON(200, common.HealthResponse{"UP"})
    })

    basket.IncludeOperations(router, store, "/basket")
    product.IncludeOperations(router, store, "/product")

    server := &Server{router: router, store: store}

    return server
}

func (s *Server) Listen(port string) {
    s.router.Run(port)
}
