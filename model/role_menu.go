package model

type RoleMenu struct {
	LunarModel
	RoleID uint `json:"roleID"`
	Role   Role `gorm:"foreignKey:RoleID" json:"-"`
	MenuID uint `json:"menuID"`
	Menu   Menu `gorm:"foreignKey:MenuID" json:"-"`
}
