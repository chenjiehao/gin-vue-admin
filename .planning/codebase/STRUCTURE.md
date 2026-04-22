# Codebase Structure

**Analysis Date:** 2026-04-21

## Project Root

```
gin-vue-admin/
├── server/           # Go backend (Gin + GORM)
├── web/              # Vue 3 frontend (Vite + Element Plus)
├── deploy/           # Docker & Kubernetes deployment configs
├── docs/             # Swagger documentation
├── Makefile          # Build automation
├── README.md         # Documentation
└── LICENSE           # MIT License
```

## Server Directory (`server/`)

```
server/
├── main.go                    # Application entry point
├── config.yaml                # Default configuration file
├── api/                       # API Controller Layer
│   └── v1/
│       ├── enter.go           # API group registry
│       │                       # Exposes: ApiGroupApp = new(ApiGroup)
│       ├── system/            # System module API handlers
│       │   ├── sys_user.go    # User login, register, CRUD
│       │   ├── sys_authority.go # Role/authority management
│       │   ├── sys_menu.go    # Menu management
│       │   ├── sys_api.go     # API endpoint management
│       │   └── ...
│       └── example/           # Example module API handlers
├── service/                   # Service Layer (Business Logic)
│   ├── enter.go               # Service group registry
│   │                           # Exposes: ServiceGroupApp = new(ServiceGroup)
│   ├── system/                # System business logic
│   │   ├── sys_user_service.go
│   │   ├── sys_authority_service.go
│   │   └── ...
│   └── example/              # Example business logic
├── router/                    # Route Definitions
│   ├── enter.go               # Router group registry
│   │                           # Exposes: RouterGroupApp = new(RouterGroup)
│   ├── system/                # System routes
│   │   ├── sys_user_router.go
│   │   ├── sys_authority_router.go
│   │   └── ...
│   └── example/              # Example routes
├── model/                      # Data Models
│   ├── system/                # System domain models
│   │   ├── sys_user.go       # User entity (GORM struct)
│   │   ├── sys_authority.go  # Role entity
│   │   ├── sys_menu.go       # Menu entity
│   │   ├── sys_api.go        # API endpoint entity
│   │   ├── request/          # Request DTOs
│   │   │   ├── sys_user.go   # GetUserList, ChangePassword, etc.
│   │   │   └── ...
│   │   └── response/         # Response DTOs
│   │       ├── sys_user.go   # SysUserResponse, LoginResponse
│   │       └── ...
│   ├── example/              # Example models
│   ├── common/               # Shared models
│   │   ├── request/          # PageInfo, GetById, etc.
│   │   └── response/         # PageResult, Response wrappers
│   └── response/             # (duplicate of common/response?)
├── initialize/                # Application Initialization
│   ├── router.go             # Route registration (core/Routers())
│   ├── plugin.go             # Plugin installation (InstallPlugin)
│   ├── gorm.go               # Database connection setup
│   ├── redis.go              # Redis connection
│   ├── mongo.go              # MongoDB connection
│   └── ...
├── middleware/                # HTTP Middleware
│   ├── jwt.go                # JWT authentication (JWTAuth())
│   ├── casbin_rbac.go        # RBAC policy enforcement (CasbinHandler())
│   ├── cors.go               # CORS configuration
│   ├── operation.go          # Operation logging
│   ├── logger.go             # Request logging
│   └── ...
├── core/                      # Core Infrastructure
│   ├── server.go             # Server startup (RunServer())
│   ├── server_run.go         # HTTP server with graceful shutdown
│   ├── viper.go              # Viper config loading
│   ├── zap.go                # Zap logger setup
│   └── internal/             # Internal utilities
├── global/                    # Global Variables
│   └── global.go             # GVA_DB, GVA_REDIS, GVA_CONFIG, etc.
├── config/                    # Configuration Structures
│   └── config.go             # Server config struct (Viper binding)
├── utils/                     # Utility Packages
│   ├── plugin/               # Plugin interface definitions
│   │   ├── plugin.go         # V1 Plugin interface
│   │   └── v2/plugin.go      # V2 Plugin interface
│   ├── captcha/             # Captcha generation
│   ├── upload/              # File upload utilities
│   ├── request/              # HTTP client (outbound requests)
│   └── ...
├── plugin/                    # Bundled Plugins
│   ├── announcement/         # Announcement plugin (V1 plugin example)
│   │   ├── api/              # Plugin API handlers
│   │   ├── config/          # Plugin configuration
│   │   ├── initialize/      # Plugin init (Gorm, Menu, Router)
│   │   ├── model/           # Plugin models
│   │   ├── router/          # Plugin route definitions
│   │   ├── service/         # Plugin business logic
│   │   ├── plugin.go        # Plugin entry point
│   │   └── plugin/          # Plugin interface impl (implements Plugin)
│   └── email/               # Email plugin
├── resource/                  # Code Generation Templates
│   ├── function/            # AutoCode template functions
│   ├── package/             # Template packages for auto-generation
│   │   ├── server/          # Server template (api, model, router, service)
│   │   └── web/             # Web template (api, view)
│   └── plugin/              # Plugin template packages
├── source/                    # Data Initialization (Seed Data)
│   ├── system/              # System module seed data
│   └── example/            # Example seed data
├── cmd/                       # CLI Commands
│   └── mcp/                 # MCP server command
├── docs/                      # Swagger generated docs
├── middleware/               # (duplicate at root level)
└── task/                     # Background task utilities
```

