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
	Image    FileType = iota + 1 // 图片
	Video                        // 视频
	Audio                        // 音频
	Document                     // 文档
	Text                         // 文本
	Other                        // 其他
)

// String 将整形转为 字符串
func (t FileType) String() string {
	switch t {
	case Image:
		return "image"
	case Video:
		return "video"
	case Audio:
		return "audio"
	case Document:
		return "document"
	case Text:
		return "text"
	default:
		return "other"
	}
}

// func 获取文件类型
func GetFileType(name string) FileType {
	ext := path.Ext(name)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return Image
	case ".mp4", ".avi", ".mov", ".wmv":
		return Video
	case ".mp3", ".wav", ".ogg", ".flac":
		return Audio
	case ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".md":
		return Document
	case ".txt", ".go", ".c", ".rs", ".js":
		return Text
	default:
		return Other
	}
}

// RenameFileByUnix 通过 时间戳 重命名文件
func RenameFileByUnix(name string) string {
	ext := path.Ext(name)
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
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
