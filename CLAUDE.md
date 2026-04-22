<!-- GSD:project-start source:PROJECT.md -->
## Project

**gin-vue-admin 精简脚手架**

基于 gin-vue-admin 的**精简框架脚手架**，移除所有业务模块，只保留核心框架能力。后续任何二次开发都基于此骨架进行，避免冗余代码干扰。

**Core Value:** 提供一个**干净的、可扩展的 RBAC 管理系统骨架**，包含用户认证、权限管理、菜单路由、API 管理等基础能力，去除所有业务相关模块。

### Constraints

- **技术约束**: 必须基于现有 GVA 框架结构，不能改变核心架构
- **兼容性约束**: 保留与现有插件系统的兼容性
<!-- GSD:project-end -->

<!-- GSD:stack-start source:codebase/STACK.md -->
## Technology Stack

## Languages
- Go 1.24.0 - Backend server implementation
- Vue 3.5.31 - Frontend UI framework
- JavaScript/TypeScript - Frontend build and tooling
## Runtime
- Go 1.24.0 with go1.24.2 toolchain
- Node.js (npm 11.3.0, core-js 3.38.1)
- npm 11.3.0 (frontend)
- Go modules (backend)
## Backend Frameworks
- Gin 1.10.0 (`github.com/gin-gonic/gin`) - HTTP web framework
- GORM 1.25.12 (`gorm.io/gorm`) - Object-relational mapper
- GORM Gen 0.3.26 (`gorm.io/gen`) - Code generation for GORM
- MySQL: `gorm.io/driver/mysql` 1.5.7
- PostgreSQL: `gorm.io/driver/postgres` 1.5.11
- SQL Server: `gorm.io/driver/sqlserver` 1.5.4
- SQLite: `github.com/glebarez/sqlite` 1.11.0
- Oracle: `github.com/dzwvip/gorm-oracle` 0.1.2
- MongoDB: `go.mongodb.org/mongo-driver` 1.17.2 + `github.com/qiniu/qmgo` 1.1.9
- JWT: `github.com/golang-jwt/jwt/v5` 5.2.2
- Casbin 2.103.0 (`github.com/casbin/casbin/v2`) - RBAC permission framework
- Casbin GORM Adapter 3.32.0 (`github.com/casbin/gorm-adapter/v3`)
- Viper 1.19.0 (`github.com/spf13/viper`) - Config management from YAML
- Zap 1.27.0 (`go.uber.org/zap`) - Structured logging
- gookit/color 1.5.4 - Terminal colors
- Redis: `github.com/redis/go-redis/v9` 9.7.0
- Swagger: `github.com/swaggo/swag` 1.16.4
- Gin Swagger: `github.com/swaggo/gin-swagger` 1.6.0
- Swagger Files: `github.com/swaggo/files` 1.0.1
- UUID: `github.com/google/uuid` 1.6.0
- Excel: `github.com/xuri/excelize/v2` 2.9.0
- Cron: `github.com/robfig/cron/v3` 3.0.1
- Captcha: `github.com/mojocn/base64Captcha` 1.3.8
- Email: `github.com/jordan-wright/email` 4.0.1-0.20210109023952
- automaxprocs 1.6.0 (`go.uber.org/automaxprocs`) - CPU limit optimization
## Frontend Frameworks
- Vue 3.5.31 - Progressive JavaScript framework
- Vue Router 4.4.3 (`vue-router`) - Client-side routing
- Pinia 2.2.2 - State management
- Vite 6.2.3 (`vite`) - Next-gen frontend build tool
- @vitejs/plugin-vue 5.0.3 - Vue 3 SFC plugin for Vite
- Element Plus 2.13.6 (`element-plus`) - Vue 3 UI library
- @element-plus/icons-vue 2.3.1 - Element Plus icons
- UnoCSS 66.5.0 (`@unocss/vite`, `@unocss/preset-wind3`, `@unocss/extractor-svelte`, `@unocss/transformer-directives`) - Atomic CSS framework
- Sass 1.78.0 - CSS preprocessor
- Axios 1.8.2 - Promise-based HTTP client
- ECharts 5.5.1 (`echarts`) - Interactive charting library
- vue-echarts 7.0.3 - ECharts wrapper for Vue
- @vueuse/core 11.0.3 - Vue composition utilities
- @vueuse/integrations 12.0.0
- mitt 3.0.1 - Event emitter
- nprogress 0.2.0 - Progress bar
- qs 6.13.0 - Query string parsing
- sortablejs 1.15.3 - Drag and drop
- spark-md5 3.0.2 - MD5 implementation
- universal-cookie 7 - Cookie management
- @form-create/designer 3.2.6 - Form designer
- @form-create/element-ui 3.2.10 - Form generation
- vform3-builds 3.0.10 - Form builder
- @wangeditor/editor 5.1.23 - Rich text editor
- @wangeditor/editor-for-vue 5.1.12 - Vue bindings
- ace-builds 1.36.4 - Code editor
- vue3-ace-editor 2.2.4 - Ace editor for Vue 3
- @vue-office/docx 1.6.2 - Word document viewer
- @vue-office/excel 1.7.11 - Excel document viewer
- @vue-office/pdf 2.0.2 - PDF document viewer
- vue-cropper 1.1.4 - Image cropping
- marked 14.1.1 - Markdown parser
- marked-highlight 2.1.4 - Markdown with highlight.js
- vue-qr 4.0.9 - QR code generator
- vite-plugin-vue-devtools 7.0.16 - Vue DevTools integration
- vite-plugin-banner 0.8.0 - Add banner to output
- vite-plugin-importer 0.2.5 - Import resolution
- vite-auto-import-svg 2.5.0 - Auto import SVG
- vite-check-multiple-dom 0.2.1 - Multi-domain checking
- @vitejs/plugin-legacy 6.0.0 - Legacy browser support
- ESLint 8.57.0 with eslint-plugin-vue 9.19.2
- @babel/eslint-parser 7.25.1
## Cloud Storage Integrations
- `minio` - MinIO configuration
- `aliyun-oss` - Aliyun OSS configuration
- `aws-s3` - AWS S3 / R2 configuration
- `tencent-cos` - Tencent COS configuration
- `qiniu` - Qiniu configuration
- `hua-wei-obs` - Huawei OBS configuration
- `cloudflare-r2` - Cloudflare R2 configuration
## Deployment
- Base image: `centos:7`
- Location: `deploy/docker/Dockerfile`
- Entrypoint: `deploy/docker/entrypoint.sh`
- Configuration: `deploy/docker-compose/docker-compose.yaml`
- Includes MySQL, Redis, Nginx, and application services
- Server manifests: `deploy/kubernetes/server/`
- Web manifests: `deploy/kubernetes/web/`
## Additional Backend Components
- `github.com/mholt/archives` 0.1.1 - Archive extraction
- `github.com/otiai10/copy` 1.14.1 - File copy utility
- `github.com/shirou/gopsutil/v3` 3.24.5 - System/utility library
- `golang.org/x/crypto` 0.37.0 - Cryptographic functions
- `github.com/unrolled/secure` 1.17.0 - Security headers middleware
- `github.com/songzhibin97/gkit` 1.2.13 - Go utility toolkit
<!-- GSD:stack-end -->

