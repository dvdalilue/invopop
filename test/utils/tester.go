package utils

import (
    "testing"
    "net/http"
    "net/http/httptest"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"

    "github.com/dvdalilue/invopop/api/server"
    "github.com/dvdalilue/invopop/test/mock"
)

type TestCase struct {
    Mock func(store *mock.MockStore)
    Assert func(t *testing.T, recoder *httptest.ResponseRecorder)
}

func Tester(t *testing.T, tc *TestCase, url string, method string) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    store := mock.NewMockStore(ctrl)
    tc.Mock(store)

    server := server.NewServer(store)
    recorder := httptest.NewRecorder()

    request, err := http.NewRequest(method, url, nil)
    assert.NoError(t, err)

    server.ServeHTTP(recorder, request)
    tc.Assert(t, recorder)
}