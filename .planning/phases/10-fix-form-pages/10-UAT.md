---
status: testing
phase: 10-fix-form-pages
source: 10-01-PLAN.md
started: 2026-04-22T15:10:00Z
updated: 2026-04-22T15:10:00Z
---

## Current Test

number: 1
name: 测试新增离线同步任务按钮
expected: |
  点击"新增任务"按钮后，页面跳转到 OfflineSyncForm，URL 类似 #/dataIntegration/offlineSyncForm
awaiting: 用户确认

## Tests

### 1. 测试新增离线同步任务按钮
expected: 点击"新增任务"按钮后，OfflineSyncForm 表单页正常显示（不是404）
result: pending

### 2. 测试编辑离线同步任务按钮
expected: 点击"编辑"按钮后，页面跳转到 OfflineSyncForm?id=xxx，表单加载对应数据
result: pending

### 3. 测试新增数据源按钮
expected: 点击"新增数据源"按钮后，dataSourceForm 表单页正常显示（不是404）
result: pending

### 4. 测试编辑数据源按钮
expected: 点击"编辑"按钮后，页面跳转到 dataSourceForm?id=xxx，表单加载对应数据
result: pending

## Summary

total: 4
passed: 0
issues: 0
pending: 4
skipped: 0
blocked: 0

## Gaps

