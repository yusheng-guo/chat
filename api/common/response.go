package common

type ResponseCode int

const (
	OK ResponseCode = iota
	BadRequest
	FileSavingErr
	InternalServerError
)

// 响应数据
type Response struct {
	Code ResponseCode // 响应码
	Msg  string       // 消息
	Data any          // 载荷
}

// NewResponse
func NewResponse(code ResponseCode, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
	}
}