## Web/Src Directory (`web/src/`)

```
web/src/
├── main.js                   # Vue app entry point
├── App.vue                   # Root component
├── api/                       # API Client Layer
│   ├── user.js               # User API (login, register, getUserInfo)
│   ├── authority.js          # Role/authority API
│   ├── menu.js               # Menu API
│   ├── autoCode.js          # AutoCode API
│   ├── system.js            # System config API
│   ├── plugin/              # Plugin API (plugin ры)
│   └── system/              # System APIs by domain
├── assets/                   # Static Assets
│   ├── icons/               # Custom SVG icons
│   └── images/              # Images (logos, placeholders)
├── components/               # Reusable Vue Components
│   ├── application/          # Application-level components
│   ├── charts/              # ECharts wrapper components
│   ├── upload/              # File upload components
│   ├── selectImage/         # Image picker component
│   ├── selectFile/          # File picker component
│   ├── richtext/            # Rich text editor
│   ├── svgIcon/            # SVG icon component
│   ├── exportExcel/         # Excel export component
│   └── ...
├── core/                     # Core Setup & Bootstrap
│   ├── gin-vue-admin.js     # App initialization logic
│   ├── gin-vue-admin/       # GVA specific setup
│   │   ├── index.js         # Setup entrance
│   │   └── ...
│   └── error-handel.js      # Global error handler
├── directive/                # Vue Directives
│   ├── auth/                # Permission directive (v-auth)
│   └── clickOutSide/        # Click outside directive
├── hooks/                    # Composition API Hooks
│   └── (hook utilities)
├── pinia/                    # Pinia State Management
│   ├── index.js             # Pinia store factory
│   │                       # Exports: store, useAppStore, useUserStore, useDictionaryStore
│   └── modules/             # Store modules
│       ├── user.js          # User state (token, userInfo, login/logout)
│       ├── router.js        # Router state (asyncRouters, SetAsyncRouter)
│       ├── app.js          # App config state (global_size, theme)
│       └── dictionary.js    # Dictionary/cache state
├── plugin/                   # Frontend Plugin Modules
│   ├── announcement/        # Announcement plugin frontend
│   │   ├── api/             # Plugin API calls
│   │   ├── view/           # Plugin page components
│   │   └── form/           # Plugin form components
│   └── email/              # Email plugin frontend
├── router/                   # Vue Router Configuration
│   └── index.js             # Router definition (hash history)
├── style/                    # Global Styles
│   ├── element/             # Element Plus theme overrides
│   └── element_visiable.scss # Element visibility styles
├── utils/                    # Utility Functions
│   ├── request.js          # Axios HTTP client (centralized)
│   │                       # Handles: JWT injection, loading, error handling, 401 redirect
│   ├── format.js           # Data formatters (boolean, date, dict lookup)
│   ├── date.js             # Date formatting utilities
│   ├── dictionary.js       # Dictionary data fetcher (getDict)
│   ├── stringFun.js        # String utilities (camel/snake case)
│   ├── bus.js              # Event bus (mitt-based emitter)
│   ├── downloadImg.js      # Image download utility
│   ├── image.js            # Image compression (ImageCompress)
│   ├── asyncRouter.js      # Dynamic router handling
│   ├── btnAuth.js          # Button permission hook (useBtnAuth)
│   └── ...
└── view/                     # Page Components
    ├── login/               # Login page
    ├── init/                # System initialization page
    ├── dashboard/           # Dashboard/home page
    ├── layout/              # Main layout components
    │   ├── index.vue        # Layout wrapper
    │   ├── aside/           # Sidebar navigation
    │   ├── header/          # Top header bar
    │   ├── tabs/            # Tab navigation (multiple tabs)
    │   ├── setting/         # Settings panel drawer
    │   └── screenfull/      # Fullscreen button
    ├── superAdmin/          # Super admin pages
    │   ├── user/            # User management
    │   ├── authority/       # Role/authority management
    │   ├── menu/            # Menu management
    │   ├── api/             # API endpoint management
    │   └── dictionary/      # Dictionary management
    ├── systemTools/         # System tool pages
    │   ├── autoCode/       # AutoCode generator UI
    │   ├── autoPkg/        # Auto package generator
    │   ├── formCreate/     # Form builder
    │   ├── system/          # System configuration
    │   ├── sysError/        # Error log viewer
    │   └── ...
    ├── example/             # Example pages
    │   ├── customer/        # Customer management example
    │   ├── fileUploadAndDownload/ # File upload example
    │   └── breakpoint/      # Breakpoint upload example
    └── person/             # Personal center page
```

## Deploy Directory (`deploy/`)

