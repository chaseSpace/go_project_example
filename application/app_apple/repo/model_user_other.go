package repo

type UserOther struct {
	Uid      int    `gorm:"column:uid" json:"uid"`
	Icon     string `gorm:"column:icon" json:"icon"`
	SignDesc string `gorm:"column:sign_desc" json:"sign_desc"`
}
