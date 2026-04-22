# Testing Patterns

**Analysis Date:** 2026-04-21

## Testing Framework

### Backend (Go)

**Testing Library:** `testing` (standard Go library) + `github.com/stretchr/testify`

**Assertion:** `github.com/stretchr/testify/assert` for assertions

**Example from** `server/utils/timer/timed_task_test.go`:
```go
import (
    "fmt"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

func TestNewTimerTask(t *testing.T) {
    tm := NewTimerTask()
    _tm := tm.(*timer)

    {
        _, err := tm.AddTaskByFunc("func", "@every 1s", mockFunc, "测试mockfunc")
        assert.Nil(t, err)
        _, ok := _tm.cronList["func"]
        if !ok {
            t.Error("no find func")
        }
    }
}
```

### Frontend (JavaScript/Vue)

**Testing:** Not visibly adopted in the codebase

**Note:** No test files found in `web/src/` directory. No testing framework configured in `web/package.json`.

## Test File Organization

### Backend Test Files

**Location:** Same directory as the code being tested

**Naming:** `*_test.go` suffix

**Examples:**
- `server/utils/validator_test.go` - Tests for validator utilities
- `server/utils/timer/timed_task_test.go` - Tests for timer functionality
- `server/utils/json_test.go` - Tests for JSON utilities
- `server/utils/ast/ast_test.go` - Tests for AST utilities
- `server/service/system/auto_code_package_test.go` - Tests for auto code package service
- `server/service/system/auto_code_template_test.go` - Tests for auto code template service
- `server/utils/ast/package_initialize_router_test.go` - Tests for package initialization

### Frontend Test Files

**Location:** Not established

**Status:** No test files detected in frontend codebase

## Test Structure

### Backend Test Patterns

**Standard Table-Driven Test:**
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
    }{
        {
            name: "测试 package",
            args: args{
                ctx: context.Background(),
                info: &request.SysAutoCodePackageCreate{
                    Template:    "package",
                    PackageName: "gva",
                },
            },
            wantErr: false,
        },
        // ... more test cases
    }
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

**Direct Test Function:**
```go
func TestVerify(t *testing.T) {
    PageInfoVerify := Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "Name": {NotEmpty()}}
    var testInfo PageInfoTest
    testInfo.Name = "test"
    testInfo.PageInfo.Page = 0
    testInfo.PageInfo.PageSize = 0
    err := Verify(testInfo, PageInfoVerify)
    if err == nil {
        t.Error("校验失败，未能捕捉0值")
    }
    // ... more assertions
}
```

**Mock Pattern:**
```go
type mockJob struct{}

func (job mockJob) Run() {
    mockFunc()
}

func mockFunc() {
    time.Sleep(time.Second)
    fmt.Println("1s...")
}
```

## Running Tests

### Backend

**Run all tests:**
```bash
go test ./...
```

**Run tests with verbose output:**
```bash
go test -v ./server/utils/...
```

**Run specific test file:**
```bash
go test -v ./server/utils/validator_test.go
```

### Frontend

**Status:** No test scripts configured

## Coverage

### Backend

**No explicit coverage enforcement** found in the codebase

**To view coverage:**
```bash
go test -cover ./...
```

### Frontend

**No testing framework configured** - coverage not applicable

## Test Categories Found

### Unit Tests

**Validator Tests** (`server/utils/validator_test.go`):
- Tests for `Verify()` function
- Tests for `NotEmpty()`, `RegexpMatch()`, comparison rules

**Timer Tests** (`server/utils/timer/timed_task_test.go`):
- Tests for task scheduling
- Tests for job and function-based tasks
- Tests for cron finding and clearing

**JSON Tests** (`server/utils/json_test.go`):
- Tests for JSON key extraction

**AST Tests** (`server/utils/ast/ast_test.go`):
- Tests for Go AST parsing and printing

### Service Tests

**Auto Code Package Tests** (`server/service/system/auto_code_package_test.go`):
- Tests for `Create()` method
- Tests for `templates()` method

**Auto Code Template Tests** (`server/service/system/auto_code_template_test.go`):
- Tests for `Create()` method
- Tests for `Preview()` method

### Plugin Tests

**Package Initialize Tests** (`server/utils/ast/package_init_test.go`):
**Plugin Initialize Tests** (`server/utils/ast/plugin_init_test.go`):
**AST Init Tests** (`server/utils/ast/ast_init_test.go`):

These test code generation for plugin infrastructure.

## CI/CD Testing Pipeline

**Status:** Not visible in the codebase

No GitHub Actions workflows or other CI configuration found in the repository.

## Test Utilities

### Test Assertions

Using `github.com/stretchr/testify/assert`:
```go
assert.Nil(t, err)
assert.Equal(t, expected, actual)
assert.True(t, condition)
```

### Test Helpers

**Paginate method** for test database queries (from `server/model/common/request/page.go`):
```go
func (r *PageInfo) Paginate() func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        if r.Page <= 0 {
            r.Page = 1
        }
        // ...
        offset := (r.Page - 1) * r.PageSize
        return db.Offset(offset).Limit(r.PageSize)
    }
}
```

## Testing Conventions Summary

| Aspect | Backend (Go) | Frontend (Vue) |
|--------|-------------|----------------|
| Framework | testing + testify/assert | Not adopted |
| Test Location | Same as source, `*_test.go` | N/A |
| Test Naming | `TestFunctionName` | N/A |
| Pattern | Table-driven tests | N/A |
| Mocking | Manual mock types | N/A |
| Coverage | Not enforced | N/A |

---

*Testing analysis: 2026-04-21*