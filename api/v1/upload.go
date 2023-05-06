package v1

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 文件上传
// @Description 文件上传
// @Tags upload
// @Accept json
// @Produce json
// @Param user_id path string true "ID"
// @Param Authorization header string true "Bearer {token}"
// @Success 200 {object} Response
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /v1/ws [post]
func Upload(c *gin.Context) {}

// UploadImage godoc
// @Summary 上传图片
// @Description 上传图片到服务器
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param file formData file true "上传的图片文件"
// @Success 200 {string} string "上传成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 401 {object} ErrorResponse "未认证授权"
// @Failure 500 {object} ErrorResponse "服务器内部错误"
// @Router /v1/upload/image [post]
func UploadImage(c *gin.Context) {
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "图片上传失败"})
		return
	}
	// 将图片保存到本地

	fileSuffix := file.Filename[strings.IndexByte(file.Filename, '.'):]
	err = c.SaveUploadedFile(file, "upload/image/"+fmt.Sprintf("%d", time.Now().Unix())+fileSuffix)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "图片保存失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "图片上传成功"})
}
