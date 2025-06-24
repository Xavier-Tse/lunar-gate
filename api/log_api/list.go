package log_api

import (
	"fmt"
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/Xavier-Tse/lunar-gate/utils/file"
	"github.com/gin-gonic/gin"
	"time"
)

func (LogApi) List(c *gin.Context) {
	var cr model.Page
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailBinding(err, c)
		return
	}
	fileName := fmt.Sprintf("logs/%s/info.log", time.Now().Format("2006-01-02"))
	lines, err := file.ReverseRead(fileName, cr.Limit)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	res.OkWithData(lines, c)
}
