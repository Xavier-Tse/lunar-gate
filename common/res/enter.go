package res

import (
	"github.com/gin-gonic/gin"
	"github.com/Xavier-Tse/lunar-gate/utils/validate"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func response(code int, data any, message string, c *gin.Context) {
	r := Response{
		Code:    code,
		Data:    data,
		Message: message,
	}
	c.JSON(http.StatusOK, r)
}

func Ok(data any, message string, c *gin.Context) {
	response(CodeOk, data, message, c)
}

func OkWithData(data any, c *gin.Context) {
	response(CodeOk, data, "成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	response(CodeOk, gin.H{}, message, c)
}

func OkWithList(list any, count int64, c *gin.Context) {
	response(CodeOk, gin.H{
		"list":  list,
		"count": count,
	}, "成功", c)
}

func Fail(code int, message string, c *gin.Context) {
	response(code, gin.H{}, message, c)
}

func FailWithMessage(message string, c *gin.Context) {
	response(CodeError, gin.H{}, message, c)
}

func FailWithError(err error, c *gin.Context) {
	response(CodeError, gin.H{}, err.Error(), c)
}

func FailBinding(err error, c *gin.Context) {
	data := validate.Error(err)
	response(CodeError, data.FieldMap, data.Message, c)
}