<!-- GSD:conventions-start source:CONVENTIONS.md -->
## Conventions

## Naming Conventions
### Backend (Go)
- Source files: `snake_case.go` (e.g., `sys_user.go`, `auto_code_service.go`)
- Test files: `*_test.go` suffix (e.g., `validator_test.go`)
- Plugin directories: `snake_case` (e.g., `announcement`, `email`)
- PascalCase: `SysUser`, `PageInfo`, `LoginResponse`
- Interface names: PascalCase with "er" suffix where applicable: `Login`, `Service`
- PascalCase for exported: `GetUserList`, `ChangePassword`, `Register`
- camelCase for unexported: `tokenNext`, `getUserInfo`
- camelCase: `userName`, `pageInfo`, `httpStatus`
- Constants: UPPER_SNAKE_CASE for true constants: `DEFAULT_REQUEST_TIMEOUT`
- Package-level vars in api/service: `var userService = service.ServiceGroupApp.SystemServiceGroup.UserService`
- snake_case in gorm tags and JSON: `json:"user_name"`, `column:"user_name"`
- PascalCase in Go struct field names: `UserName string`
### Frontend (JavaScript/Vue)
- JavaScript files: `camelCase.js` (e.g., `user.js`, `sysDictionary.js`)
- Vue components: `PascalCase.vue` or `kebab-case.vue`
- API files: `camelCase.js` matching backend module
- PascalCase for component names in templates: `<UserDialog>`, `<CustomPic>`
- kebab-case for file names: `customPic.vue`, `warningBar.vue`
- camelCase: `tableData`, `searchInfo`, `pageInfo`
- Constants: UPPER_SNAKE_CASE: `DEFAULT_REQUEST_TIMEOUT`
- camelCase store names: `useUserStore`, `useAppStore`
## Code Style
### Backend Formatting
### Frontend Formatting
## Import Organization
### Backend Go Imports
### Frontend Imports
- `@/` maps to `web/src/`
- Use absolute paths, never relative
## Error Handling
### Backend Patterns
- Return `error` on failure
- Return `nil` on success
- Never handle `gin.Context` in service layer
- Check error returned from service
- Use response helpers: `response.FailWithMessage()`, `response.FailWithDetailed()`
- Log errors with zap: `global.GVA_LOG.Error("操作失败!", zap.Error(err))`
- Use `utils.Verify()` with Rules map
- Validation rules defined in `server/utils/validator.go`
- Common patterns: `NotEmpty()`, `Lt()`, `Le()`, `Ge()`, `Gt()`, `RegexpMatch()`
- `SUCCESS = 0` in `server/model/common/response/response.go`
- `ERROR = 7` for failures
- HTTP status: 200 OK for business errors, 401 for unauthorized, etc.
### Frontend Patterns
- API functions return response objects with `code` property
- Check `res.code === 0` for success
- Use `ElMessage` for user-facing error display
### Route Name Convention (CRITICAL)
**组件的 `defineOptions.name`、数据库 `sys_base_menus.name`、`router.push({ name: '...' })` 三者的值必须完全一致，且统一使用 kebab-case（小写+中划线）。**

