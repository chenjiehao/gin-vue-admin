# Requirements: gin-vue-admin 精简脚手架

**Defined:** 2026-04-21
**Core Value:** 提供一个干净的、可扩展的 RBAC 管理系统骨架

## v1 Requirements

### 框架清理

- [ ] **CLEAN-01**: 删除 `server/model/example/` 目录及所有示例模型
- [ ] **CLEAN-02**: 删除 `server/api/v1/example/` 目录及所有示例 API
- [ ] **CLEAN-03**: 删除 `server/service/example/` 目录及所有示例服务
- [ ] **CLEAN-04**: 删除 `server/router/example/` 目录及所有示例路由
- [ ] **CLEAN-05**: 删除 `web/src/view/example/` 目录及所有示例页面
- [ ] **CLEAN-06**: 删除 `web/src/api/example.js` 示例 API 文件
- [ ] **CLEAN-07**: 删除 `server/plugin/announcement/` 插件（如有）
- [ ] **CLEAN-08**: 删除 `server/plugin/email/` 插件（如有）

### 后端保留

- [ ] **CORE-01**: 保留用户认证（登录、登出、获取用户信息）
- [ ] **CORE-02**: 保留角色管理（CRUD、权限关联）
- [ ] **CORE-03**: 保留菜单管理（CRUD、按钮权限）
- [ ] **CORE-04**: 保留 API 管理（CRUD、角色关联）
- [ ] **CORE-05**: 保留 JWT + Casbin 鉴权体系
- [ ] **CORE-06**: 保留插件系统入口和注册机制
- [ ] **CORE-07**: 保留 Swagger 文档注释
- [ ] **CORE-08**: 保留数据库初始化和迁移机制

### 前端保留

- [ ] **FRONT-01**: 保留登录页面
- [ ] **FRONT-02**: 保留主布局（侧边栏、顶栏、内容区）
- [ ] **FRONT-03**: 保留框架级别 API（用户、角色、菜单、API）
- [ ] **FRONT-04**: 保留动态路由加载机制
- [ ] **FRONT-05**: 移除所有业务示例页面

### 数据清理

- [ ] **DATA-01**: 清理初始化数据中的示例菜单
- [ ] **DATA-02**: 清理初始化数据中的示例角色
- [ ] **DATA-03**: 清理初始化数据中的示例用户（保留超级管理员）
- [ ] **DATA-04**: 保留超级管理员默认账号（admin/123456）

## v2 Requirements

（无）

## Out of Scope

| Feature | Reason |
|---------|--------|
| 代码生成器 | 属于高级工具，非最小框架必需 |
| 多语言支持 | i18n 属于业务层需求 |
| 主题定制 | 基础样式足够，后续业务自行扩展 |
| 云存储集成 | 对象存储不是框架核心 |
| 消息通知 | 邮件/短信属于业务功能 |

## Traceability

| Requirement | Phase | Status |
|-------------|-------|--------|
| CLEAN-01 ~ CLEAN-08 | Phase 1 | Pending |
| CORE-01 ~ CORE-08 | Phase 1 | Pending |
| FRONT-01 ~ FRONT-05 | Phase 2 | Pending |
| DATA-01 ~ DATA-04 | Phase 2 | Pending |

**Coverage:**
- v1 requirements: 22 total
- Mapped to phases: 22
- Unmapped: 0 ✓

---
*Requirements defined: 2026-04-21*
*Last updated: 2026-04-21 after initial definition*
