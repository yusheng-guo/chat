package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/chat/api/common"
	"github.com/yushengguo557/chat/internal/service"
)

// UploadImage godoc
// @Summary 上传文件
// @Description 上传文件
// @Accept multipart/form-data
// @Produce json
// @Security ApiKeyAuth
// @Param file formData file true "上传的文件"
// @Success 200 {object} UploadResponse "上传成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 401 {object} ErrorResponse "未认证授权"
// @Failure 500 {object} ErrorResponse "服务器内部错误"
// @Router /upload [post]
func Upload(c *gin.Context) {
	// 1.获取文件
	fileHeader, err := c.FormFile("file")
	if err != nil {
		rsp := common.NewResponse(common.BadRequest, "文件上传失败")
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	// 2.将文件保存到服务器
	svc := service.NewService(c.Request.Context())
	info, err := svc.UploadFile(fileHeader)
	if err != nil {
		rsp := common.NewResponse(common.FileSavingErr, "文件保存失败")
		c.JSON(http.StatusInternalServerError, rsp)
		return
	}
	// 3.响应 可访问地址
	rsp := common.NewResponse(common.OK, info.AccessUrl)
	c.JSON(http.StatusOK, rsp)
}

// func UploadImage(c *gin.Context) {
// 	// 获取文件
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"msg": "图片上传失败"})
// 		return
// 	}
// 	// 将图片保存到服务器
// 	fileSuffix := file.Filename[strings.IndexByte(file.Filename, '.'):]
// 	filename := fmt.Sprintf("%d", time.Now().Unix()) + fileSuffix
// 	filepath := "storage/image/" + filename
// 	err = c.SaveUploadedFile(file, filepath)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"msg": "图片保存失败"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"msg": c.Request.Host + "/image/" + filename})
// }
