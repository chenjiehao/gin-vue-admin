# Phase 11-01: form-split-panel - Summary

**Completed:** 2026-04-22

## Goal

将"新增/编辑"操作从全页路由跳转改为在列表页右侧 div 区域局部加载，左侧列表保持显示，右侧动态切换列表/表单。

## What Was Done

将 offlineSync.vue 和 dataSource.vue 两个列表页重构为分栏布局：

1. **offlineSync.vue** — 列表页内嵌表单
   - 页面初始只显示左侧搜索+表格+分页，右侧隐藏
   - 点击"新增任务"：右侧显示表单，左侧收缩
   - 点击"编辑"：右侧显示预填数据表单
   - 表单提交/取消后自动返回列表视图

2. **dataSource.vue** — 列表页内嵌表单
   - 同样的分栏布局模式
   - 表单内联到列表页，不再路由跳转

## Files Modified

- `web/src/view/dataIntegration/offlineSync.vue` — 重构为分栏布局，内嵌表单
- `web/src/view/dataIntegration/dataSource.vue` — 重构为分栏布局，内嵌表单

## Key Changes

### 布局变化
```
之前: router.push → 新页面
现在: 列表页左右分栏，右侧动态显示/隐藏表单
```

### 状态管理
- 新增 `formVisible` 控制表单显示/隐藏
- 新增 `isEdit` 区分新增/编辑模式
- 新增 `closeForm()` 关闭表单并重置状态
- 新增 `goCreate()` / `goEdit()` 打开表单

### 样式
- `.page-container` flex 容器
- `.list-panel` 列表区域，`collapsed` 时宽度收缩
- `.form-panel` 表单区域，`visible` 时显示

## Notes

- 两个独立表单页（offlineSyncForm.vue, dataSourceForm.vue）保留未删除，因为 router/index.js 中仍有静态路由注册
- 如果后续确认不再需要独立表单页，可删除这两个文件和 router/index.js 中的静态路由
