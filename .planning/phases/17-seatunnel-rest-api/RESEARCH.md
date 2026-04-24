# Phase 17 Research: Apache SeaTunnel v2 REST API

## 研究目标
了解 Apache SeaTunnel v2 的 REST API，为后续实现任务提交和执行监控做准备。

## SeaTunnel Engine v2 REST API 概述

SeaTunnel Engine 提供 REST API 用于：
- 作业提交 (Job Submission)
- 作业监控 (Job Monitoring)
- 作业控制 (Job Control: pause, resume, stop)
- 系统信息 (System Info)

### 典型 REST API 端点

#### 1. 作业提交
```
POST /seatunnel/submit
Content-Type: application/json

{
  "env": {...},
  "source": {...},
  "sink": {...}
}
```

#### 2. 获取作业状态
```
GET /seatunnel/job/:jobId/status
```

#### 3. 获取作业信息
```
GET /seatunnel/job/:jobId/info
```

#### 4. 作业控制
```
POST /seatunnel/job/:jobId/stop
POST /seatunnel/job/:jobId/pause
POST /seatunnel/job/:jobId/resume
```

#### 5. 系统信息
```
GET /seatunnel/system/info
GET /seatunnel/system/metrics
```

## 待研究问题

- [ ] SeaTunnel Engine 的认证方式 (API Key / Basic Auth / None)
- [ ] API 基础 URL 配置方式
- [ ] 作业提交后返回的响应格式
- [ ] 实时日志获取方式
- [ ] 作业状态流转 (pending → running → finished/failed)

## 数据结构

### Job Status 枚举
- `WAITING` - 等待中
- `RUNNING` - 运行中
- `FINISHED` - 已完成
- `FAILED` - 失败
- `CANCELED` - 已取消
- `PAUSED` - 已暂停

### 典型作业响应
```json
{
  "jobId": 123,
  "status": "RUNNING",
  "createTime": "2026-04-24T10:00:00",
  "metrics": {
    "sourceProcessedCount": 1000,
    "sinkProcessedCount": 1000,
    "sourceFailedCount": 0,
    "sinkFailedCount": 0
  }
}
```

## 行动计划
1. 确认 SeaTunnel Engine 部署方式 (standalone / cluster)
2. 获取 API 文档
3. 测试作业提交接口
4. 实现后端定时任务调用接口
5. 实现作业状态轮询监控

## 参考资料
- 待补充：SeaTunnel 官方文档链接