示例：
- 数据库菜单 name: `dataSourceForm`
- 组件定义: `defineOptions({ name: 'dataSourceForm' })`
- 路由跳转: `router.push({ name: 'dataSourceForm' })`
- pathInfo.json 映射: `"/src/view/dataIntegration/dataSourceForm.vue": "dataSourceForm"`

**常见错误：混用大小写**（如 `DataSourceForm` vs `dataSourceForm`），导致 "No match for" 路由错误。
## Logging
### Backend (Zap)
- `Error` for failures that need attention
- `Info` for significant operations
- `Debug` for development details
### Frontend
## Swagger Annotations
### Required for Every API Function
- `@Tags` - API group name (usually module name like "SysUser", "Base")
- `@Summary` - Brief description of what the endpoint does
- `@Security` - Authentication method (ApiKeyAuth for most, optional for public endpoints)
- `@accept` - Request content type (application/json)
- `@Produce` - Response content type
- `@Param` - Parameters (data body, path params, query params)
- `@Success` - Response structure
- `@Router` - HTTP method and path
## Layer Responsibilities
### Service Layer (`server/service/`)
- Implement business logic
- Perform database CRUD operations via GORM
- Return `(result, error)` tuples
- Validate input data types
- Handle HTTP context (gin.Context)
- Write response directly
- Access request body directly
### API Layer (`server/api/v1/`)
- Parse request parameters (ShouldBindJSON, ShouldBindQuery)
- Validate input using utils.Verify
- Call appropriate service methods
- Return formatted responses using response helpers
- Log errors with global.GVA_LOG
- Write business logic directly
- Perform database operations
- Handle middleware logic
### Router Layer (`server/router/`)
- Define routes and HTTP methods
- Apply middleware (OperationRecord, Auth middleware)
- Group routes logically
## Request/Response Model Conventions
### Backend Request Models
- Embed `common.PageInfo` for pagination
- Use `json` and `form` tags for binding
- Include `OrderKey` and `Desc` for sorting
### Backend Response Models
### Frontend API Functions
## Code Organization
### Global Instance Pattern
### API Service References
### Plugin Structure
## Git Commit Message Format
- `feat` - New feature
- `fix` - Bug fix
- `refactor` - Code refactoring
- `Update` - Documentation or README changes
- `发布` - Release/version changes
## Testing Patterns
### Backend Go Tests
### Frontend Testing

## Development Guidelines

### Service Restart Protocol
**重启服务前必须先停止旧进程。**

使用 `go run .` 或编译后的可执行文件启动后端服务时，每次运行都会创建一个新进程。如果不先停止旧进程，会导致：
- 多个实例同时运行，端口冲突
- 旧代码的进程占用端口，新代码无法生效
- 调试时请求打到旧服务上，返回 404 或旧逻辑

**正确流程：**
1. 停止旧服务：`lsof -ti :<端口号> | xargs kill`
2. 启动新服务：`go run .` 或 `./可执行文件`

**前端同理**：重启 Vite 开发服务器前，先用 `lsof -ti :<端口号> | xargs kill` 停止旧进程。

示例（后端端口 8888，前端端口 8080）：
```bash
# 停止后端
lsof -ti :8888 | xargs kill 2>/dev/null
# 停止前端
lsof -ti :8080 | xargs kill 2>/dev/null
# 重启后端
cd server && go run .
# 重启前端
cd web && npm run dev
```

