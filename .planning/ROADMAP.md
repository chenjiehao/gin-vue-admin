# Roadmap: gin-vue-admin 精简脚手架

## Milestones

- ✅ **v1.0 精简脚手架** — Phases 1-3 (shipped 2026-04-21)
- ✅ **v1.1 功能增强** — Phase 4 (shipped 2026-04-22)
- ✅ **v1.2 菜单与数据集成** — Phase 5 (shipped 2026-04-22)
- 🔄 **v1.3 数据集成 - 离线同步** — Phase 6 (in progress)

## Phases

<details>
<summary>✅ v1.0 精简脚手架 (Phases 1-3) — SHIPPED 2026-04-21</summary>

- [x] Phase 1: Backend Cleanup (3/3 plans) — completed 2026-04-21
- [x] Phase 2: Frontend Cleanup (2/2 plans) — completed 2026-04-21
- [x] Phase 3: Data Cleanup (1/1 plans) — completed 2026-04-21

</details>

<details>
<summary>✅ v1.1 功能增强 (Phase 4) — SHIPPED 2026-04-22</summary>

- [x] Phase 4: 首页仪表盘 UI 重设计 (1/1 plans) — completed 2026-04-22

</details>

<details>
<summary>✅ v1.2 菜单与数据集成 (Phase 5) — SHIPPED 2026-04-22</summary>

- [x] Phase 5: 恢复审计日志菜单 + 新增数据集成模版一级菜单 (1/1 plans) — completed 2026-04-22
- [x] Phase 5 (额外): 数据集成模版菜单优化 — 2026-04-22

</details>

<details>
<summary>🔄 v1.3 数据集成 - 离线同步 — IN PROGRESS</summary>

- [x] Phase 6: 离线同步功能 (1/1 plans) — completed 2026-04-22
  - [x] 06-01-PLAN.md — 创建离线同步Vue页面 + API stubs + 菜单
- [x] Phase 7: 优化离线同步默认列表展示 (1/1 plans) — completed 2026-04-22
  - [x] 07-01-PLAN.md — 修复 permission.js 路由重定向逻辑，hidden 子菜单不再导致异常跳转
- [x] Phase 8: 数据源管理 (1/1 plans) — completed 2026-04-22
  - [x] 08-01-PLAN.md — 在数据集成菜单下新建数据源管理二级菜单
- [ ] Phase 9: 修复菜单生成逻辑 (1/1 plans) — pending
  - [ ] 09-01-PLAN.md — 清理错误创建的表单页菜单记录

</details>

## Progress

| Phase | Milestone | Plans Complete | Status | Completed |
|-------|-----------|----------------|--------|-----------|
| 1. Backend Cleanup | v1.0 | 3/3 | Complete | 2026-04-21 |
| 2. Frontend Cleanup | v1.0 | 2/2 | Complete | 2026-04-21 |
| 3. Data Cleanup | v1.0 | 1/1 | Complete | 2026-04-21 |
| 4. 首页仪表盘 UI 重设计 | v1.1 | 1/1 | Complete | 2026-04-22 |
| 5. 恢复审计日志菜单 + 新增数据集成模版一级菜单 | v1.2 | 1/1 | Complete | 2026-04-22 |
| 6. 离线同步功能 | v1.3 | 1/1 | Complete | 2026-04-22 |
| 7. 优化离线同步默认列表展示 | v1.3 | 1/1 | Complete | 2026-04-22 |
| 8. 数据源管理 | v1.3 | 1/1 | Complete | 2026-04-22 |
| 9. 修复菜单生成逻辑：禁止为二级菜单下的非菜单页面自动添加菜单项 | v1.3 | 0/1 | Not planned | — |

### Phase 9: 修复菜单生成逻辑：禁止为二级菜单下的非菜单页面自动添加菜单项

**Goal:** 从数据库删除 offlineSyncForm (id=14) 和 dataSourceForm (id=16) 的错误菜单记录
**Requirements:** N/A (数据清理任务)
**Depends on:** Phase 8
**Plans:** 1 plan

Plans:
- [ ] 09-01-PLAN.md — 清理错误创建的表单页菜单记录

---

_For milestone details, see [.planning/milestones/v1.0-ROADMAP.md](.planning/milestones/v1.0-ROADMAP.md)_
