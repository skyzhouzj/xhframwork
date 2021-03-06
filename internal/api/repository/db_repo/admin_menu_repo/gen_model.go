package admin_menu_repo

import "time"

// AdminMenu 管理员菜单栏表
//go:generate gormgen -structs AdminMenu -input .
type AdminMenu struct {
	Id          int32     // 主键
	AdminId     int32     // 管理员ID
	MenuId      int32     // 菜单栏ID
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
}

func (*AdminMenu) TableName() string {
	return "sys_admin_menu"
}
