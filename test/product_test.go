package test

import (
    "io/ioutil"
    "testing"
    "encoding/json"
    "net/http"
    "net/http/httptest"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"

    "github.com/dvdalilue/invopop/db"
    "github.com/dvdalilue/invopop/api/server"
    "github.com/dvdalilue/invopop/api/product"
    "github.com/dvdalilue/invopop/test/mock"
)

func randomProduct() db.Product {
    return db.Product{
        ID:    RandomInt(1, 100),
        Name:  RandomString(10),
        Price: RandomFloat(2.0, 42.0),
    }
}

func randomProducts(n int64) []*db.Product {
    products := make([]*db.Product, n)

    var p db.Product

    for i := int64(0); i < n; i++ {
        p = randomProduct()
        products[i] = &p
    }

    return products
}

func TestGetProductsAPI(t *testing.T) {
    RandomInit()

    n := RandomInt(0, 10)
    products := randomProducts(n)

    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    store := mock.NewMockStore(ctrl)
    store.EXPECT().GetProducts(gomock.Any()).Times(1).Return(products)

    server := server.NewServer(store)
    recorder := httptest.NewRecorder()

    url := "/product/"
    request, err := http.NewRequest(http.MethodGet, url, nil)
    assert.NoError(t, err)

    server.ServeHTTP(recorder, request)
    assert.Equal(t, http.StatusOK, recorder.Code)

    var res product.Products = *product.ToProductsDto(products)

    data, err := ioutil.ReadAll(recorder.Body)
    assert.NoError(t, err)

    var gotProducts product.Products
    err = json.Unmarshal(data, &gotProducts)
    assert.NoError(t, err)

    assert.Equal(t, gotProducts, res)
}
