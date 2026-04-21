# gin-vue-admin 精简脚手架

## What This Is

基于 gin-vue-admin 的**精简框架脚手架**，移除所有业务模块，只保留核心框架能力。后续任何二次开发都基于此骨架进行，避免冗余代码干扰。

## Core Value

提供一个**干净的、可扩展的 RBAC 管理系统骨架**，包含用户认证、权限管理、菜单路由、API 管理等基础能力，去除所有业务相关模块。

## Requirements

### Validated

基于代码库分析，以下是 gin-vue-admin 框架已有的核心能力（将被保留）：

- ✓ **用户管理** — 用户 CRUD、登录登出、密码修改
- ✓ **角色权限管理** — RBAC 权限模型、角色与权限关联
- ✓ **菜单管理** — 动态菜单、按钮级别权限
- ✓ **API 管理** — API 路由注册、角色与 API 关联
- ✓ **JWT 认证** — Token 生成、刷新、验证
- ✓ **Casbin 鉴权** — 请求级别的权限校验
- ✓ **数据库初始化** — GORM 自动迁移
- ✓ **插件系统** — 插件注册与挂载机制
- ✓ **Swagger 文档** — API 自动生成文档
- ✓ **统一响应格式** — 标准化 JSON 响应结构
- ✓ **请求日志** — 操作记录中间件
- ✓ **CORS 配置** — 跨域请求支持

### Active

- [ ] **CORE-01**: 清理所有示例模块（example 模块）
- [ ] **CORE-02**: 清理所有业务相关菜单和权限数据
- [ ] **CORE-03**: 保留最小用户表结构（仅保留登录必需字段）
- [ ] **CORE-04**: 保留完整的权限框架（Casbin + 菜单 + API）
- [ ] **CORE-05**: 保留插件系统入口
- [ ] **CORE-06**: 前端移除所有业务页面，只保留登录页和框架布局

### Out of Scope

- [业务模块] — 业务功能不属于框架本身
- [多数据库支持] — 框架默认只保留 MySQL 支持
- [云存储集成] — OSS/S3 等存储不是框架必需
- [消息通知] — 邮件/短信等功能不属于核心框架
- [代码生成器] — 属于高级功能，非最小框架必需

## Context

**技术栈：**
- 后端：Go 1.24 + Gin 1.10 + GORM 1.25 + Casbin 2.103 + JWT 5.2
- 前端：Vue 3.5 + Vite 6.2 + Pinia 2.2 + Element Plus 2.10 + UnoCSS
- 部署：Docker 支持

**框架特性：**
- 前后端分离架构
- 严格的分层设计（Router → API → Service → Model）
- 插件化设计，支持独立插件扩展
- 完整的 Swagger API 文档

## Constraints

- **技术约束**: 必须基于现有 GVA 框架结构，不能改变核心架构
- **兼容性约束**: 保留与现有插件系统的兼容性

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| 只保留最小框架 | 脚手架应该干净，去掉业务干扰 | — Pending |
| 保留完整权限体系 | 二次开发都需要 RBAC 能力 | — Pending |
| 前端只留登录和框架 | 最小可用 UI 骨架 | — Pending |

---

## Evolution

This document evolves at phase transitions and milestone boundaries.

**After each phase transition** (via `/gsd:transition`):
1. Requirements invalidated? → Move to Out of Scope with reason
2. Requirements validated? → Move to Validated with phase reference
3. New requirements emerged? → Add to Active
4. Decisions to log? → Add to Key Decisions
5. "What This Is" still accurate? → Update if drifted

**After each milestone** (via `/gsd:complete-milestone`):
1. Full review of all sections
2. Core Value check — still the right priority?
3. Audit Out of Scope — reasons still valid?
4. Update Context with current state

---
*Last updated: 2026-04-21 after initialization*
