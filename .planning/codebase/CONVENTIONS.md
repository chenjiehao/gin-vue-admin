# Coding Conventions

**Analysis Date:** 2026-04-21

## Naming Conventions

### Backend (Go)

**Files:**
- Source files: `snake_case.go` (e.g., `sys_user.go`, `auto_code_service.go`)
- Test files: `*_test.go` suffix (e.g., `validator_test.go`)
- Plugin directories: `snake_case` (e.g., `announcement`, `email`)

**Types/Structs:**
- PascalCase: `SysUser`, `PageInfo`, `LoginResponse`
- Interface names: PascalCase with "er" suffix where applicable: `Login`, `Service`

**Functions/Methods:**
- PascalCase for exported: `GetUserList`, `ChangePassword`, `Register`
- camelCase for unexported: `tokenNext`, `getUserInfo`

**Variables:**
- camelCase: `userName`, `pageInfo`, `httpStatus`
- Constants: UPPER_SNAKE_CASE for true constants: `DEFAULT_REQUEST_TIMEOUT`
- Package-level vars in api/service: `var userService = service.ServiceGroupApp.SystemServiceGroup.UserService`

**Database/JSON Fields:**
- snake_case in gorm tags and JSON: `json:"user_name"`, `column:"user_name"`
- PascalCase in Go struct field names: `UserName string`

### Frontend (JavaScript/Vue)

**Files:**
- JavaScript files: `camelCase.js` (e.g., `user.js`, `sysDictionary.js`)
- Vue components: `PascalCase.vue` or `kebab-case.vue`
- API files: `camelCase.js` matching backend module

**Components:**
- PascalCase for component names in templates: `<UserDialog>`, `<CustomPic>`
- kebab-case for file names: `customPic.vue`, `warningBar.vue`

**Variables:**
- camelCase: `tableData`, `searchInfo`, `pageInfo`
- Constants: UPPER_SNAKE_CASE: `DEFAULT_REQUEST_TIMEOUT`

**State (Pinia):**
- camelCase store names: `useUserStore`, `useAppStore`

## Code Style

### Backend Formatting

**Tool:** gofmt (standard Go formatter)

**Line Length:** No strict limit, but typically under 120 characters

**Import Organization:**
```go
import (
    "standard-library"
    "external-packages"
    "internal/project-packages"
)
```
Group order:
1. Standard library (fmt, time, context)
2. External packages (gin, gorm, zap, jwt)
3. Internal packages (github.com/flipped-aurora/gin-vue-admin/server/...)

**Example from** `server/api/v1/system/sys_user.go`:
```go
import (
    "strconv"
    "time"

    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/system"
    systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
    systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"
    "go.uber.org/zap"
)
```

### Frontend Formatting

**Tool:** ESLint + Prettier (via eslint.config.mjs)

**ESLint Configuration** (from `web/eslint.config.mjs`):
```javascript
{
  name: 'app/files-to-lint',
  files: ['**/*.{js,mjs,jsx,vue}'],
  rules: {
    'vue/max-attributes-per-line': 0,
    'vue/no-v-model-argument': 0,
    'vue/multi-word-component-names': 'off',
    'no-lone-blocks': 'off',
    'no-extend-native': 'off',
    'no-unused-vars': ['error', { argsIgnorePattern: '^_' }]
  }
}
```

**UnoCSS:** Using presetWind3 with transformerDirectives for Tailwind-like utilities

## Import Organization

### Backend Go Imports

**Order:**
1. Standard library (fmt, os, time, context, errors)
2. Third-party packages (gin, gorm, zap, jwt, redis)
3. Internal packages (github.com/flipped-aurora/gin-vue-admin/server/...)

**Alias pattern:** `systemReq` for request models, `systemRes` for response models

### Frontend Imports

**Order:**
1. Vue/core imports (vue, reactiveness APIs)
2. External libraries (element-plus, axios, pinia)
3. Internal modules (@/api, @/utils, @/components, @/pinia)

**Path aliases:**
- `@/` maps to `web/src/`
- Use absolute paths, never relative

## Error Handling

### Backend Patterns

