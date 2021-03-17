package product_test

import (
    "testing"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "encoding/json"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"

    "github.com/dvdalilue/invopop/api/product"
    "github.com/dvdalilue/invopop/test/mock"
    "github.com/dvdalilue/invopop/test/utils"
)

func TestGetProductsAPI(t *testing.T) {
    utils.RandomInit()

    n := utils.RandomInt(0, 10)
    products := utils.RandomProducts(n)

    tcs := []utils.TestCase{
        {
            Mock: func(store *mock.MockStore) {
                store.EXPECT().
                    GetProducts(gomock.Any()).
                    Times(1).
                    Return(products)
            },
            Assert: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                assert.Equal(t, http.StatusOK, recorder.Code)

                res := *product.ToProductsDto(products)

                data, err := ioutil.ReadAll(recorder.Body)
                assert.NoError(t, err)

                var responseProducts product.Products
                err = json.Unmarshal(data, &responseProducts)
                assert.NoError(t, err)

                assert.Equal(t, responseProducts, res)
            },
        },
    }

    for _, tc := range tcs {
        utils.Tester(t, &tc, "/product/", http.MethodGet)
    }
}
