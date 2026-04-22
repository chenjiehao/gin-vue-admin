# Architecture

**Analysis Date:** 2026-04-21

## Overall Architecture

**Pattern:** Monolithic full-stack application with frontend-backend separation and plugin extensibility

The gin-vue-admin (GVA) framework is organized as two separate deployments:
- **Backend (`server/`)**: Go + Gin REST API server
- **Frontend (`web/`)**: Vue 3 SPA with Vite bundler

Communication between frontend and backend is via HTTP/REST over JSON. The backend also supports an MCP (Model Context Protocol) service for AI tooling integration.

## High-Level Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                         Frontend (Vue 3 SPA)                     │
│  ┌─────────┐   ┌─────────┐   ┌─────────┐   ┌─────────┐          │
│  │  Views  │───│  API    │───│ Pinia   │───│  Router │          │
│  │         │   │ Client  │   │ Stores  │   │         │          │
│  └─────────┘   └────┬────┘   └─────────┘   └─────────┘          │
│                      │                                           │
│                      ▼                                           │
│              @/utils/request.js                                 │
│              (Axios wrapper, JWT handling)                      │
└────────────────────────────┬────────────────────────────────────┘
                             │ HTTP/REST (JSON)
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Backend (Go + Gin)                         │
│  ┌──────────────────────────────────────────────────────────┐   │
│  │                    Middleware Chain                      │   │
│  │  GinRecovery → JWTAuth → CasbinHandler → RouteHandler   │   │
│  └──────────────────────────────────────────────────────────┘   │
│                              │                                   │
│                              ▼                                   │
│  ┌──────────┐     ┌──────────┐     ┌──────────┐                │
│  │   API    │────▶│ Service  │────▶│  Model   │                │
│  │  Layer   │     │  Layer   │     │  Layer   │                │
│  └──────────┘     └──────────┘     └──────────┘                │
│        │                                               ▲        │
│        ▼                                               │        │
│  ┌──────────┐     ┌──────────┐                         │        │
│  │  Router  │     │  Plugin  │─────────────────────────┘        │
│  │  Layer   │     │  System  │                                 │
│  └──────────┘     └──────────┘                                 │
│                                                                  │
│  ┌──────────┐     ┌──────────┐     ┌──────────┐                │
│  │   MCP    │     │  Global  │     │   Core   │                │
│  │  Server  │     │  Vars    │     │ Init     │                │
│  └──────────┘     └──────────┘     └──────────┘                │
└──────────────────────────────────────────────────────────────────┘
                             │
                             ▼
