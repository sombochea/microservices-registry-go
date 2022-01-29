package model

type ApiErrorResponse struct {
	StatusCode        int    `json:"status_code"`
	StatusDescription string `json:"status_description"`
}

type ApiResponse struct {
	StatusCode        int         `json:"status_code"`
	StatusDescription string      `json:"status_description"`
	Data              interface{} `json:"data"`
}