---
phase: 01-backend-cleanup
plan: 03
type: execute
wave: 2
depends_on: ["01-01", "01-02"]
files_modified:
  - server/main.go
  - server/initialize/router.go
  - server/initialize/plugin.go
autonomous: true
requirements: [CORE-01, CORE-02, CORE-03, CORE-04, CORE-05, CORE-06, CORE-07, CORE-08]
must_haves:
  truths:
    - "User authentication APIs work correctly (login, logout, getUserInfo)"
    - "Role management APIs work correctly (CRUD, permission association)"
    - "Menu management APIs work correctly (CRUD, button permissions)"
    - "API management APIs work correctly (CRUD, role association)"
    - "JWT + Casbin authentication system works"
    - "Plugin system entry point can mount normally"
    - "Swagger documentation is accessible"
    - "Database migration works correctly"
  artifacts:
    - path: "server/api/v1/system/sys_user.go"
      ensures: "User auth APIs exist"
    - path: "server/api/v1/system/sys_authority.go"
      ensures: "Role management APIs exist"
    - path: "server/api/v1/system/sys_menu.go"
      ensures: "Menu management APIs exist"
    - path: "server/api/v1/system/sys_api.go"
      ensures: "API management APIs exist"
    - path: "server/middleware/jwt.go"
      ensures: "JWT middleware exists"
    - path: "server/middleware/casbin_rbac.go"
      ensures: "Casbin middleware exists"
    - path: "server/utils/plugin/plugin.go"
      ensures: "Plugin interface exists"
  key_links:
    - from: "server/api/v1/system/"
      to: "server/service/system/"
      pattern: "service\\.ServiceGroupApp\\.SystemServiceGroup"
    - from: "server/router/system/"
      to: "server/api/v1/system/"
      pattern: "api\\.ApiGroupApp\\.SystemApiGroup"
    - from: "server/middleware/"
      to: "server/service/"
      pattern: "CasbinHandler"
---

## Plan 01-03: Verify Core Framework Functionality

### Objective
Verify that core framework functionality remains intact after example and plugin deletion.

### Context

This plan verifies the 8 core requirements after Plans 01-01 and 01-02 have removed example modules and plugins:

1. CORE-01: User authentication (login, logout, getUserInfo)
2. CORE-02: Role management (CRUD, permission association)
3. CORE-03: Menu management (CRUD, button permissions)
4. CORE-04: API management (CRUD, role association)
5. CORE-05: JWT + Casbin authentication
6. CORE-06: Plugin system entry
7. CORE-07: Swagger documentation
8. CORE-08: Database migration

### Tasks

#### Task 1: Verify Go code compiles without errors

<read_first>
- server/main.go
- server/initialize/router.go
- server/initialize/plugin.go
</read_first>

<action>
Run Go build to verify all code compiles after example and plugin deletion:

```bash
cd server && go build ./...
```

This verifies:
- No broken import statements in enter.go files
- No dangling references to deleted example modules
- Plugin system still initializes correctly
</action>

<acceptance_criteria>
- go build ./... completes with exit code 0
- No "cannot find package" errors
- No "undefined" errors
</acceptance_criteria>

#### Task 2: Verify core API files exist and are intact

<read_first>
- server/api/v1/system/sys_user.go
- server/api/v1/system/sys_authority.go
- server/api/v1/system/sys_menu.go
- server/api/v1/system/sys_api.go
</read_first>

<action>
Verify the following core API files exist and contain expected endpoints:

1. sys_user.go should contain:
   - Login function (user login)
   - Logout function (user logout)
   - GetUserInfo function (get current user info)
   - ChangePassword function

2. sys_authority.go should contain:
   - CreateAuthority (create role)
   - DeleteAuthority (delete role)
   - UpdateAuthority (update role)
   - GetAuthorityList (list roles)

3. sys_menu.go should contain:
   - CreateMenu (create menu)
   - DeleteMenu (delete menu)
   - UpdateMenu (update menu)
   - GetMenuList (list menus)

4. sys_api.go should contain:
   - CreateApi (create API)
   - DeleteApi (delete API)
   - UpdateApi (update API)
   - GetApiList (list APIs)
   - GetApiByAuthority (get APIs by role)
</action>

<acceptance_criteria>
- All four API files exist
- Each file contains at least 3 of the expected functions listed above
- Files have valid Go syntax
</acceptance_criteria>

#### Task 3: Verify middleware files exist

<read_first>
- server/middleware/jwt.go
- server/middleware/casbin_rbac.go
</read_first>

<action>
Verify JWT and Casbin middleware files exist and contain expected functions:

1. jwt.go should contain:
   - JWTAuth() middleware function
   - ParseToken() function to parse JWT tokens

2. casbin_rbac.go should contain:
   - CasbinHandler() middleware function
   - Casbin() function to enforce policies
</action>

<acceptance_criteria>
- server/middleware/jwt.go exists with JWTAuth and ParseToken functions
- server/middleware/casbin_rbac.go exists with CasbinHandler and Casbin functions
</acceptance_criteria>

#### Task 4: Verify plugin system interface exists

<read_first>
- server/utils/plugin/plugin.go
</read_first>

<action>
Verify the plugin system interface is intact:

1. plugin.go should define the Plugin interface with methods:
   - Register(*gin.RouterGroup) - for registering routes
   - RouterPath() string - returns the base path for plugin routes

2. Verify initialize/plugin.go exists and contains InstallPlugin function
</action>

<acceptance_criteria>
- server/utils/plugin/plugin.go defines Plugin interface
- server/initialize/plugin.go contains InstallPlugin function
- Plugin interface has Register and RouterPath methods
</acceptance_criteria>

#### Task 5: Verify Swagger documentation setup

<read_first>
- server/main.go (look for swagger docs registration)
- server/docs/ directory
</read_first>

<action>
Verify Swagger documentation is properly configured:

1. Check server/docs/ directory exists with swagger files
2. Verify swagger docs are registered in main.go or router initialization
3. Look for _ "github.com/swaggo/swag/cmd/swag" in main.go imports
</action>

<acceptance_criteria>
- server/docs/ directory exists with swagger JSON/YAML files
- Swagger is registered in the application startup
</acceptance_criteria>

#### Task 6: Verify database migration setup

<read_first>
- server/initialize/gorm.go
- server/model/ system models
</read_first>

<action>
Verify database migration mechanism is intact:

1. Check initialize/gorm.go contains AutoMigrate calls
2. Verify system models (sys_user, sys_authority, sys_menu, sys_api) have proper GORM tags
3. Verify RegisterTables function exists in initialize package
</action>

<acceptance_criteria>
- initialize/gorm.go contains db.AutoMigrate() calls
- All system models have gorm.Model embedded or equivalent fields
- RegisterTables function registers all system tables
</acceptance_criteria>

### Verification

Run these commands to verify:

```bash
# 1. Verify code compiles
cd server && go build ./...

# 2. List system API files
ls server/api/v1/system/

# 3. List middleware files
ls server/middleware/

# 4. List plugin utils
ls server/utils/plugin/

# 5. List swagger docs
ls server/docs/
```

### Success Criteria

1. Go code compiles without errors
2. Core API files exist (sys_user.go, sys_authority.go, sys_menu.go, sys_api.go)
3. JWT middleware exists with proper functions
4. Casbin middleware exists with proper functions
5. Plugin interface is defined in server/utils/plugin/plugin.go
6. Swagger docs directory exists
7. Database migration (AutoMigrate) is configured
8. No compilation errors from removed example/plugin modules

### Output

After completion, create `.planning/phases/01-backend-cleanup/01-03-verify-core-framework-SUMMARY.md`
