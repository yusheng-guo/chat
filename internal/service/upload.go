package service

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/yushengguo557/chat/global"
	"github.com/yushengguo557/chat/internal/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

// UploadFile 上传文件
func (s *Service) UploadFile(fileHeader *multipart.FileHeader) (*FileInfo, error) {
	filetype := upload.GetFileType(fileHeader.Filename)           // 获取文件类型
	filename := upload.RenameFileByUnix(fileHeader.Filename)      // 重命名
	uploadpath := global.Storage.SavePath                         // 上传路径
	dstdir := filepath.Join(uploadpath, filetype.String())        // 目标存储目录
	dst := filepath.Join(uploadpath, filetype.String(), filename) // 目标存储目录 + 文件名
	if _, err := os.Stat(dstdir); os.IsNotExist(err) {            // 目录不存在 创建目录
		if err = upload.CreateSavePath(dstdir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("`create save path in `UploadFile` Function: %w", err)
		}
	}
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, fmt.Errorf("save file in `UploadFile` Function: %w", err)
	}
	accessUrl := fmt.Sprintf("%sstorage/%s/%s", global.Storage.ServerUrl, filetype.String(), filename)
	return &FileInfo{Name: filename, AccessUrl: accessUrl}, nil
}
