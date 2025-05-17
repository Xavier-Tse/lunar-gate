package data_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
	"time"
)

type UserTimeAggregationResponse struct {
	DateList          []string `json:"dateList"`
	LoginCountList    []int    `json:"loginCountList"`
	RegisterCountList []int    `json:"registerCountList"`
}

func (DataApi) UserTimeAggregation(c *gin.Context) {
	var data UserTimeAggregationResponse

	now := time.Now().Format("2006-01-02") + " 23:59:59"
	before7Day := time.Now().AddDate(0, 0, -6).Format("2006-01-02") + " 00:00:00"

	var userLoginList []model.UserLogin
	global.DB.Find(&userLoginList, "created_at > ? and created_at < ?", before7Day, now)
	var userLoginMap = map[string]int{}
	for _, u := range userLoginList {
		key := u.CreatedAt.Format("2006-01-02")
		userLoginMap[key] = userLoginMap[key] + 1
	}

	var userRegisterList []model.User
	global.DB.Find(&userRegisterList, "created_at > ? and created_at < ?", before7Day, now)
	var userRegisterMap = map[string]int{}
	for _, u := range userRegisterList {
		key := u.CreatedAt.Format("2006-01-02")
		userRegisterMap[key] = userRegisterMap[key] + 1
	}

	before7Date := time.Now().AddDate(0, 0, -6)
	for i := 0; i < 7; i++ {
		date := before7Date.AddDate(0, 0, i)
		key := date.Format("2006-01-02")
		loginCount := userLoginMap[key]
		registerCount := userRegisterMap[key]
		data.DateList = append(data.DateList, key)
		data.LoginCountList = append(data.LoginCountList, loginCount)
		data.RegisterCountList = append(data.RegisterCountList, registerCount)
	}
	res.OkWithData(data, c)
}
