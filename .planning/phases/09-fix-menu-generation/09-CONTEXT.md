# Phase 09: fix-menu-generation - Context

**Gathered:** 2026-04-22
**Status:** Ready for planning

<domain>
## Phase Boundary

修复菜单生成逻辑 — 禁止为二级菜单下的非菜单页面（如表单页）自动添加菜单项。

**问题现象：**
- `OfflineSyncForm` (id=14) 和 `dataSourceForm` (id=16) 被错误地创建为独立的菜单项
- 它们以 `hidden=1` 隐藏，但仍然存在于 `sys_base_menus` 表中
- 这导致它们成为"菜单项"，只是不显示在侧边栏

**预期行为：**
- 只有明确要创建为菜单的才会生成菜单项
- 表单页（如 `*Form.vue`）不应该成为菜单项，只应该是可以通过 `router.push` 访问的路由

**修复范围：**
- 仅限数据集成下的 `offlineSyncForm` 和 `dataSourceForm`
- 不涉及后端 API、数据库结构、其他菜单层级

**约束：**
- 不修改已有页面的菜单状态（除非本身属于误添加）
- 不改变路由访问路径，仅影响菜单显示

</domain>

<decisions>
## Implementation Decisions

### D-01: 清理现有错误菜单记录
- **决定：** 从 `sys_base_menus` 删除 `OfflineSyncForm` (id=14) 和 `dataSourceForm` (id=16) 的记录
- **原因：** 它们不应该作为菜单项存在，只是表单页

### D-02: 表单页不走菜单创建流程
- **决定：** `*Form.vue` 页面完全不走菜单创建流程
- **原因：** 表单页只应该作为代码层面通过 `router.push` 访问的路由，不应该成为菜单项
- **实现：** 后续开发中，表单页不录入任何菜单创建流程

### D-03: 路由仍可正常访问
- **决定：** 删除菜单记录后，表单页仍可通过 `router.push({ name: 'offlineSyncForm' })` 正常访问
- **原因：** 只是移除菜单项，不影响路由本身

### D-04: 清理范围
- **决定：** 仅清理 `offlineSyncForm` 和 `dataSourceForm` 两个已知的错误记录
- **不涉及：** 系统其他菜单的检查

</decisions>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

- `server/model/system/sys_base_menu.go` — 菜单数据模型
- `web/src/view/dataIntegration/offlineSync.vue` — 离线同步列表页（参考）
- `web/src/view/dataIntegration/offlineSyncForm.vue` — 离线同步表单页（待清理）
- `web/src/view/dataIntegration/dataSource.vue` — 数据源列表页（参考）
- `web/src/view/dataIntegration/dataSourceForm.vue` — 数据源表单页（待清理）
- `.planning/ROADMAP.md` — Phase 9 定义

</canonical_refs>

<codebase_context>
## Existing Code Insights

### Database Records to Delete
```
id=14: OfflineSyncForm  parent_id=11  hidden=1  (需删除)
id=16: dataSourceForm  parent_id=11  hidden=1  (需删除)
```

### Reusable Assets
- `router.push({ name: 'xxxForm', query: { id: row.id } })` — 表单页的标准跳转方式（已有模式）

### Integration Points
- `sys_base_menus` 表 — 存储菜单记录
- `sys_authority_menus` 表 — 存储菜单与角色的关联（需同步清理）

</codebase_context>

<specifics>
## Specific Ideas

- 删除菜单记录时，需要同时检查并删除 `sys_authority_menus` 中的关联记录
- 前端 `pathInfo.json` 中的映射可以保留（用于路由名解析），不影响菜单显示

</specifics>

<deferred>
## Deferred Ideas

- 系统级菜单全面检查 — 留待后续phase处理

</deferred>

---

*Phase: 09-fix-menu-generation*
*Context gathered: 2026-04-22*