┌──────────────────────────────────────────────────────────────────┐
│                    External Services                              │
│  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌─────────┐            │
│  │   MySQL │  │  Redis  │  │  Mongo  │  │ Cloud   │            │
│  │         │  │         │  │         │  │ Storage │            │
│  └─────────┘  └─────────┘  └─────────┘  └─────────┘            │
└──────────────────────────────────────────────────────────────────┘
```

## Frontend Architecture

**Pattern:** Composition API + Pinia state management + Vue Router + Axios

### Layers and Their Interactions

```
┌─────────────────────────────────────────────────────────┐
│                    Views/Pages                          │
│  (src/view/*) — Business pages, el-drawer for forms    │
│                         │                                │
│                         ▼                                │
│  ┌─────────────────────────────────────────────────┐    │
│  │                 API Layer                         │    │
│  │  (src/api/*.js) — service(url, method, data)    │    │
│  │  Uses @/utils/request.js (Axios wrapper)        │    │
│  └─────────────────────────────────────────────────┘    │
│                         │                                │
│                         ▼                                │
│  ┌─────────────────────────────────────────────────┐    │
│  │              State Management (Pinia)            │    │
│  │  src/pinia/modules/*.js — reactive state        │    │
│  │  useUserStore, useRouterStore, useAppStore       │    │
│  └─────────────────────────────────────────────────┘    │
│                         │                                │
│                         ▼                                │
│  ┌─────────────────────────────────────────────────┐    │
│  │              Router (Vue Router)                 │    │
│  │  src/router/index.js — hash history, lazy load  │    │
│  │  permission.js — route guards, async router     │    │
│  └─────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────┘
```

### Request Flow (Frontend)

1. **View** calls API function from `src/api/*.js`
2. **API function** invokes `service` (axios instance from `@/utils/request.js`)
3. **Request interceptor** attaches `x-token` JWT from `useUserStore().token`
4. **Backend** processes request and returns JSON `{code, data, msg}`
5. **Response interceptor** checks `code === 0`, handles 401 redirect, auto-refreshes JWT if `new-token` header present
6. **API function** returns data to View

### Key Frontend Architecture Decisions

- **Pinia stores expose state and methods**: `useUserStore()` returns `{userInfo, token, LoginIn, LoginOut, ...}`
- **Dynamic router loading**: Routes fetched from backend on login via `useRouterStore().SetAsyncRouter()`, then flattened and registered via `permission.js` router guards
- **API layer co-located by domain**: `src/api/user.js`, `src/api/authority.js`, `src/api/system.js`
- **Component libraries**: Element Plus for UI, UnoCSS for atomic styling, ECharts for charts

## Backend Architecture

**Pattern:** Layered architecture with API → Service → Model chain

### Backend Layers

```
┌─────────────────────────────────────────────────────────────┐
│                       Router Layer                           │
│  (src/router/*) — Route definitions, middleware attachment  │
│  References API handlers via router.RouterGroupApp           │
│  Routes: PublicGroup vs PrivateGroup (JWT + Casbin)         │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        API Layer                             │
│  (src/api/v1/*) — HTTP handlers, param binding, response   │
│  References Services via api.ApiGroupApp.SystemApiGroup     │
│  Example: sys_user.go → userService → response.OkWith...   │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      Service Layer                           │
│  (src/service/*) — Business logic, database CRUD           │
│  NEVER touches gin.Context — pure Go functions              │
│  Returns (data, error) tuples                               │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                       Model Layer                            │
│  (src/model/*) — GORM structs, request/response DTOs        │
│  system/ — core business models (SysUser, SysAuthority)     │
│  common/request|response — shared pagination, response      │
└─────────────────────────────────────────────────────────────┘
```

### Request/Response Flow (Backend)

1. **Request arrives** at Gin route handler
2. **Middleware chain**: GinRecovery (panic logging) → JWTAuth (JWT validation) → CasbinHandler (RBAC check)
3. **Router** maps URL to API handler function in `api/v1/system/`
4. **API handler** (`sys_user.go:Login`):
   - Binds JSON body via `c.ShouldBindJSON(&l)`
   - Validates input via `utils.Verify(l, utils.LoginVerify)`
   - Calls `userService.Login(u)` (Service layer)
   - On error: `response.FailWithMessage(err.Error(), c)`
   - On success: `response.OkWithDetailed(systemRes.LoginResponse{...}, "登录成功", c)`
5. **Service layer** (`service/system/sys_user_service.go:Login`):
   - Performs DB operations via GORM
   - Returns `(user *system.SysUser, err error)`
   - NEVER uses `gin.Context`
6. **Model layer** (`model/system/sys_user.go`):
   - GORM struct defining database schema
   - Inherits `global.GVA_MODEL` for `ID`, `CreatedAt`, `UpdatedAt`
7. **Response** serialized as JSON `{code: 0, data: {...}, msg: "登录成功"}`

### Layer Group Pattern (enter.go)

All layers use `enter.go` to expose grouped instances:

```go
// server/api/enter.go
var ApiGroupApp = new(ApiGroup)
type ApiGroup struct {
    SystemApiGroup  system.ApiGroup
    ExampleApiGroup example.ApiGroup
}

// server/service/enter.go
var ServiceGroupApp = new(ServiceGroup)
type ServiceGroup struct {
    SystemServiceGroup  system.ServiceGroup
    ExampleServiceGroup example.ServiceGroup
}

// server/router/enter.go
var RouterGroupApp = new(RouterGroup)
type RouterGroup struct {
    System  system.RouterGroup
    Example example.RouterGroup
}
```

Cross-layer references:
- Router → API: `api.ApiGroupApp.SystemApiGroup.BaseApi.Login`
- API → Service: `var userService = service.ServiceGroupApp.SystemServiceGroup.UserService`
- Service → Model: direct struct usage

### Middleware Chain

```
[Gin Recovery] → [Logger] → [JWT Auth] → [Casbin RBAC] → [Route Handler]
     │              │             │              │
     ▼              ▼             ▼              ▼
  Panic         Request        Token          Policy
  recovery      logging        validation     enforcement
                                    │
                                    ▼
                             [Extract claims to Context]
```

The `PrivateGroup` in `initialize/router.go` applies `JWTAuth()` and `CasbinHandler()`:
```go
PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
```

### Global State (server/global/global.go)

```go
var (
    GVA_DB        *gorm.DB          // Primary database connection
    GVA_DBList    map[string]*gorm.DB // Multi-database support
    GVA_REDIS     redis.UniversalClient
    GVA_REDISList map[string]redis.UniversalClient
    GVA_MONGO     *qmgo.QmgoClient
    GVA_CONFIG    config.Server      // Viper config
    GVA_LOG       *zap.Logger        // Zap logger
    GVA_VP        *viper.Viper
    BlackCache    local_cache.Cache  // JWT blacklist cache
)
```

## Plugin Architecture

GVA supports two plugin versions:

### Plugin V1 (server/plugin/announcement)

Plugin interface defined in `server/utils/plugin/plugin.go`:
```go
type Plugin interface {
    Register(group *gin.RouterGroup)
    RouterPath() string
}
```

Plugin structure (`server/plugin/announcement/`):
```
announcement/
├── api/          # API handlers
├── config/       # Plugin config
├── initialize/   # DB migration, menu, API registration
│   ├── api.go    # Register plugin APIs
│   ├── gorm.go   # AutoMigrate models
│   ├── menu.go   # Create sidebar menus
│   └── router.go # Register routes
├── model/        # Plugin data models
├── plugin/       # Plugin interface implementation
├── router/       # Route definitions
├── service/      # Business logic
└── plugin.go     # Entry point (calls Register in init())
```

Plugin registration via `server/plugin/register.go`:
```go
_ "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement"
```

Plugin activation in `initialize/plugin.go:InstallPlugin`:
```go
func InstallPlugin(PrivateGroup, PublicGroup *gin.RouterGroup, engine *gin.Engine) {
    bizPluginV1(PrivateGroup, PublicGroup)
    bizPluginV2(engine)
}
```

### Plugin V2 (server/utils/plugin/v2/plugin.go)

Simpler interface:
```go
type Plugin interface {
    Register(group *gin.Engine)
}
```

V2 plugins receive the full Gin engine and handle their own routing groups.

## Data Flow: Typical API Call

### Login Flow (Full Stack)

```
[Frontend] User submits login form
    │
    ▼
src/view/login/index.vue
    │ calls LoginIn(loginInfo)
    ▼
src/pinia/modules/user.js: LoginIn()
    │ calls login(loginInfo) from @/api/user.js
    ▼
src/api/user.js: login()
    │ service({url: '/base/login', method: 'post', data})
    ▼
src/utils/request.js (Axios instance)
    │ interceptors: attach x-token header
    │ POST /base/login JSON
    ▼
[Backend] Gin Router
    │
    ├─ PublicGroup.POST("/base/login", BaseApi.Login)
    │
    ▼
api/v1/system/sys_user.go: Login()
    │ 1. c.ShouldBindJSON(&l) — parse body
    │ 2. utils.Verify(l, LoginVerify) — validate
    │ 3. userService.Login(u) — business logic
    ▼
service/system/sys_user_service.go: Login()
    │ 1. global.GVA_DB.Where("username = ?", u.Username).First(&user)
    │ 2. bcrypt.Compare(user.Password, u.Password)
    │ 3. return &user, nil
    ▼
model/system/sys_user.go: SysUser struct
    │ GORM model with UUID, Username, Password (json:"-"), etc.
    ▼
response.OkWithDetailed(LoginResponse{User, Token, ExpiresAt}, "登录成功", c)
    │
    ▼
[Gin Context] JSON response: {code: 0, data: {user, token, ...}, msg: "登录成功"}
    │
    ▼
[Frontend] Response interceptor
    │
    ├─ If new-token header: userStore.setToken(new-token)
    ├─ If code === 0: return response.data
    └─ Pinia userInfo updated, router navigates to defaultRouter
```

## Key Architectural Decisions

1. **Strict layer isolation**: `Router → API → Service → Model`. API never touches DB directly, Service never uses `gin.Context`

2. **Global instance pattern**: All cross-module references use `*GroupApp` singletons (`ServiceGroupApp`, `ApiGroupApp`, `RouterGroupApp`) to avoid circular dependencies

3. **JWT in headers**: Token passed via `x-token` header, not cookies. `new-token` header for silent refresh

4. **Casbin for RBAC**: Policy-based access control enforced via middleware. Policies stored in DB (`sys_casbin` table)

5. **Response standardization**: All APIs return `{code, data, msg}`. `code=0` is success, non-zero is error

6. **Swagger for API docs**: Every API handler has full Swagger comments parsed by `swag init`. Docs at `/swagger/index.html`

7. **Auto-migration for plugins**: Plugins call `db.AutoMigrate()` in their `initialize/gorm.go` to install DB tables

8. **Config via Viper**: All configuration in `config.yaml`, loaded via `core.Viper()`, stored in `global.GVA_CONFIG`