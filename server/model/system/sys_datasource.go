package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type SysDataSource struct {
	global.GVA_MODEL
	Name     string `json:"name" gorm:"comment:数据源名称"`      // 数据源名称
	Type     string `json:"type" gorm:"comment:数据源类型"`      // MySQL/达梦/人大金仓/PostgreSQL/Oracle/SQLServer
	Host     string `json:"host" gorm:"comment:主机地址"`        // 主机地址
	Port     int    `json:"port" gorm:"comment:端口"`           // 端口
	Database string `json:"database" gorm:"comment:数据库名称"`  // 数据库名称
	Username string `json:"username" gorm:"comment:用户名"`      // 用户名
	Password string `json:"password" gorm:"comment:密码"`       // 密码
	Status   uint   `json:"status" gorm:"default:1;comment:状态 0禁用 1启用"` // 状态 0禁用 1启用
	Remark   string `json:"remark" gorm:"comment:备注"`         // 备注
}

func (SysDataSource) TableName() string {
	return "sys_datasource"
}
