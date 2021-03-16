package db

import (
    "github.com/dvdalilue/invopop/api/common"
)

type Error struct {
    Code int
    Message string
}

func (e *Error) ToAPIResponse() (*common.APIResponse) {
    return &common.APIResponse{
        Code: e.Code,
        Message: e.Message,
    }
}
