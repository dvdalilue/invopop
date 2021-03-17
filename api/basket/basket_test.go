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

    // "github.com/dvdalilue/invopop/db"
    "github.com/dvdalilue/invopop/api/server"
    // "github.com/dvdalilue/invopop/api/basket"
    "github.com/dvdalilue/invopop/test/mock"
    "github.com/dvdalilue/invopop/test/utils"
    // "github.com/dvdalilue/invopop/api/product"
)

func TestBasketsAPI(t *testing.T) {
    utils.RandomInit()

    n := utils.RandomInt(1, 10)
    baskets := utils.RandomBaskets(n)

    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    store := mock.NewMockStore(ctrl)
    store.EXPECT().GetBaskets(gomock.Any()).Times(1).Return(baskets)

    bps := store.EXPECT().
        GetBasketProducts(gomock.Any(), gomock.Any()).
        Times(int(n))
    bps.Do(func(ctx context.Context, id int64) {
        k := utils.RandomInt(0, 10)
        bps.Return(utils.RandomProducts(k), nil)
    })

    server := server.NewServer(store)
    recorder := httptest.NewRecorder()

    url := "/basket/"
    request, err := http.NewRequest(http.MethodGet, url, nil)
    assert.NoError(t, err)

    server.ServeHTTP(recorder, request)
    assert.Equal(t, http.StatusOK, recorder.Code)
}