### Casbin 权限配置
**新增 API 接口后必须同步配置 Casbin 权限。**

gin-vue-admin 使用 Casbin 做 RBAC 权限控制，所有私有 API（需要登录访问的）都会经过 Casbin 中间件验证。如果接口的权限策略未配置，即使登录用户也会返回"权限不足"或"未登录"错误。

**常见错误表现：**
- 接口返回 `code: 7, msg: "未登录或非法访问，请登录"`
- 即使已登录且 Token 有效，仍然提示权限不足

**权限配置方式：**
1. **方式一：通过数据初始化脚本**（推荐）
   在 `server/source/system/casbin.go` 中添加新的权限策略：
   ```go
   {
       Path:       "/dataSource/testConnection",
       Method:     "POST",
       RoleID:     "1",  // 1 通常是超级管理员角色
   },
   ```

2. **方式二：通过代码生成器自动录入**
   使用 GVA 代码生成器生成模块时，会自动录入 API 权限到 `sys_api` 表，并通过 `sys_menu` 关联到角色。

3. **方式三：手动在数据库中添加**
   在 `sys_menu` 和 `sys_authority_menu` 表中添加对应菜单和角色关联。

**新增模块权限检查清单：**
- [ ] API 路由已注册
- [ ] API 信息已录入 `sys_api` 表（或通过 source 文件初始化）
- [ ] 菜单已录入 `sys_menu` 表
- [ ] 角色（通常是 admin）已绑定对应菜单
- [ ] 测试接口访问是否正常

**调试权限问题：**
```bash
# 检查服务是否正常启动
curl http://localhost:8888/health

# 检查 Token 是否有效
curl -H "Authorization: Bearer <你的Token>" http://localhost:8888/dataSource/getDataSourceList

# 如果返回 "未登录或非法访问"，则是 Casbin 权限未配置
```
<!-- GSD:conventions-end -->

<!-- GSD:architecture-start source:ARCHITECTURE.md -->
## Architecture

## Overall Architecture
- **Backend (`server/`)**: Go + Gin REST API server
- **Frontend (`web/`)**: Vue 3 SPA with Vite bundler
## High-Level Architecture Diagram
```
```
## Frontend Architecture
### Layers and Their Interactions
```
```
### Request Flow (Frontend)
### Key Frontend Architecture Decisions
- **Pinia stores expose state and methods**: `useUserStore()` returns `{userInfo, token, LoginIn, LoginOut, ...}`
- **Dynamic router loading**: Routes fetched from backend on login via `useRouterStore().SetAsyncRouter()`, then flattened and registered via `permission.js` router guards
- **API layer co-located by domain**: `src/api/user.js`, `src/api/authority.js`, `src/api/system.js`
- **Component libraries**: Element Plus for UI, UnoCSS for atomic styling, ECharts for charts
## Backend Architecture
### Backend Layers
```
```
### Request/Response Flow (Backend)
### Layer Group Pattern (enter.go)
```go
```
- Router → API: `api.ApiGroupApp.SystemApiGroup.BaseApi.Login`
- API → Service: `var userService = service.ServiceGroupApp.SystemServiceGroup.UserService`
- Service → Model: direct struct usage
### Middleware Chain
```
```
```go
```
### Global State (server/global/global.go)
```go
```
## Plugin Architecture
### Plugin V1 (server/plugin/announcement)
```go
```
```
```
```go
```
```go
```
### Plugin V2 (server/utils/plugin/v2/plugin.go)
```go
```
## Data Flow: Typical API Call
### Login Flow (Full Stack)
```
```
## Key Architectural Decisions
<!-- GSD:architecture-end -->

<!-- GSD:workflow-start source:GSD defaults -->
## GSD Workflow Enforcement

Before using Edit, Write, or other file-changing tools, start work through a GSD command so planning artifacts and execution context stay in sync.

Use these entry points:
- `/gsd:quick` for small fixes, doc updates, and ad-hoc tasks
- `/gsd:debug` for investigation and bug fixing
- `/gsd:execute-phase` for planned phase work

Do not make direct repo edits outside a GSD workflow unless the user explicitly asks to bypass it.
<!-- GSD:workflow-end -->



<!-- GSD:profile-start -->
## Developer Profile

> Profile not yet configured. Run `/gsd:profile-user` to generate your developer profile.
> This section is managed by `generate-claude-profile` -- do not edit manually.
<!-- GSD:profile-end -->