**Service Layer:**
- Return `error` on failure
- Return `nil` on success
- Never handle `gin.Context` in service layer

**API Layer:**
- Check error returned from service
- Use response helpers: `response.FailWithMessage()`, `response.FailWithDetailed()`
- Log errors with zap: `global.GVA_LOG.Error("操作失败!", zap.Error(err))`

**Example from** `server/api/v1/system/sys_user.go`:
```go
func (b *BaseApi) GetUserList(c *gin.Context) {
    var pageInfo systemReq.GetUserList
    err := c.ShouldBindJSON(&pageInfo)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    err = utils.Verify(pageInfo, utils.PageInfoVerify)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    list, total, err := userService.GetUserInfoList(pageInfo)
    if err != nil {
        global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}
```

**Validation:**
- Use `utils.Verify()` with Rules map
- Validation rules defined in `server/utils/validator.go`
- Common patterns: `NotEmpty()`, `Lt()`, `Le()`, `Ge()`, `Gt()`, `RegexpMatch()`

**Error Codes:**
- `SUCCESS = 0` in `server/model/common/response/response.go`
- `ERROR = 7` for failures
- HTTP status: 200 OK for business errors, 401 for unauthorized, etc.

### Frontend Patterns

**Error Handling:**
- API functions return response objects with `code` property
- Check `res.code === 0` for success
- Use `ElMessage` for user-facing error display

**Example from** `web/src/view/superAdmin/user/user.vue`:
```javascript
const res = await deleteUser({ id: row.ID })
if (res.code === 0) {
    ElMessage.success('删除成功')
    await getTableData()
} else {
    ElMessage.error(res.msg || '操作失败')
}
```

## Logging

### Backend (Zap)

**Logger:** `global.GVA_LOG` of type `*zap.Logger`

**Patterns:**
```go
global.GVA_LOG.Error("登录失败! 用户名不存在或者密码错误!", zap.Error(err))
global.GVA_LOG.Info("Some info message")
global.GVA_LOG.Debug("Debug details")
```

**Log Levels:**
- `Error` for failures that need attention
- `Info` for significant operations
- `Debug` for development details

### Frontend

**Console:** Standard `console.log()`, `console.error()`, `console.warn()`

**Pattern in store:**
```javascript
console.error('LoginIn error:', error)
```

## Swagger Annotations

### Required for Every API Function

**Location:** Directly above API handler function in `server/api/v1/`

**Example from** `server/api/v1/system/sys_user.go`:
```go
// GetUserList
// @Tags      SysUser
// @Summary   分页获取用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.GetUserList                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router    /user/getUserList [post]
func (b *BaseApi) GetUserList(c *gin.Context) {
    // ...
}
```

**Annotation Fields:**
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

**DO:**
- Implement business logic
- Perform database CRUD operations via GORM
- Return `(result, error)` tuples
- Validate input data types

**DON'T:**
- Handle HTTP context (gin.Context)
- Write response directly
- Access request body directly

**Example:** `server/service/system/sys_user_service.go`

### API Layer (`server/api/v1/`)

**DO:**
- Parse request parameters (ShouldBindJSON, ShouldBindQuery)
- Validate input using utils.Verify
- Call appropriate service methods
- Return formatted responses using response helpers
- Log errors with global.GVA_LOG

**DON'T:**
- Write business logic directly
- Perform database operations
- Handle middleware logic

### Router Layer (`server/router/`)

**DO:**
- Define routes and HTTP methods
- Apply middleware (OperationRecord, Auth middleware)
- Group routes logically

**Example from** `server/router/system/sys_user.go`:
```go
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
    userRouter := Router.Group("user").Use(middleware.OperationRecord())
    userRouterWithoutRecord := Router.Group("user")
    {
        userRouter.POST("admin_register", baseApi.Register)
        userRouter.POST("changePassword", baseApi.ChangePassword)
        // ...
    }
    {
        userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList)
        // ...
    }
}
```

## Request/Response Model Conventions

### Backend Request Models

**Location:** `server/model/system/request/` or `server/model/common/request/`

**Naming:** `XxxRequest`, `XxxSearch`, `GetXxxList`

