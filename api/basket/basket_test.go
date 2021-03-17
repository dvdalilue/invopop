package basket_test

import (
    // "io/ioutil"
    "testing"
    // "encoding/json"
    "net/http"
    "context"
    "net/http/httptest"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"

    "github.com/dvdalilue/invopop/db"
    "github.com/dvdalilue/invopop/test/mock"
    "github.com/dvdalilue/invopop/test/utils"
    // "github.com/dvdalilue/invopop/api/product"
)

func TestBasketsAPI(t *testing.T) {
    utils.RandomInit()

    tcs := []utils.TestCase{
        {
            Mock: func(store *mock.MockStore) {
                n := utils.RandomInt(1, 10)
                baskets := utils.RandomBaskets(n)

                store.EXPECT().
                    GetBaskets(gomock.Any()).
                    Times(1).
                    Return(baskets)

                bps := store.EXPECT().
                    GetBasketProducts(gomock.Any(), gomock.Any()).
                    Times(int(n))
                bps.Do(func(ctx context.Context, id int64) {
                    k := utils.RandomInt(0, 10)
                    bps.Return(utils.RandomProducts(k), nil)
                })
            },
            Assert: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                assert.Equal(t, http.StatusOK, recorder.Code)
            },
        },
        {
            Mock: func(store *mock.MockStore) {
                n := utils.RandomInt(1, 10)
                baskets := utils.RandomBaskets(n)

                store.EXPECT().
                    GetBaskets(gomock.Any()).
                    Times(1).
                    Return(baskets)

                store.EXPECT().
                    GetBasketProducts(gomock.Any(), gomock.Any()).
                    Return(nil, &db.Error{Code: http.StatusNotFound, Message: ""})
            },
            Assert: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                assert.Equal(t, http.StatusNotFound, recorder.Code)
            },
        },
    }

    for _, tc := range tcs {
        utils.Tester(t, &tc, "/basket/", http.MethodGet)
    }
}

func TestDeleteBasketAPI(t *testing.T) {
    utils.RandomInit()

    tcs := []utils.TestCase{
        {
            Mock: func(store *mock.MockStore) {
                store.EXPECT().
                    DeleteBasket(gomock.Any(), gomock.Any()).
                    Times(1).
                    Return(nil)
            },
            Assert: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                assert.Equal(t, http.StatusNoContent, recorder.Code)
            },
        },
        {
            Mock: func(store *mock.MockStore) {
                store.EXPECT().
                    DeleteBasket(gomock.Any(), gomock.Any()).
                    Times(1).
                    Return(&db.Error{Code: http.StatusNotFound, Message: ""})
            },
            Assert: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                assert.Equal(t, http.StatusNotFound, recorder.Code)
            },
        },
    }

    for _, tc := range tcs {
        utils.Tester(t, &tc, "/basket/1", http.MethodDelete)
    }
}
