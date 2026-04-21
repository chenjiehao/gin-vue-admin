# Roadmap: gin-vue-admin 精简脚手架

## Overview

将 gin-vue-admin 框架精简为最小可用的 RBAC 管理系统骨架。Phase 1 清理后端示例模块并保留核心框架；Phase 2 清理前端业务页面；Phase 3 清理数据库初始化数据中的示例数据。

## Phases

- [ ] **Phase 1: Backend Cleanup** - 删除后端示例模块，保留核心框架
- [ ] **Phase 2: Frontend Cleanup** - 删除前端业务页面，保留框架骨架
- [ ] **Phase 3: Data Cleanup** - 清理数据库初始化中的示例数据

## Phase Details

### Phase 1: Backend Cleanup
**Goal**: 后端所有示例模块已删除，核心框架功能完整保留
**Depends on**: Nothing (first phase)
**Requirements**: CLEAN-01, CLEAN-02, CLEAN-03, CLEAN-04, CLEAN-07, CLEAN-08, CORE-01, CORE-02, CORE-03, CORE-04, CORE-05, CORE-06, CORE-07, CORE-08
**Success Criteria** (what must be TRUE):
  1. server/model/example/ 目录不存在，无示例模型文件
  2. server/api/v1/example/ 目录不存在，无示例API文件
  3. server/service/example/ 目录不存在，无示例服务文件
  4. server/router/example/ 目录不存在，无示例路由文件
  5. server/plugin/announcement/ 目录不存在
  6. server/plugin/email/ 目录不存在
  7. 用户认证API可正常工作（登录、登出、获取用户信息）
  8. 角色管理API可正常工作（CRUD、权限关联）
  9. 菜单管理API可正常工作（CRUD、按钮权限）
  10. API管理API可正常工作（CRUD、角色关联）
  11. JWT + Casbin 鉴权体系正常工作
  12. 插件系统入口可正常挂载
  13. Swagger 文档可访问
  14. 数据库自动迁移正常工作
**Plans**: 3 plans

Plans:
- [ ] 01-01: 删除 server/model/example/、server/api/v1/example/、server/service/example/、server/router/example/
- [ ] 01-02: 删除 server/plugin/announcement/ 和 server/plugin/email/
- [ ] 01-03: 验证核心框架功能（用户、角色、菜单、API管理 + JWT + Casbin + Swagger）

### Phase 2: Frontend Cleanup
**Goal**: 前端所有示例页面和API已删除，只保留框架登录页和主布局
**Depends on**: Phase 1
**Requirements**: FRONT-01, FRONT-02, FRONT-03, FRONT-04, FRONT-05
**Success Criteria** (what must be TRUE):
  1. web/src/view/example/ 目录不存在，无示例页面
  2. web/src/api/example.js 文件不存在
  3. 登录页面可正常加载和提交
  4. 主布局（侧边栏、顶栏、内容区）可正常渲染
  5. 框架API（用户、角色、菜单、API）可正常调用
  6. 动态路由加载正常，框架菜单可显示
  7. 导航中无业务/示例页面入口
**Plans**: 2 plans

Plans:
- [ ] 02-01: 删除 web/src/view/example/ 和 web/src/api/example.js
- [ ] 02-02: 验证前端框架功能（登录、主布局、动态路由、框架API调用）

### Phase 3: Data Cleanup
**Goal**: 数据库初始化数据已清理，无示例菜单、角色、用户数据
**Depends on**: Phase 2
**Requirements**: DATA-01, DATA-02, DATA-03, DATA-04
**Success Criteria** (what must be TRUE):
  1. 初始化数据中不包含示例菜单
  2. 初始化数据中不包含示例角色
  3. 初始化数据中不包含示例用户（除超级管理员外）
  4. 超级管理员默认账号（admin/123456）可正常登录
**Plans**: 1 plan

Plans:
- [ ] 03-01: 清理 source/ 目录中的示例数据（菜单、角色、用户），保留超级管理员

## Progress

**Execution Order:**
Phases execute in numeric order: 1 → 2 → 3

| Phase | Plans Complete | Status | Completed |
|-------|----------------|--------|-----------|
| 1. Backend Cleanup | 0/3 | Not started | - |
| 2. Frontend Cleanup | 0/2 | Not started | - |
| 3. Data Cleanup | 0/1 | Not started | - |
