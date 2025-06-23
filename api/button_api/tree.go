package button_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
	"sort"
)

type GroupListResponse struct {
	GroupTitle string         `json:"groupTitle"`
	List       []model.Button `json:"list"`
}

func (ButtonApi) Tree(c *gin.Context) {
	var data = map[string][]model.Button{}
	var _list []model.Button
	global.DB.Find(&_list)
	var keys []string
	for _, u := range _list {
		_, ok := data[u.Group]
		if !ok {
			keys = append(keys, u.Group)
		}
		data[u.Group] = append(data[u.Group], u)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	var list = make([]GroupListResponse, 0)
	for _, key := range keys {
		list = append(list, GroupListResponse{
			GroupTitle: key,
			List:       data[key],
		})
	}
	res.OkWithData(list, c)
}