**Pattern:**
```go
type GetUserList struct {
    common.PageInfo
    Username string `json:"username" form:"username"`
    NickName string `json:"nickName" form:"nickName"`
    Phone    string `json:"phone" form:"phone"`
    Email    string `json:"email" form:"email"`
    OrderKey string `json:"orderKey" form:"orderKey"`
    Desc     bool   `json:"desc" form:"desc"`
}
```

**Common Patterns:**
- Embed `common.PageInfo` for pagination
- Use `json` and `form` tags for binding
- Include `OrderKey` and `Desc` for sorting

### Backend Response Models

**Location:** `server/model/system/response/` or `server/model/common/response/`

**Standard Response Structure** (in `server/model/common/response/response.go`):
```go
type Response struct {
    Code int         `json:"code"`
    Data interface{} `json:"data"`
    Msg  string      `json:"msg"`
}
```

**Page Result:**
```go
type PageResult struct {
    List     interface{} `json:"list"`
    Total    int64       `json:"total"`
    Page     int         `json:"page"`
    PageSize int         `json:"pageSize"`
}
```

### Frontend API Functions

**Pattern:**
```javascript
import service from '@/utils/request'

// @Tags User
// @Summary 获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页参数"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
export const getUserList = (data) => {
  return service({
    url: '/user/getUserList',
    method: 'post',
    data: data
  })
}
```

## Code Organization

### Global Instance Pattern

**Service Group** (`server/service/enter.go`):
```go
var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
    SystemServiceGroup  system.ServiceGroup
    ExampleServiceGroup example.ServiceGroup
}
```

**API Group** (`server/api/v1/enter.go`):
```go
var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
    SystemApiGroup  system.ApiGroup
    ExampleApiGroup example.ApiGroup
}
```

**Router Group** (`server/router/enter.go`):
```go
var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
    System  system.RouterGroup
    Example example.RouterGroup
}
```

### API Service References

**In API layer** (`server/api/v1/system/enter.go`):
```go
var (
    userService = service.ServiceGroupApp.SystemServiceGroup.UserService
    menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
    // ... other services
)
```

### Plugin Structure

**Plugin Entry** (`server/plugin/announcement/plugin.go`):
```go
var _ interfaces.Plugin = (*plugin)(nil)
var Plugin = new(plugin)

type plugin struct{}

func init() {
    interfaces.Register(Plugin)
}

func (p *plugin) Register(group *gin.Engine) {
    // Initialize API, Menu, Dictionary, Gorm, Router
}
```

## Git Commit Message Format

**From recent commits:**
```
feat: 添加获取 MCP 路由列表的接口
fix: 修复获取x-token的方式，确保正确获取用户token
refactor: 调整ai工具rules
Update README.md
```

**Format:** `type: description`

**Types:**
- `feat` - New feature
- `fix` - Bug fix
- `refactor` - Code refactoring
- `Update` - Documentation or README changes
- `发布` - Release/version changes

## Testing Patterns

### Backend Go Tests

**Test Framework:** `testing` package with `testify/assert`

**File Location:** Same package as implementation, `*_test.go` suffix

**Example from** `server/utils/validator_test.go`:
```go
func TestVerify(t *testing.T) {
    PageInfoVerify := Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "Name": {NotEmpty()}}
    var testInfo PageInfoTest
    testInfo.Name = "test"
    // ...
    err := Verify(testInfo, PageInfoVerify)
    if err == nil {
        t.Error("校验失败，未能捕捉0值")
    }
}
```

**Example from** `server/service/system/auto_code_package_test.go`:
```go
func Test_autoCodePackage_Create(t *testing.T) {
    type args struct {
        ctx  context.Context
        info *request.SysAutoCodePackageCreate
    }
    tests := []struct {
        name    string
        args    args
        wantErr bool
    }{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            a := &autoCodePackage{}
            if err := a.Create(tt.args.ctx, tt.args.info); (err != nil) != tt.wantErr {
                t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

### Frontend Testing

**Not widely adopted** - No visible test files in `web/src/`

---

*Convention analysis: 2026-04-21*