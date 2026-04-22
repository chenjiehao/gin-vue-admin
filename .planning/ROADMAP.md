# Roadmap: gin-vue-admin 精简脚手架

## Milestones

- ✅ **v1.0 精简脚手架** — Phases 1-3 (shipped 2026-04-21)
- ✅ **v1.1 功能增强** — Phase 4 (shipped 2026-04-22)
- ✅ **v1.2 菜单与数据集成** — Phase 5 (shipped 2026-04-22)
- ✅ **v1.3 数据集成 - 离线同步** — Phases 6-11 (shipped 2026-04-22)

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
<summary>✅ v1.3 数据集成 - 离线同步 — SHIPPED 2026-04-22</summary>

- [x] Phase 6: 离线同步功能 (1/1 plans) — completed 2026-04-22
  - [x] 06-01-PLAN.md — 创建离线同步Vue页面 + API stubs + 菜单
- [x] Phase 7: 优化离线同步默认列表展示 (1/1 plans) — completed 2026-04-22
  - [x] 07-01-PLAN.md — 修复 permission.js 路由重定向逻辑，hidden 子菜单不再导致异常跳转
- [x] Phase 8: 数据源管理 (1/1 plans) — completed 2026-04-22
  - [x] 08-01-PLAN.md — 在数据集成菜单下新建数据源管理二级菜单
- [x] Phase 9: 修复菜单生成逻辑 (1/1 plans) — completed 2026-04-22
  - [x] 09-01-PLAN.md — 清理错误创建的表单页菜单记录
- [x] Phase 10: 修复数据源和离线同步表单页无法访问的问题 (1/1 plans) — completed 2026-04-22
  - [x] 10-01-PLAN.md — 注册表单页静态路由
- [x] Phase 11: 将表单页从全页跳转改为整块替换加载 (1/1 plans) — completed 2026-04-22
  - [x] 11-01-PLAN.md — 重构列表页，整块替换为表单视图，多列网格布局

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
| 9. 修复菜单生成逻辑：禁止为二级菜单下的非菜单页面自动添加菜单项 | v1.3 | 1/1 | Complete | 2026-04-22 |
| 10. 修复数据源和离线同步表单页无法访问的问题 | v1.3 | 1/1 | Complete | 2026-04-22 |
| 11. 将表单页从全页跳转改为右侧 div 局部加载 | v1.3 | 1/1 | Complete | 2026-04-22 |
| 12. 数据源管理的增强 | v1.3 | 2/3 | Complete    | 2026-04-22 |
| 13. 数据源表列表获取 | v1.3 | 2/2 | Complete | 2026-04-22 |

### Phase 12: 数据源管理的增强

**Goal:** 增强数据源管理功能 — 在 Phase 8/11 完成的 CRUD 基础上，新增连接测试和批量操作能力
**Requirements**: DS-01, DS-02, DS-03, DS-04, DS-05, DS-06, DS-07
**Depends on:** Phase 11
**Plans:** 3/3 plans complete

Plans:
- [x] 12-01-PLAN.md — 后端连接测试 API + Service 层逻辑 (Wave 1)
- [x] 12-02-PLAN.md — 后端批量操作 API + Service 层逻辑 (Wave 2)
- [x] 12-03-PLAN.md — 前端连接测试 UI + 批量操作 UI (Wave 3)

### Phase 13: 数据源表列表获取

**Goal:** 新增后端 API 获取数据源的表列表，前端联动实现离线同步任务配置时的表选择
**Requirements**: DS-08
**Depends on:** Phase 12
**Plans:** 1/2 plans executed

Plans:
- [x] 13-01-PLAN.md — 后端 GetTables API + Service 层逻辑 (Wave 1)
- [ ] 13-02-PLAN.md — 前端表选择下拉框联动 (Wave 2)

---

_For milestone details, see [.planning/milestones/v1.0-ROADMAP.md](.planning/milestones/v1.0-ROADMAP.md)_