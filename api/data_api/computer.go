package data_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/utils/computer"
	"github.com/gin-gonic/gin"
)

type ComputerResponse struct {
	Cpu  float64 `json:"cpu"`
	Mem  float64 `json:"mem"`
	Disk float64 `json:"disk"`
}

func (DataApi) Computer(c *gin.Context) {
	data := ComputerResponse{
		Cpu:  computer.GetCpuPercent(),
		Mem:  computer.GetMemPercent(),
		Disk: computer.GetDiskPercent(),
	}
	res.OkWithData(data, c)
}
