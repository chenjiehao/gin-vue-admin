package request

// CreateOfflineSyncRequest 创建离线同步任务请求
type CreateOfflineSyncRequest struct {
	TaskName          string `json:"taskName" binding:"required"`                    // 任务名称
	SourceType        string `json:"sourceType" binding:"required"`                 // 源类型
	SourceDataSourceId uint   `json:"sourceDataSourceId" binding:"required"`          // 源数据源ID
	SourceDatabase    string `json:"sourceDatabase" binding:"required"`             // 源数据库
	SourceTable       string `json:"sourceTable" binding:"required"`                 // 源表
	TargetType        string `json:"targetType" binding:"required"`                 // 目标类型
	TargetDataSourceId uint   `json:"targetDataSourceId" binding:"required"`          // 目标数据源ID
	TargetDatabase    string `json:"targetDatabase" binding:"required"`             // 目标数据库
	TargetTable       string `json:"targetTable" binding:"required"`                 // 目标表
	Parallelism       int    `json:"parallelism"`                                    // 并行度
	BatchSize         int    `json:"batchSize"`                                      // 批处理大小
	SplitSize         int    `json:"splitSize"`                                      // 分割大小
	UpdateMode        string `json:"updateMode"`                                     // 更新模式
	DataSaveMode      string `json:"dataSaveMode"`                                   // 数据保存模式
	SinkBatchSize     int    `json:"sinkBatchSize"`                                  // 目标批处理大小
	Column            []string `json:"column"`                                       // 同步的列
	CronExpression    string `json:"cronExpression"`                                // Cron表达式
}

// UpdateOfflineSyncRequest 更新离线同步任务请求
type UpdateOfflineSyncRequest struct {
	ID                uint   `json:"ID" binding:"required"`                         // 任务ID
	TaskName          string `json:"taskName" binding:"required"`                   // 任务名称
	SourceType        string `json:"sourceType" binding:"required"`                  // 源类型
	SourceDataSourceId uint   `json:"sourceDataSourceId" binding:"required"`         // 源数据源ID
	SourceDatabase    string `json:"sourceDatabase" binding:"required"`              // 源数据库
	SourceTable       string `json:"sourceTable" binding:"required"`                // 源表
	TargetType        string `json:"targetType" binding:"required"`                 // 目标类型
	TargetDataSourceId uint   `json:"targetDataSourceId" binding:"required"`         // 目标数据源ID
	TargetDatabase    string `json:"targetDatabase" binding:"required"`            // 目标数据库
	TargetTable       string `json:"targetTable" binding:"required"`               // 目标表
	Parallelism       int    `json:"parallelism"`                                   // 并行度
	BatchSize         int    `json:"batchSize"`                                     // 批处理大小
	SplitSize         int    `json:"splitSize"`                                     // 分割大小
	UpdateMode        string `json:"updateMode"`                                    // 更新模式
	DataSaveMode      string `json:"dataSaveMode"`                                  // 数据保存模式
	SinkBatchSize     int    `json:"sinkBatchSize"`                                 // 目标批处理大小
	Column            []string `json:"column"`                                      // 同步的列
	CronExpression    string `json:"cronExpression"`                               // Cron表达式
	Status            uint   `json:"status"`                                       // 状态
}

// GenerateSeatunnelScriptAndSaveRequest 生成脚本并保存请求
type GenerateSeatunnelScriptAndSaveRequest struct {
	TaskName           string   `json:"taskName" binding:"required"`                    // 任务名称
	SourceDataSourceId uint     `json:"sourceDataSourceId" binding:"required"`         // 源数据源ID
	SourceDatabase     string   `json:"sourceDatabase" binding:"required"`             // 源数据库
	SourceTable        string   `json:"sourceTable" binding:"required"`                // 源表
	TargetDataSourceId uint     `json:"targetDataSourceId" binding:"required"`         // 目标数据源ID
	TargetDatabase     string   `json:"targetDatabase" binding:"required"`            // 目标数据库
	TargetTable        string   `json:"targetTable" binding:"required"`                // 目标表
	Parallelism        int      `json:"parallelism"`                                   // 并行度
	BatchSize          int      `json:"batchSize"`                                     // 批处理大小
	SplitSize          int      `json:"splitSize"`                                     // 分割大小
	UpdateMode         string   `json:"updateMode"`                                    // 更新模式
	DataSaveMode       string   `json:"dataSaveMode"`                                  // 数据保存模式
	SinkBatchSize      int      `json:"sinkBatchSize"`                                 // 目标批处理大小
	Column             []string `json:"column"`                                       // 同步的列
	CronExpression     string   `json:"cronExpression"`                               // Cron表达式
	Status             uint     `json:"status"`                                        // 状态
}

