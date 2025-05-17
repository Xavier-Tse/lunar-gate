package model

type UserLogin struct {
	LunarModel
	UserID uint   `json:"userID"`
	IP     string `gorm:"size:64" json:"ip"`
	Addr   string `gorm:"size:64" json:"addr"`
	UA     string `gorm:"size:128" json:"ua"`
}
