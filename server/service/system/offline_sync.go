package system

import (
	"encoding/json"
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// OfflineSyncService 离线同步任务服务
type OfflineSyncService struct{}

var OfflineSyncServiceApp = new(OfflineSyncService)

// CreateOfflineSync 创建离线同步任务
func (s *OfflineSyncService) CreateOfflineSync(req request.CreateOfflineSyncRequest) (*system.SysOfflineSync, error) {
	// 序列化 column 字段
	columnJSON, err := json.Marshal(req.Column)
	if err != nil {
		return nil, errors.New("序列化列配置失败: " + err.Error())
	}

	// 生成 SeaTunnel 脚本
	scriptReq := request.GenerateSeatunnelScriptRequest{
		TaskName:           req.TaskName,
		SourceDataSourceId:  req.SourceDataSourceId,
		SourceDatabase:     req.SourceDatabase,
		SourceTable:        req.SourceTable,
		TargetDataSourceId: req.TargetDataSourceId,
		TargetDatabase:     req.TargetDatabase,
		TargetTable:        req.TargetTable,
		Parallelism:        req.Parallelism,
		BatchSize:          req.BatchSize,
		SplitSize:          req.SplitSize,
		UpdateMode:         req.UpdateMode,
		DataSaveMode:       req.DataSaveMode,
		SinkBatchSize:      req.SinkBatchSize,
		Column:             req.Column,
	}

	script, err := SeatunnelScriptServiceApp.GenerateScript(scriptReq)
	if err != nil {
		return nil, errors.New("生成脚本失败: " + err.Error())
	}

	scriptContent, err := json.Marshal(script)
	if err != nil {
		return nil, errors.New("序列化脚本失败: " + err.Error())
	}

	// 创建任务记录
	offlineSync := &system.SysOfflineSync{
		TaskName:           req.TaskName,
		SourceType:         req.SourceType,
		SourceDataSourceId: req.SourceDataSourceId,
		SourceDatabase:     req.SourceDatabase,
		SourceTable:        req.SourceTable,
		TargetType:         req.TargetType,
		TargetDataSourceId: req.TargetDataSourceId,
		TargetDatabase:     req.TargetDatabase,
		TargetTable:        req.TargetTable,
		Parallelism:        req.Parallelism,
		BatchSize:          req.BatchSize,
		SplitSize:          req.SplitSize,
		UpdateMode:         req.UpdateMode,
		DataSaveMode:       req.DataSaveMode,
		SinkBatchSize:      req.SinkBatchSize,
		Column:             string(columnJSON),
		ScriptContent:      string(scriptContent),
		CronExpression:     req.CronExpression,
		Status:             1,
	}

	if err := global.GVA_DB.Create(offlineSync).Error; err != nil {
		return nil, errors.New("创建任务失败: " + err.Error())
	}

	return offlineSync, nil
}

// UpdateOfflineSync 更新离线同步任务
func (s *OfflineSyncService) UpdateOfflineSync(req request.UpdateOfflineSyncRequest) error {
	// 序列化 column 字段
	columnJSON, err := json.Marshal(req.Column)
	if err != nil {
		return errors.New("序列化列配置失败: " + err.Error())
	}

	// 重新生成 SeaTunnel 脚本
	scriptReq := request.GenerateSeatunnelScriptRequest{
		TaskName:           req.TaskName,
		SourceDataSourceId:  req.SourceDataSourceId,
		SourceDatabase:     req.SourceDatabase,
		SourceTable:        req.SourceTable,
		TargetDataSourceId: req.TargetDataSourceId,
		TargetDatabase:     req.TargetDatabase,
		TargetTable:        req.TargetTable,
		Parallelism:        req.Parallelism,
		BatchSize:          req.BatchSize,
		SplitSize:          req.SplitSize,
		UpdateMode:         req.UpdateMode,
		DataSaveMode:       req.DataSaveMode,
		SinkBatchSize:      req.SinkBatchSize,
		Column:             req.Column,
	}

	script, err := SeatunnelScriptServiceApp.GenerateScript(scriptReq)
	if err != nil {
		return errors.New("生成脚本失败: " + err.Error())
	}

	scriptContent, err := json.Marshal(script)
	if err != nil {
		return errors.New("序列化脚本失败: " + err.Error())
	}

	updates := map[string]interface{}{
		"task_name":            req.TaskName,
		"source_type":          req.SourceType,
		"source_data_source_id": req.SourceDataSourceId,
		"source_database":      req.SourceDatabase,
		"source_table":         req.SourceTable,
		"target_type":          req.TargetType,
		"target_data_source_id": req.TargetDataSourceId,
		"target_database":      req.TargetDatabase,
		"target_table":         req.TargetTable,
		"parallelism":          req.Parallelism,
		"batch_size":           req.BatchSize,
		"split_size":           req.SplitSize,
		"update_mode":          req.UpdateMode,
		"data_save_mode":       req.DataSaveMode,
		"sink_batch_size":      req.SinkBatchSize,
		"column":               string(columnJSON),
		"script_content":       string(scriptContent),
		"cron_expression":      req.CronExpression,
		"status":               req.Status,
	}

	return global.GVA_DB.Model(&system.SysOfflineSync{}).Where("id = ?", req.ID).Updates(updates).Error
}

// DeleteOfflineSync 删除离线同步任务
func (s *OfflineSyncService) DeleteOfflineSync(id uint) error {
	return global.GVA_DB.Where("id = ?", id).Delete(&system.SysOfflineSync{}).Error
}

// GetOfflineSyncById 根据ID获取离线同步任务
func (s *OfflineSyncService) GetOfflineSyncById(id uint) (*system.SysOfflineSync, error) {
	var offlineSync system.SysOfflineSync
	if err := global.GVA_DB.First(&offlineSync, id).Error; err != nil {
		return nil, err
	}
	return &offlineSync, nil
}

// GetOfflineSyncList 分页获取离线同步任务列表
func (s *OfflineSyncService) GetOfflineSyncList(page, pageSize int) ([]system.SysOfflineSync, int64, error) {
	var list []system.SysOfflineSync
	var total int64

	db := global.GVA_DB.Model(&system.SysOfflineSync{})

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}
	if err := db.Order("id DESC").Offset(offset).Limit(pageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// RegenerateScript 重新生成脚本
func (s *OfflineSyncService) RegenerateScript(id uint) (string, error) {
	var offlineSync system.SysOfflineSync
	if err := global.GVA_DB.First(&offlineSync, id).Error; err != nil {
		return "", errors.New("任务不存在: " + err.Error())
	}

	// 反序列化 column
	var column []string
	if offlineSync.Column != "" {
		if err := json.Unmarshal([]byte(offlineSync.Column), &column); err != nil {
			return "", errors.New("解析列配置失败: " + err.Error())
		}
	}

	// 生成脚本
	scriptReq := request.GenerateSeatunnelScriptRequest{
		TaskName:           offlineSync.TaskName,
		SourceDataSourceId:  offlineSync.SourceDataSourceId,
		SourceDatabase:     offlineSync.SourceDatabase,
		SourceTable:        offlineSync.SourceTable,
		TargetDataSourceId: offlineSync.TargetDataSourceId,
		TargetDatabase:     offlineSync.TargetDatabase,
		TargetTable:        offlineSync.TargetTable,
		Parallelism:        offlineSync.Parallelism,
		BatchSize:          offlineSync.BatchSize,
		SplitSize:          offlineSync.SplitSize,
		UpdateMode:         offlineSync.UpdateMode,
		DataSaveMode:       offlineSync.DataSaveMode,
		SinkBatchSize:      offlineSync.SinkBatchSize,
		Column:             column,
	}

	script, err := SeatunnelScriptServiceApp.GenerateScript(scriptReq)
	if err != nil {
		return "", errors.New("生成脚本失败: " + err.Error())
	}

	scriptContent, err := json.Marshal(script)
	if err != nil {
		return "", errors.New("序列化脚本失败: " + err.Error())
	}

	// 更新数据库中的脚本内容
	if err := global.GVA_DB.Model(&system.SysOfflineSync{}).Where("id = ?", id).Update("script_content", string(scriptContent)).Error; err != nil {
		return "", errors.New("更新脚本失败: " + err.Error())
	}

	return string(scriptContent), nil
}

// GetScriptById 根据ID获取脚本内容
func (s *OfflineSyncService) GetScriptById(id uint) (string, error) {
	var offlineSync system.SysOfflineSync
	if err := global.GVA_DB.Select("script_content").First(&offlineSync, id).Error; err != nil {
		return "", err
	}
	return offlineSync.ScriptContent, nil
}