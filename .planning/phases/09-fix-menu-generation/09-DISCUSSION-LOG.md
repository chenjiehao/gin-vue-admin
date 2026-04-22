# Phase 09: fix-menu-generation - Discussion Log

> **Audit trail only.** Do not use as input to planning, research, or execution agents.
> Decisions are captured in CONTEXT.md — this log preserves the alternatives considered.

**Date:** 2026-04-22
**Phase:** 09-fix-menu-generation
**Areas discussed:** 修复策略, 识别机制, 清理范围

---

## 修复策略

| Option | Description | Selected |
|--------|-------------|----------|
| 直接删除问题菜单记录 | 从 sys_base_menus 删除 OfflineSyncForm 和 dataSourceForm 的记录 | ✓ |
| 修复新建菜单的生成逻辑 | 确保未来创建类似页面时不会错误创建为菜单项 | |
| 两者都要修复 | 既删除现有错误记录，也修复新建逻辑 | ✓ |

**User's choice:** 两者都要修复
**Notes:** 用户确认需要同时处理现有问题和预防未来问题

---

## 识别机制

| Option | Description | Selected |
|--------|-------------|----------|
| 强制命名规范 | 规定表单页必须以 Form 结尾，系统检测到 *Form.vue 就不创建菜单项 | |
| 手动标记机制 | 在 pathInfo.json 或菜单创建时显式标记 isMenu=true/false | |
| 表单不走菜单创建流程 | 表单页根本不录入系统，完全由代码层面 router.push 访问 | ✓ |

**User's choice:** 表单不走菜单创建流程
**Notes:** 用户明确表示表单页不应该成为菜单项，只应该是可通过 router.push 访问的路由

---

## 清理范围

| Option | Description | Selected |
|--------|-------------|----------|
| 只清理已知的两个 | 只删除 OfflineSyncForm 和 dataSourceForm，其他以后再说 | ✓ |
| 检查全部二级菜单 | 扫描 sys_base_menus 中所有 parent_id != 0 的记录 | |

**User's choice:** 只清理已知的两个
**Notes:** 用户选择聚焦在已知问题上，不扩大范围

---

## Deferred Ideas

- 系统级菜单全面检查 — 留待后续phase处理

