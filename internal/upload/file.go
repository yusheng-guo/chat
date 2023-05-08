package upload

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"time"
)

type FileType int

const (
	TypeImage    FileType = iota + 1 // 图片
	TypeMarkDown                     // markdown
)

// GetFileType 获取文件类型
func (t FileType) String() string {
	switch t {
	case TypeImage:
		return "image"
	case TypeMarkDown:
		return "markdown"
	}
	return ""
}

// RenameFileByUnix 通过 时间戳 重命名文件
func RenameFileByUnix(name string) string {
	fileSuffix := path.Ext(name)
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), fileSuffix)
}

// RenameFileByMD5 通过 md5 加密算法重命名文件
func RenameFileByMD5(name string) string {
	m := md5.New()
	m.Write([]byte(name))
	return hex.EncodeToString(m.Sum(nil))
}

// CreateSavePath 创建保存路径
func CreateSavePath(path string, perm os.FileMode) (err error) {
	err = os.MkdirAll(path, perm)
	if err != nil {
		return fmt.Errorf("create path in `CreateSavePath` Function: %w", err)
	}
	return nil
}

// SaveFile 以指定的文件名 保存文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("open file in `SaveFile` Function: %w", err)
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create file in `SaveFile` Function: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return fmt.Errorf("copy file in `SaveFile` Function: %w", err)
	}
	return nil
}
