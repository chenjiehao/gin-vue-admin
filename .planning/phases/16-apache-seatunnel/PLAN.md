# Phase 16: 实现Apache SeaTunnel任务脚本生成

## 目标

后端根据离线同步配置生成 Apache SeaTunnel 任务 JSON 脚本。

## SeaTunnel JSON 结构分析

### 目标结构
```json
{
  "env": {
    "job": { "name": "taskName", "mode": "BATCH" },
    "parallelism": 4,
    "batchSize": 5000
  },
  "source": {
    "MySQL": {
      "baseUrl": "jdbc:mysql://host:port",
      "username": "xxx",
      "password": "xxx",
      "database": "xxx",
      "table": "xxx",
      "column": [],
      "splitSize": 8096
    }
  },
  "sink": {
    "MySQL": {
      "baseUrl": "jdbc:mysql://host:port",
      "username": "xxx",
      "password": "xxx",
      "database": "xxx",
      "table": "xxx",
      "updateMode": "INSERT",
      "dataSaveMode": "OVERWRITE",
      "batchSize": 1024
    }
  }
}
```

## 前端现状 vs 需求对比

| SeaTunnel 字段 | 前端表单字段 | 状态 |
|----------------|-------------|------|
| env.job.name | taskName | ✓ 已有 |
| env.parallelism | - | ✗ 需添加或用默认值 4 |
| env.batchSize | - | ✗ 需添加或用默认值 5000 |
| source[type].baseUrl | sourceDataSourceId | ✓ 通过数据源查找 |
| source[type].username | sourceDataSourceId | ✓ 通过数据源查找 |
| source[type].password | sourceDataSourceId | ✓ 通过数据源查找 |
| source[type].database | sourceDatabase | ✓ 已有 |
| source[type].table | sourceTable | ✓ 已有 |
| source[type].column | - | ✗ 需添加或用默认值 [] |
| source[type].splitSize | - | ✗ 需添加或用默认值 8096 |
| sink[type].baseUrl | targetDataSourceId | ✓ 通过数据源查找 |
| sink[type].username | targetDataSourceId | ✓ 通过数据源查找 |
| sink[type].password | targetDataSourceId | ✓ 通过数据源查找 |
| sink[type].database | targetDatabase | ✓ 已有 |
| sink[type].table | targetTable | ✓ 已有 |
| sink[type].updateMode | - | ✗ 需添加或用默认值 "INSERT" |
| sink[type].dataSaveMode | - | ✗ 需添加或用默认值 "OVERWRITE" |
| sink[type].batchSize | - | ✗ 需添加或用默认值 1024 |

## 实现计划

### 任务 1: 前端 - 添加 SeaTunnel 配置字段

前端表单增加以下字段（可选，使用默认值）：

1. **env.parallelism** - 同步并行度（默认 4）
2. **env.batchSize** - 批处理大小（默认 5000）
3. **source.column** - 要同步的列（默认空数组）
4. **source.splitSize** - 分割大小（默认 8096）
5. **sink.updateMode** - 更新模式（默认 INSERT）
6. **sink.dataSaveMode** - 数据保存模式（默认 OVERWRITE）
7. **sink.batchSize** - 批处理大小（默认 1024）

### 任务 2: 后端 - 创建 SeaTunnel 脚本生成服务

新建 `SeatunnelScriptGenerator` 或在现有服务中增加方法：

1. 根据离线同步配置构建 SeaTunnel JSON
2. 从数据源配置获取 baseUrl、username、password
3. 组装 env、source、sink 三个部分

### 任务 3: 后端 - 创建 API 接口

新增 API：
- `POST /offlineSync/generateScript` - 根据同步任务生成 SeaTunnel 脚本

## API 设计

### Request
```json
{
  "taskName": "orders-sync",
  "sourceDataSourceId": 1,
  "sourceDatabase": "orders_db",
  "sourceTable": "orders",
  "targetDataSourceId": 2,
  "targetDatabase": "dwd_db",
  "targetTable": "orders",
  // 以下字段可选，使用默认值
  "parallelism": 4,
  "batchSize": 5000,
  "column": [],
  "splitSize": 8096,
  "updateMode": "INSERT",
  "dataSaveMode": "OVERWRITE",
  "sinkBatchSize": 1024
}
```

### Response
```json
{
  "code": 0,
  "data": {
    "script": { /* SeaTunnel JSON 结构 */ }
  },
  "msg": "生成成功"
}
```

## 文件变更

### 后端新增
- `server/model/request/offline_sync.go` - 请求模型
- `server/service/system/seatunnel_script.go` - 脚本生成服务
- `server/api/v1/system/offline_sync.go` - API 接口（如不存在）

### 后端修改
- `server/service/system/offline_sync.go` - 增加生成脚本方法
- `server/router/system/offline_sync.go` - 注册路由

### 前端修改
- `web/src/view/dataIntegration/offlineSync.vue` - 增加配置表单
- `web/src/api/offlineSync.js` - 增加 generateScript API
