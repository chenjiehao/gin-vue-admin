---
status: testing
phase: 06-offline-sync
source: 06-01-SUMMARY.md
started: 2026-04-22T03:15:00.000Z
updated: 2026-04-22T03:15:00.000Z
---

## Current Test

number: 1
name: 离线同步列表页面显示
expected: |
  访问数据集成 > 离线同步子菜单，看到任务列表页面，显示6条模拟任务数据（用户数据同步、订单数据备份、日志采集、商品数据同步、文件归档任务、第三方API数据拉取），状态标签显示正确颜色（成功=绿色，运行中=黄色，待执行=灰色，失败=红色）
awaiting: user response

## Tests

### 1. 离线同步列表页面显示
expected: 访问数据集成 > 离线同步子菜单，看到任务列表页面，显示6条模拟任务数据（用户数据同步、订单数据备份、日志采集、商品数据同步、文件归档任务、第三方API数据拉取），状态标签显示正确颜色（成功=绿色，运行中=黄色，待执行=灰色，失败=红色）
result: pending

### 2. 搜索功能
expected: 输入任务名称"用户"，点击查询，列表筛选为2条（用户数据同步、第三方API数据拉取）；状态筛选选择"成功"，显示3条成功状态任务
result: pending

### 3. 分页功能
expected: 列表底部显示分页组件，当前页码1，每页10条，总数6
result: pending

### 4. 新增任务跳转
expected: 点击"新增任务"按钮，页面跳转到 /dataIntegration/offlineSyncForm，显示任务表单页面，顶部有"返回列表"按钮
result: pending

### 5. 编辑任务跳转
expected: 点击某行的"编辑"按钮，页面跳转到 /dataIntegration/offlineSyncForm?id=xxx，表单中预填充了该任务的数据
result: pending

### 6. 返回列表功能
expected: 在表单页面点击"返回列表"按钮，页面返回到任务列表页，数据完整显示
result: pending

### 7. 历史记录对话框
expected: 点击某行的"历史"按钮，弹出对话框显示该任务的同步执行历史记录（执行时间、状态、耗时、同步数量、错误信息）
result: pending

### 8. 触发同步功能
expected: 点击某行的"触发"按钮，弹出确认框，点击确定后显示"触发成功"提示
result: pending

### 9. 删除任务功能
expected: 点击某行的"删除"按钮，弹出确认框，点击确定后显示"删除成功"提示，列表刷新
result: pending

## Summary

total: 9
passed: 0
issues: 0
pending: 9
skipped: 0

## Gaps

[none yet]
