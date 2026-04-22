# Phase 12: 数据源管理的增强 - Discussion Log

> **Audit trail only.** Do not use as input to planning, research, or execution agents.
> Decisions are captured in CONTEXT.md — this log preserves the alternatives considered.

**Date:** 2026-04-22
**Phase:** 12-数据源管理的增强
**Areas discussed:** 连接测试功能, 批量操作

---

## 连接测试功能

| Option | Description | Selected |
|--------|-------------|----------|
| 保存时自动测试 | 每次保存前自动测试连接，用户等待 | |
| 手动点击按钮测试 | 用户自行点击"测试连接"按钮 | ✓ |
| 两者都要 | 保存时自动测试 + 手动也可测试 | |

**User's choice:** 技术决策
**Notes:** 避免保存时强制等待连接超时；用户可自行选择时机测试

| Option | Description | Selected |
|--------|-------------|----------|
| 成功/失败两种状态 | 简单状态提示 | |
| 成功/失败 + 详细错误信息 | 显示连接超时、认证失败等具体原因 | ✓ |

**User's choice:** 技术决策
**Notes:** 详细信息帮助用户定位配置问题

---

## 批量操作

| Option | Description | Selected |
|--------|-------------|----------|
| 批量删除 | 只支持批量删除 | |
| 批量删除 + 启用/禁用 | 支持删除和状态切换 | ✓ |

**User's choice:** 技术决策
**Notes:** 两者都是高频管理操作

| Option | Description | Selected |
|--------|-------------|----------|
| 弹窗确认（显示数量） | "确定删除选中的 3 个数据源吗？" | ✓ |
| 逐个确认 | 逐个弹窗确认 | |

**User's choice:** 技术决策
**Notes:** 用户需明确知道影响范围

| Option | Description | Selected |
|--------|-------------|----------|
| 显示每个操作的成功/失败 | 列表展示每个结果 | |
| 成功后刷新，失败时提示数量 | 简化 UX | ✓ |

**User's choice:** 技术决策
**Notes:** 失败场景较少见，不逐个显示

---

## Claude's Discretion

- 连接测试的触发时机、结果展示、实现方式 — 技术决策
- 批量操作类型、确认方式、反馈方式 — 技术决策
- 前端按钮位置 — 技术决策

## Deferred Ideas

- 字段扩展（超时时间、连接池大小、字符编码）
- 分类标签（数据源分组）
