package model

type RoleButton struct {
	LunarModel
	ButtonID uint   `json:"buttonID"`
	Button   Button `gorm:"foreignKey:ButtonID" json:"-"`
	RoleID   uint   `json:"roleID"`
	Role     Role   `gorm:"foreignKey:RoleID" json:"-"`
}
