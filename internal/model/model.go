package model

import "time"

// Model 公共模型
type Model struct {
	ID        string     `json:"id" gorethink:"id,omitempty"`                 // 模型ID
	CreatedAt *time.Time `json:"created_at" gorethink:"created_at,omitempty"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at" gorethink:"updated_at,omitempty"` // 更新时间
	DeletedAt *time.Time `json:"deleted_at" gorethink:"deleted_at,omitempty"` // 删除时间
	IsDel     bool       `json:"is_del" gorethink:"is_del,omitempty"`         // 是否删除
}
