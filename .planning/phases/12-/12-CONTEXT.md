# Phase 12: 数据源管理的增强 - Context

**Gathered:** 2026-04-22
**Status:** Ready for planning

<domain>
## Phase Boundary

增强数据源管理功能 — 在 Phase 8/11 完成的 CRUD 基础上，新增连接测试和批量操作能力。

**当前已有功能：**
- 数据源列表（名称、类型、主机、端口、数据库、状态、备注）
- 新增/编辑/删除数据源
- 表单内嵌在列表页（Phase 11 成果）

**本 phase 新增：**
- 连接测试功能
- 批量操作（删除、启用/禁用）

</domain>

<decisions>
## Implementation Decisions

### D-01: 连接测试 — 触发时机
- **决定：** 手动点击"测试连接"按钮，不在保存时自动测试
- **原因：** 避免保存时强制等待连接超时；用户可自行选择时机测试

### D-02: 连接测试 — 结果展示
- **决定：** 成功/失败 + 详细错误信息（连接超时、认证失败、数据库不存在等）
- **原因：** 详细信息帮助用户定位配置问题

### D-03: 连接测试 — 实现方式
- **决定：** 前端点击按钮 → 调用后端 API → 后端尝试建立真实连接 → 返回详细结果
- **后端 API：** `POST /dataSource/testConnection` — 接收数据源配置，返回 `{success: bool, message: string}`

### D-04: 批量操作 — 操作类型
- **决定：** 支持批量删除 + 批量启用/禁用
- **原因：** 两者都是高频管理操作

### D-05: 批量操作 — 确认方式
- **决定：** 弹窗确认，显示选中数量（"确定删除选中的 3 个数据源吗？"）
- **原因：** 用户需明确知道影响范围

### D-06: 批量操作 — 反馈方式
- **决定：** 成功后刷新列表；失败时 ElMessage 提示失败数量，不逐个显示
- **原因：** 简化 UX，失败场景较少见

### D-07: 前端交互 — 按钮位置
- **决定：** 批量操作按钮（批量删除、批量启用、批量禁用）放在表格上方右侧
- **原因：** 与现有"新增数据源"按钮在同一行，位置醒目

</decisions>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

- `web/src/view/dataIntegration/dataSource.vue` — 数据源列表页（参考现有布局）
- `web/src/api/dataSource.js` — 现有 API 定义（需新增 testConnection 和批量 API）
- `server/api/v1/system/data_source.go` — 后端 API（需新增 testConnection 接口）
- `server/service/system/data_source.go` — 后端服务层（需实现连接测试逻辑）

</canonical_refs>

<codebase_context>
## Existing Code Insights

### Reusable Assets
- `dataSource.vue` — 现有列表页，可直接复用搜索区、表格、分页布局
- `form-view` 样式 — Phase 11 的内嵌表单样式，可复用
- `ElMessage` / `ElMessageBox` — 统一的消息和确认框

### Integration Points
- 后端需新增 `/dataSource/testConnection` API
- 后端需新增 `/dataSource/batchDelete` 和 `/dataSource/batchUpdateStatus` API
- 前端 `dataSource.vue` 需新增批量选择逻辑和批量操作按钮

### Batch Operation Logic
```javascript
// 多选后的操作
const batchDelete = () => {
  const ids = multipleSelection.value.map(item => item.id)
  // 调用批量删除 API
}
```

</codebase_context>

<specifics>
## Specific Ideas

- 连接测试后端实现：使用各数据库官方 driver 的 Ping/Connect 方法
- 批量操作需处理空选择情况（未选中时按钮禁用或提示）

</specifics>

<deferred>
## Deferred Ideas

- 字段扩展（超时时间、连接池大小、字符编码）— 属于下一个增强方向
- 分类标签（数据源分组）— 属于下一个增强方向

</deferred>

---

*Phase: 12-data-source-enhancement*
*Context gathered: 2026-04-22*
