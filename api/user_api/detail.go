package user_api

import (
	"github.com/Xavier-Tse/lunar-gate/common/res"
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/middleware"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/gin-gonic/gin"
	"time"
)

type UserDetailResponse struct {
	UserID    uint      `json:"userID"`
	Nickname  string    `json:"nickname"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	RoleList  []string  `json:"roleList"`
	Addr      string    `json:"addr"`
}

func (UserApi) UserDetail(c *gin.Context) {
	claims := middleware.GetAuth(c)

	var user model.User
	err := global.DB.Preload("RoleList").Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}

	var userLoginModel model.UserLogin
	global.DB.Order("created_at desc").Take(&userLoginModel, "user_id = ?", user.ID)

	var roleList []string
	for _, u := range user.RoleList {
		roleList = append(roleList, u.Title)
	}

	data := UserDetailResponse{
		UserID:    user.ID,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		RoleList:  roleList,
		Addr:      userLoginModel.Addr,
	}
	res.OkWithData(data, c)
}
