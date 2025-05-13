package model

type RoleApi struct {
	LunarModel
	ApiID  uint `json:"apiID"`
	Api    Api  `gorm:"foreignKey:ApiID" json:"-"`
	RoleID uint `json:"roleID"`
	Role   Role `gorm:"foreignKey:RoleID" json:"role"`
}
