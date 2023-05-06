package common

type Response map[string]any

// UploadResponse 上传响应
// Code
// =0 表示上传成功 msg携带 url
// ≠0 表示上传失败 msg携带错误消息
type UploadResponse struct {
	Code int    // 响应码
	Msg  string // 消息
}
