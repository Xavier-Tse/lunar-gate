package model

type UserRole struct {
	LunarModel
	UserID uint `json:"userID"`
	User   User `gorm:"foreignKey:UserID" json:"-"`
	RoleID uint `json:"roleID"`
	Role   Role `gorm:"foreignKey:RoleID" json:"-"`
}
