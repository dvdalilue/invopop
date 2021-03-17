package basket_test

import (
    "io"
    "fmt"
    "testing"
    "bytes"
    "encoding/json"
    "net/http"
    "context"
    "net/http/httptest"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"

    "github.com/dvdalilue/invopop/db"
    "github.com/dvdalilue/invopop/api/basket"
    "github.com/dvdalilue/invopop/test/mock"
    "github.com/dvdalilue/invopop/test/utils"
)

func TestBasketsAPI(t *testing.T) {
    utils.RandomInit()

    tcs := []utils.TestCase{
        {
            Mock: func(store *mock.MockStore) {
                n := utils.RandomInt(1, 100)
                baskets := utils.RandomBaskets(n)

                store.EXPECT().
                    GetBaskets(gomock.Any()).
                    Times(1).
                    Return(baskets)

                bps := store.EXPECT().
                    GetBasketProducts(gomock.Any(), gomock.Any()).
                    Times(int(n))
                bps.Do(func(ctx context.Context, id int64) {
                    k := utils.RandomInt(5, 10)
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
        {
            Mock: func(store *mock.MockStore) {
                basket := utils.RandomBasket()

                store.EXPECT().
                    CreateBasket(gomock.Any()).
                    Times(1).
                    Return(&basket)

                store.EXPECT().
                    GetBasketProducts(gomock.Any(), gomock.Any()).
                    Times(1).
                    Return(utils.RandomProducts(0), nil)
            },
            Assert: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                assert.Equal(t, http.StatusOK, recorder.Code)
            },
            Method: http.MethodPost,
        },
    }

    for _, tc := range tcs {
        utils.Tester(t, &tc, "/basket/", http.MethodGet)
    }
}

func TestBasketProductAPI(t *testing.T) {
    utils.RandomInit()

    b := utils.RandomBasket()
    p := utils.RandomProduct()

    tcs := []utils.TestCase{
        {
            Mock: func(store *mock.MockStore) {

                store.EXPECT().
                    AddBasketProduct(gomock.Any(), gomock.Any(), gomock.Any()).
                    Times(1).
                    Return(&b, nil)

                store.EXPECT().
                    GetBasketProducts(gomock.Any(), gomock.Any()).
                    Times(1).
                    Return([]*db.Product{&p}, nil)
            },
            Assert: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                assert.Equal(t, http.StatusOK, recorder.Code)
            },
            Body: func() io.Reader {
                payload := new(bytes.Buffer)
                body := &basket.AddBasketProduct{ProductID: p.ID}

                json.NewEncoder(payload).Encode(body)

                return payload
            }(),
        },
    }

    url := fmt.Sprintf("/basket/%d/product", b.ID)

    for _, tc := range tcs {
        utils.Tester(t, &tc, url, http.MethodPost)
    }
}

func TestBasketAPI(t *testing.T) {
    utils.RandomInit()

    tcs := []utils.TestCase{
        {
            Mock: func(store *mock.MockStore) {
                b := utils.RandomBasket()

                store.EXPECT().
                    GetBasket(gomock.Any(), gomock.Any()).
                    Times(1).
                    Return(&b, nil)

                bps := store.EXPECT().
                    GetBasketProducts(gomock.Any(), gomock.Any()).
                    Times(1)
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
                b := utils.RandomBasket()

                store.EXPECT().
                    GetBasket(gomock.Any(), gomock.Any()).
                    Times(1).
                    Return(&b, nil)

                store.EXPECT().
                    GetBasketProducts(gomock.Any(), gomock.Any()).
                    Return(nil, &db.Error{Code: http.StatusNotFound, Message: ""})
            },
            Assert: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                assert.Equal(t, http.StatusNotFound, recorder.Code)
            },
        },
        {
            Mock: func(store *mock.MockStore) {
                store.EXPECT().
                    GetBasket(gomock.Any(), gomock.Any()).
                    Times(1).
                    Return(nil, &db.Error{Code: http.StatusNotFound, Message: ""})
            },
            Assert: func(t *testing.T, recorder *httptest.ResponseRecorder) {
                assert.Equal(t, http.StatusNotFound, recorder.Code)
            },
        },
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
            Method: http.MethodDelete,
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
            Method: http.MethodDelete,
        },
    }

    for _, tc := range tcs {
        utils.Tester(t, &tc, "/basket/1", http.MethodGet)
    }
}