// GenerateSeatunnelScriptRequest 生成 SeaTunnel 脚本请求
type GenerateSeatunnelScriptRequest struct {
	// 任务名称
	TaskName string `json:"taskName" binding:"required"`
	// 源数据源ID
	SourceDataSourceId uint `json:"sourceDataSourceId" binding:"required"`
	// 源数据库
	SourceDatabase string `json:"sourceDatabase" binding:"required"`
	// 源表
	SourceTable string `json:"sourceTable" binding:"required"`
	// 目标数据源ID
	TargetDataSourceId uint `json:"targetDataSourceId" binding:"required"`
	// 目标数据库
	TargetDatabase string `json:"targetDatabase" binding:"required"`
	// 目标表
	TargetTable string `json:"targetTable" binding:"required"`
	// 并行度（可选，默认4）
	Parallelism int `json:"parallelism"`
	// 源批处理大小（可选，默认5000）
	BatchSize int `json:"batchSize"`
	// 要同步的列（可选，默认空数组表示全部列）
	Column []string `json:"column"`
	// 分割大小（可选，默认8096）
	SplitSize int `json:"splitSize"`
	// 更新模式（可选，默认INSERT）
	UpdateMode string `json:"updateMode"`
	// 数据保存模式（可选，默认OVERWRITE）
	DataSaveMode string `json:"dataSaveMode"`
	// 目标批处理大小（可选，默认1024）
	SinkBatchSize int `json:"sinkBatchSize"`
}

// GetParallelism 获取并行度（默认值4）
func (r *GenerateSeatunnelScriptRequest) GetParallelism() int {
	if r.Parallelism <= 0 {
		return 4
	}
	return r.Parallelism
}

// GetBatchSize 获取批处理大小（默认值5000）
func (r *GenerateSeatunnelScriptRequest) GetBatchSize() int {
	if r.BatchSize <= 0 {
		return 5000
	}
	return r.BatchSize
}

// GetColumn 获取列（默认空数组）
func (r *GenerateSeatunnelScriptRequest) GetColumn() []string {
	if r.Column == nil {
		return []string{}
	}
	return r.Column
}

// GetSplitSize 获取分割大小（默认值8096）
func (r *GenerateSeatunnelScriptRequest) GetSplitSize() int {
	if r.SplitSize <= 0 {
		return 8096
	}
	return r.SplitSize
}

// GetUpdateMode 获取更新模式（默认INSERT）
func (r *GenerateSeatunnelScriptRequest) GetUpdateMode() string {
	if r.UpdateMode == "" {
		return "INSERT"
	}
	return r.UpdateMode
}

// GetDataSaveMode 获取数据保存模式（默认OVERWRITE）
func (r *GenerateSeatunnelScriptRequest) GetDataSaveMode() string {
	if r.DataSaveMode == "" {
		return "OVERWRITE"
	}
	return r.DataSaveMode
}

// GetSinkBatchSize 获取目标批处理大小（默认值1024）
func (r *GenerateSeatunnelScriptRequest) GetSinkBatchSize() int {
	if r.SinkBatchSize <= 0 {
		return 1024
	}
	return r.SinkBatchSize
}

// DeleteByIdRequest 删除请求
type DeleteByIdRequest struct {
	ID uint `json:"id" binding:"required"`
}

// GetByIdRequest 获取详情请求
type GetByIdRequest struct {
	ID uint `json:"id" binding:"required"`
}

// PageInfo 分页查询
type PageInfo struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Keyword  string `json:"keyword" form:"keyword"`
}

// DataSourcePageInfo 数据源分页查询（支持类型过滤）
type DataSourcePageInfo struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Keyword  string `json:"keyword" form:"keyword"`
	Type     string `json:"type" form:"type"`
}