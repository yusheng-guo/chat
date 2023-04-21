package v1

// 错误响应
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
