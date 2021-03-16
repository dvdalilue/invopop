package common

type APIResponse struct {

    // code
    Code int `json:"code,omitempty"`

    // message
    Message string `json:"message,omitempty"`
}

type HealthResponse struct {

    // message
    Status string `json:"status,omitempty"`
}

type Object struct {

	// id
    ID int64 `uri:"id"`
}