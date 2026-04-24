package system

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// SeaTunnelScript SeaTunnel 任务脚本结构
type SeaTunnelScript struct {
	Env    SeaTunnelEnv              `json:"env"`
	Source map[string]interface{}    `json:"source"`
	Sink   map[string]interface{}    `json:"sink"`
}

// SeaTunnelEnv 环境配置
type SeaTunnelEnv struct {
	Job        SeaTunnelJob `json:"job"`
	Parallelism int          `json:"parallelism"`
	BatchSize  int          `json:"batchSize"`
}

// SeaTunnelJob 作业配置
type SeaTunnelJob struct {
	Name string `json:"name"`
	Mode string `json:"mode"`
}

// SeatunnelScriptService SeaTunnel 脚本服务
type SeatunnelScriptService struct{}

var SeatunnelScriptServiceApp = new(SeatunnelScriptService)

// GenerateScript 生成 SeaTunnel 任务脚本
func (s *SeatunnelScriptService) GenerateScript(req request.GenerateSeatunnelScriptRequest) (*SeaTunnelScript, error) {
	// 1. 获取源数据源配置
	var sourceDS system.SysDataSource
	if err := global.GVA_DB.First(&sourceDS, req.SourceDataSourceId).Error; err != nil {
		return nil, fmt.Errorf("源数据源不存在: %w", err)
	}

	// 2. 获取目标数据源配置
	var targetDS system.SysDataSource
	if err := global.GVA_DB.First(&targetDS, req.TargetDataSourceId).Error; err != nil {
		return nil, fmt.Errorf("目标数据源不存在: %w", err)
	}

	// 3. 构建 SeaTunnel 脚本
	script := &SeaTunnelScript{
		Env: SeaTunnelEnv{
			Job: SeaTunnelJob{
				Name: req.TaskName,
				Mode: "BATCH",
			},
			Parallelism: req.GetParallelism(),
			BatchSize:   req.GetBatchSize(),
		},
		Source: s.buildSource(sourceDS, req),
		Sink:   s.buildSink(targetDS, req),
	}

	return script, nil
}

// buildSource 构建源配置
func (s *SeatunnelScriptService) buildSource(ds system.SysDataSource, req request.GenerateSeatunnelScriptRequest) map[string]interface{} {
	source := map[string]interface{}{
		"baseUrl":   s.buildBaseUrl(ds),
		"username":  ds.Username,
		"password":  ds.Password,
		"database":  req.SourceDatabase,
		"table":     req.SourceTable,
		"column":    req.GetColumn(),
		"splitSize": req.GetSplitSize(),
	}
	return map[string]interface{}{
		ds.Type: source,
	}
}

// buildSink 构建目标配置
func (s *SeatunnelScriptService) buildSink(ds system.SysDataSource, req request.GenerateSeatunnelScriptRequest) map[string]interface{} {
	sink := map[string]interface{}{
		"baseUrl":      s.buildBaseUrl(ds),
		"username":     ds.Username,
		"password":     ds.Password,
		"database":     req.TargetDatabase,
		"table":        req.TargetTable,
		"updateMode":   req.GetUpdateMode(),
		"dataSaveMode": req.GetDataSaveMode(),
		"batchSize":    req.GetSinkBatchSize(),
	}
	return map[string]interface{}{
		ds.Type: sink,
	}
}

// buildBaseUrl 构建 JDBC URL
func (s *SeatunnelScriptService) buildBaseUrl(ds system.SysDataSource) string {
	switch ds.Type {
	case "MySQL":
		return fmt.Sprintf("jdbc:mysql://%s:%d", ds.Host, ds.Port)
	case "PostgreSQL":
		return fmt.Sprintf("jdbc:postgresql://%s:%d", ds.Host, ds.Port)
	case "SQLServer":
		return fmt.Sprintf("jdbc:sqlserver://%s:%d", ds.Host, ds.Port)
	default:
		return fmt.Sprintf("jdbc:%s://%s:%d", ds.Type, ds.Host, ds.Port)
	}
}