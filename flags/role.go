package flags

import (
	"github.com/Xavier-Tse/lunar-gate/global"
	"github.com/Xavier-Tse/lunar-gate/model"
	"github.com/sirupsen/logrus"
)

type role struct{}

func (role) Init() {
	var clearTableList = []string{
		"apis",
		"casbin_rule",
		"menus",
		"role_apis",
		"role_menus",
		"roles",
		"user_roles",
	}

	for _, tableName := range clearTableList {
		global.DB.Exec("truncate " + tableName)
		logrus.Infof("清空表 %s", tableName)
	}

	//3. 创建api
	var apiList = []model.Api{
		{Group: "用户管理", Name: "用户列表", Path: "/api/user/list", Method: "GET"},
		{Group: "用户管理", Name: "删除用户", Path: "/api/user/remove", Method: "DELETE"},
		{Group: "用户管理", Name: "用户修改个人信息", Path: "/api/user/info", Method: "PUT"},
		{Group: "用户管理", Name: "用户详情", Path: "/api/user/info", Method: "GET"},
		{Group: "用户管理", Name: "用户改密码", Path: "/api/user/password", Method: "PUT"},
		{Group: "用户管理", Name: "给用户分配角色", Path: "/api/user/role", Method: "PUT"},
		{Group: "用户管理", Name: "用户登录", Path: "/api/user/login", Method: "POST"},
		{Group: "用户管理", Name: "用户注册", Path: "/api/user/register", Method: "POST"},
	}

	global.DB.Create(&apiList)
	logrus.Infof("创建api记录 %d 条", len(apiList))

	//4. 创建菜单
	var menuList = []model.Menu{
		{
			LunarModel: model.LunarModel{
				ID: 1,
			},
			Enable:    true,
			Name:      "data",
			Path:      "/admin/data",
			Component: "@/view/admin/data.vue",
			Meta: model.Meta{
				Title: "数据统计",
				Icon:  "fa fa-data",
			},
			Sort: 10,
		},
		{
			LunarModel: model.LunarModel{
				ID: 2,
			},
			Enable:    true,
			Name:      "permission",
			Path:      "/admin/permission",
			Component: "@/view/admin/permission.vue",
			Meta: model.Meta{
				Title: "权限管理",
				Icon:  "fa fa-data",
			},
			Sort: 9,
		},
		{
			LunarModel: model.LunarModel{
				ID: 3,
			},
			Enable:    true,
			Name:      "permission_user",
			Path:      "/admin/permission/user",
			Component: "@/view/admin/permission/user.vue",
			Meta: model.Meta{
				Title: "用户列表",
				Icon:  "fa fa-data",
			},
			Sort:         5,
			ParentMenuID: uint2Ptr(2),
		},
		{
			LunarModel: model.LunarModel{
				ID: 4,
			},
			Enable:    true,
			Name:      "permission_role",
			Path:      "/admin/permission/role",
			Component: "@/view/admin/permission/role.vue",
			Meta: model.Meta{
				Title: "角色列表",
				Icon:  "fa fa-data",
			},
			Sort:         5,
			ParentMenuID: uint2Ptr(2),
		},
	}
	global.DB.Create(&menuList)
	logrus.Infof("创建菜单 %d 条", len(menuList))
}

func uint2Ptr(u uint) *uint {
	return &u
}
