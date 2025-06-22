package image_api

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/middleware"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
	"time"
)

var whiteMap = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
}

func (ImageApi) Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.FailWithMessage("请上传文件", c)
		return
	}

	if fileHeader.Size > 1024*1024*global.Config.Info.File.Size {
		res.FailWithMessage("文件过大", c)
		return
	}

	_, ok := whiteMap[strings.ToLower(path.Ext(fileHeader.Filename))]
	if !ok {
		res.FailWithMessage("文件格式错误", c)
		return
	}

	auth := middleware.GetAuth(c)
	dist := fmt.Sprintf("static/%s/%s/%d/%s", global.Config.Info.File.Dir, auth.Username, time.Now().UnixNano(), fileHeader.Filename)
	if err := c.SaveUploadedFile(fileHeader, dist); err != nil {
		res.FailWithMessage("上传失败", c)
		return
	}
	res.Ok("/"+dist, "图片上传成功", c)
}