```
deploy/
├── docker/                  # Docker configuration
│   ├── docker-compose.yml   # Local development compose
│   └── Dockerfile           # Multi-stage build
├── docker-compose/          # Production compose files
│   └── ...
└── kubernetes/              # K8s deployment manifests
    └── ...
```

## Key Entry Points

### Backend: `server/main.go`
```go
func main() {
    initializeSystem()      // Load config → Init DB → Register tables
    core.RunServer()         // Setup Redis → Mongo → Load sys data → Start HTTP server
}
```

Initialization sequence:
1. `core.Viper()` — Load `config.yaml` via Viper into `global.GVA_CONFIG`
2. `initialize.OtherInit()` — Setup system utilities
3. `core.Zap()` — Initialize Zap logger into `global.GVA_LOG`
4. `initialize.Gorm()` — Connect to MySQL/PostgreSQL via GORM
5. `initialize.Timer()` — Setup background scheduler
6. `initialize.DBList()` — Setup additional databases
7. `initialize.RegisterTables()` — AutoMigrate all model tables
8. `initialize.Routers()` — Register all API routes and middleware

### Frontend: `web/src/main.js`
```javascript
const app = createApp(App)
app.use(run)          // GVA initialization plugin
app.use(ElementPlus)   // UI component library
app.use(store)        // Pinia state management
app.use(router)       // Vue Router
app.mount('#app')
```

### Frontend Router: `web/src/router/index.js`
```javascript
const router = createRouter({
  history: createWebHashHistory(),  // Uses hash to avoid server config for SPA
  routes: [
    { path: '/', redirect: '/login' },
    { path: '/login', name: 'Login', component: ... },
    { path: '/init', name: 'Init', component: ... },
    { path: '/:catchAll(.*)', component: ErrorPage }
  ]
})
```

### Permission Guard: `web/src/permission.js`
- Handles white-list routes (Login, Init)
- On login: fetches async router from backend, flattens nested structure, registers to Vue Router
- On each navigation: checks JWT token, handles 401 redirect

## Directory Purposes Summary

| Directory | Purpose |
|-----------|---------|
| `server/api/` | HTTP request handlers — parse params, call services, return responses |
| `server/service/` | Business logic — database CRUD, validation, pure Go (no gin.Context) |
| `server/router/` | URL-to-handler mapping + middleware attachment |
| `server/model/` | GORM structs, request DTOs, response DTOs |
| `server/initialize/` | App startup sequence — config, DB, routes, plugins |
| `server/middleware/` | HTTP pipeline — JWT, Casbin RBAC, CORS, logging |
| `server/core/` | Server bootstrap — Viper, Zap, HTTP server |
| `server/global/` | Shared global variables — DB, Redis, Config references |
| `server/plugin/` | Self-contained feature plugins (announcement, email) |
| `server/resource/` | Code generation templates for auto-code feature |
| `web/src/api/` | Frontend API call wrappers using axios |
| `web/src/pinia/` | Frontend reactive state (user, router, app config) |
| `web/src/view/` | Page-level Vue components |
| `web/src/components/` | Reusable Vue components |
| `web/src/utils/` | Utility functions — HTTP client, formatters, helpers |
| `web/src/router/` | Vue Router definition and route guards |

## Where to Add New Code

### New Backend API (System Module)

| Layer | Path | Action |
|-------|------|--------|
| Model | `server/model/system/request/` | Add request DTO struct |
| Model | `server/model/system/response/` | Add response DTO struct |
| Model | `server/model/system/` | Add GORM entity if new table needed |
| Service | `server/service/system/` | Add `xxx_service.go` with business logic |
| Service | `server/service/enter.go` | Add new service to `ServiceGroup` struct |
| API | `server/api/v1/system/` | Add `xxx_api.go` with HTTP handlers |
| API | `server/api/enter.go` | Add new API to `ApiGroup` struct |
| Router | `server/router/system/` | Add `xxx_router.go` with route definitions |
| Router | `server/router/enter.go` | Add new router to `RouterGroup` struct |
| Swagger | Add comments to API handler functions | Regenerate docs via `swag init` |

### New Plugin

| Part | Path | Description |
|------|------|-------------|
| Backend Entry | `server/plugin/<name>/plugin.go` | Implement Plugin interface |
| Backend API | `server/plugin/<name>/api/` | API handlers |
| Backend Service | `server/plugin/<name>/service/` | Business logic |
| Backend Model | `server/plugin/<name>/model/` | GORM models |
| Backend Init | `server/plugin/<name>/initialize/` | `gorm.go`, `menu.go`, `router.go` |
| Frontend API | `web/src/plugin/<name>/api/` | API calls |
| Frontend View | `web/src/plugin/<name>/view/` | Vue components |
| Registration | `server/plugin/register.go` | Add `_ "path/to/plugin"` |

### New Frontend Page

| Part | Path |
|------|------|
| Page Component | `web/src/view/<module>/page.vue` |
| API Calls | `web/src/api/<module>.js` |
| Pinia Store | `web/src/pinia/modules/<module>.js` (if needed) |
| Router | Dynamically loaded from backend via `SetAsyncRouter()` |

---

*Structure analysis: 2026-04-21*