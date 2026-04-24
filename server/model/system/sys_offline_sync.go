package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysOfflineSync 离线同步任务配置
type SysOfflineSync struct {
	global.GVA_MODEL
	TaskName          string `json:"taskName" gorm:"comment:任务名称"`                    // 任务名称
	SourceType        string `json:"sourceType" gorm:"comment:源类型"`                  // MySQL/PostgreSQL/SQLServer等
	SourceDataSourceId uint   `json:"sourceDataSourceId" gorm:"comment:源数据源ID"`        // 源数据源ID
	SourceDatabase    string `json:"sourceDatabase" gorm:"comment:源数据库"`             // 源数据库
	SourceTable       string `json:"sourceTable" gorm:"comment:源表"`                    // 源表
	TargetType        string `json:"targetType" gorm:"comment:目标类型"`                  // MySQL/PostgreSQL/SQLServer等
	TargetDataSourceId uint   `json:"targetDataSourceId" gorm:"comment:目标数据源ID"`        // 目标数据源ID
	TargetDatabase    string `json:"targetDatabase" gorm:"comment:目标数据库"`             // 目标数据库
	TargetTable       string `json:"targetTable" gorm:"comment:目标表"`                  // 目标表
	Parallelism       int    `json:"parallelism" gorm:"default:4;comment:并行度"`          // 并行度
	BatchSize         int    `json:"batchSize" gorm:"default:5000;comment:批处理大小"`     // 批处理大小
	SplitSize         int    `json:"splitSize" gorm:"default:8096;comment:分割大小"`       // 分割大小
	UpdateMode        string `json:"updateMode" gorm:"default:INSERT;comment:更新模式"`    // 更新模式
	DataSaveMode      string `json:"dataSaveMode" gorm:"default:OVERWRITE;comment:数据保存模式"` // 数据保存模式
	SinkBatchSize     int    `json:"sinkBatchSize" gorm:"default:1024;comment:目标批处理大小"` // 目标批处理大小
	Column            string `json:"column" gorm:"type:text;comment:同步的列，JSON数组"`      // 同步的列（JSON数组格式存储）
	ScriptContent     string `json:"scriptContent" gorm:"type:text;comment:生成的SeaTunnel脚本JSON"` // 生成的脚本内容
	Status            uint   `json:"status" gorm:"default:1;comment:状态 0禁用 1启用"`       // 状态
	CronExpression    string `json:"cronExpression" gorm:"comment:Cron表达式"`             // Cron表达式
	LastSyncTime      *uint  `json:"lastSyncTime" gorm:"comment:上次同步时间戳"`            // 上次同步时间
	NextSyncTime      *uint  `json:"nextSyncTime" gorm:"comment:下次同步时间戳"`            // 下次同步时间
	LastSyncStatus    string `json:"lastSyncStatus" gorm:"default:pending;comment:上次同步状态 pending/running/success/failed"` // 上次同步状态
	LastSyncMsg       string `json:"lastSyncMsg" gorm:"type:text;comment:上次同步消息/错误信息"` // 上次同步消息
}

func (SysOfflineSync) TableName() string {
	return "sys_offline_sync"
}