---
phase: 01-backend-cleanup
plan: 01
type: execute
wave: 1
depends_on: []
files_modified:
  - server/api/v1/enter.go
  - server/service/enter.go
  - server/router/enter.go
autonomous: true
requirements: [CLEAN-01, CLEAN-02, CLEAN-03, CLEAN-04]
must_haves:
  truths:
    - "server/model/example/ directory does not exist"
    - "server/api/v1/example/ directory does not exist"
    - "server/service/example/ directory does not exist"
    - "server/router/example/ directory does not exist"
    - "enter.go files no longer reference example packages"
  artifacts:
    - path: "server/model/example/"
      ensures: "Example models removed"
    - path: "server/api/v1/example/"
      ensures: "Example API handlers removed"
    - path: "server/service/example/"
      ensures: "Example services removed"
    - path: "server/router/example/"
      ensures: "Example routes removed"
  key_links:
    - from: "server/api/v1/enter.go"
      to: "example package"
      pattern: "example\\.ApiGroup"
    - from: "server/service/enter.go"
      to: "example package"
      pattern: "example\\.ServiceGroup"
    - from: "server/router/enter.go"
      to: "example package"
      pattern: "example\\.RouterGroup"
---

## Plan 01-01: Delete Example Module Directories

### Objective
Delete all server/ example module directories and update enter.go files to remove example package references.

### Context

@server/api/v1/enter.go
```go
var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
    SystemApiGroup  system.ApiGroup
    ExampleApiGroup example.ApiGroup  // REMOVE THIS LINE
}
```

@server/service/enter.go
```go
type ServiceGroup struct {
    SystemServiceGroup  system.ServiceGroup
    ExampleServiceGroup example.ServiceGroup  // REMOVE THIS LINE
}
```

@server/router/enter.go
```go
type RouterGroup struct {
    System  system.RouterGroup
    Example example.RouterGroup  // REMOVE THIS LINE
}
```

### Tasks

#### Task 1: Remove example imports and fields from enter.go files

<read_first>
- server/api/v1/enter.go
- server/service/enter.go
- server/router/enter.go
</read_first>

<action>
Update the three enter.go files to remove example package references:

1. server/api/v1/enter.go:
   - Remove import: `"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"`
   - Remove field: `ExampleApiGroup example.ApiGroup` from ApiGroup struct

2. server/service/enter.go:
   - Remove import: `"github.com/flipped-aurora/gin-vue-admin/server/service/example"`
   - Remove field: `ExampleServiceGroup example.ServiceGroup` from ServiceGroup struct

3. server/router/enter.go:
   - Remove import: `"github.com/flipped-aurora/gin-vue-admin/server/router/example"`
   - Remove field: `Example example.RouterGroup` from RouterGroup struct
</action>

<acceptance_criteria>
- server/api/v1/enter.go contains only system package imports and SystemApiGroup field
- server/service/enter.go contains only system package imports and SystemServiceGroup field
- server/router/enter.go contains only system package imports and System field
- No compilation errors from removed imports
</acceptance_criteria>

#### Task 2: Delete server/model/example/ directory

<read_first>
- server/model/enter.go (to verify no cross-module dependencies)
- List all files in server/model/example/ to confirm contents
</read_first>

<action>
Delete the entire server/model/example/ directory and all its contents:
- server/model/example/enter.go
- server/model/example/exa.go
- server/model/example/exa_customer.go
- server/model/example/exa_file_upload_download.go
- server/model/example/exa_breakpoint_continue.go
- server/model/example/exa_attachment_category.go
- server/model/example/request/ (entire directory)
- server/model/example/response/ (entire directory)
</action>

<acceptance_criteria>
- server/model/example/ directory does not exist
- ls server/model/ shows only: common/ system/ enter.go
</acceptance_criteria>

#### Task 3: Delete server/api/v1/example/ directory

<read_first>
- server/api/v1/example/enter.go (to verify service references before deletion)
</read_first>

<action>
Delete the entire server/api/v1/example/ directory and all its contents:
- server/api/v1/example/enter.go
- server/api/v1/example/ (all .go files in this directory)
- server/api/v1/example/ (all subdirectories)
</action>

<acceptance_criteria>
- server/api/v1/example/ directory does not exist
- ls server/api/v1/ shows only: enter.go system/
</acceptance_criteria>

#### Task 4: Delete server/service/example/ directory

<read_first>
- server/service/example/enter.go (to verify structure)
</read_first>

<action>
Delete the entire server/service/example/ directory and all its contents:
- server/service/example/enter.go
- server/service/example/ (all .go files in this directory)
- server/service/example/ (all subdirectories)
</action>

<acceptance_criteria>
- server/service/example/ directory does not exist
- ls server/service/ shows only: enter.go system/
</acceptance_criteria>

#### Task 5: Delete server/router/example/ directory

<read_first>
- server/router/example/enter.go (to verify API references)
</read_first>

<action>
Delete the entire server/router/example/ directory and all its contents:
- server/router/example/enter.go
- server/router/example/ (all .go files in this directory)
- server/router/example/ (all subdirectories)
</action>

<acceptance_criteria>
- server/router/example/ directory does not exist
- ls server/router/ shows only: enter.go system/
</acceptance_criteria>

### Verification

Run these commands to verify deletion:
```bash
# Verify directories removed
ls server/model/      # Should show: common/ system/ enter.go
ls server/api/v1/     # Should show: enter.go system/
ls server/service/    # Should show: enter.go system/
ls server/router/     # Should show: enter.go system/

# Verify Go files compile (no broken imports)
cd server && go build ./...
```

### Success Criteria

1. server/model/example/ does not exist
2. server/api/v1/example/ does not exist
3. server/service/example/ does not exist
4. server/router/example/ does not exist
5. server/api/v1/enter.go, server/service/enter.go, server/router/enter.go no longer reference example packages
6. Go code compiles without errors

### Output

After completion, create `.planning/phases/01-backend-cleanup/01-01-delete-example-dirs-SUMMARY.md`
