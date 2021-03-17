package product_test

import (
    "testing"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "encoding/json"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"

    "github.com/dvdalilue/invopop/api/server"
    "github.com/dvdalilue/invopop/api/product"
    "github.com/dvdalilue/invopop/test/mock"
    "github.com/dvdalilue/invopop/test/utils"
)

func TestGetProductsAPI(t *testing.T) {
    utils.RandomInit()

    n := utils.RandomInt(0, 10)
    products := utils.RandomProducts(n)

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
