# Codebase Concerns

**Analysis Date:** 2026-04-21

## Tech Debt

### SQL Injection Risk in Auto Code History
- **Issue:** Table name is directly concatenated into DROP TABLE statement without sanitization
- **File:** `server/service/system/auto_code_history.go:213-215`
- **Code:**
```go
return global.MustGetGlobalDBByDBName(BusinessDb).Exec("DROP TABLE " + tableName).Error
return global.GVA_DB.Exec("DROP TABLE " + tableName).Error
```
- **Impact:** If `tableName` comes from user input, SQL injection is possible
- **Fix approach:** Validate table name against a whitelist of allowed characters or use parameterized queries for object names

### Unchecked Error in Tencent COS Upload
- **Issue:** Uses `panic(err)` for upload failures instead of returning error
- **File:** `server/utils/upload/tencent_cos.go:33`
- **Code:**
```go
_, err := client.Object.Put(context.Background(), global.GVA_CONFIG.TencentCOS.PathPrefix+"/"+fileKey, f, nil)
if err != nil {
    panic(err)
}
```
- **Impact:** Crashes the request goroutine instead of gracefully handling the error
- **Fix approach:** Return error to caller instead of panicking

### In-Memory Token Cache Not Horizontally Scalable
- **Issue:** Export tokens stored in package-level in-memory maps
- **File:** `server/api/v1/system/sys_export_template.go:22-26`
- **Code:**
```go
var (
    exportTokenCache      = make(map[string]interface{})
    exportTokenExpiration = make(map[string]time.Time)
    tokenMutex            sync.RWMutex
)
```
- **Impact:** In multi-instance deployments, tokens created on one instance cannot be validated on another. Additionally, the `cleanupExpiredTokens` goroutine started in `init()` never terminates gracefully.
- **Fix approach:** Use Redis for token storage instead of in-memory maps

### Goroutine Leak in Export Template
- **Issue:** `init()` function spawns an unbounded cleanup goroutine with no shutdown mechanism
- **File:** `server/api/v1/system/sys_export_template.go:44-46`
- **Code:**
```go
func init() {
    go cleanupExpiredTokens()
}
```
- **Impact:** Goroutine runs for entire application lifetime with no way to gracefully shut down
- **Fix approach:** Use a context with cancel, or implement shutdown hook

### Raw SQL Query String Concatenation
- **Issue:** Multiple places use string formatting to build SQL queries
- **Files:**
  - `server/service/system/auto_code_mysql.go:17` - `SELECT SCHEMA_NAME AS database FROM INFORMATION_SCHEMA.SCHEMATA;`
  - `server/service/system/auto_code_mysql.go:31` - `select table_name as table_name from information_schema.tables where table_schema = ?`
  - `server/service/system/auto_code_mysql.go:46-75` - Large raw SQL query with `?` placeholders
  - `server/service/system/sys_export_template.go:607` - Raw SQL without placeholders
- **Impact:** While some use parameterized queries, raw SQL is harder to audit and maintain
- **Fix approach:** Consider using GORM's query builder for better type safety

## Known Bugs

### No graceful shutdown for background cleanup
- **Symptoms:** Application hangs on shutdown because `cleanupExpiredTokens` goroutine never exits
- **File:** `server/api/v1/system/sys_export_template.go:29-41`
- **Trigger:** Sending SIGTERM to application
- **Workaround:** None - requires code change

## Security Considerations

### JWT Signing Key from Configuration
- **Risk:** JWT signing key is loaded from config file (`global.GVA_CONFIG.JWT.SigningKey`)
- **File:** `server/utils/jwt.go:28`
- **Current mitigation:** Key should be strong and kept secret
- **Recommendations:** Consider using environment variables or secrets manager for production

### Incomplete Error Message Information Leakage
- **Risk:** Error messages may reveal internal system details
- **File:** `server/middleware/jwt.go:41`
- **Code:** `response.NoAuth(err.Error())` - exposes internal JWT error details
- **Recommendations:** Return generic error messages to clients while logging detailed errors internally

### File Upload Path Traversal Potential
- **Risk:** File keys are generated with just timestamp prefix
- **File:** `server/utils/upload/tencent_cos.go:29`
- **Code:** `fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)`
- **Recommendations:** Sanitize `file.Filename` to prevent path traversal attacks (e.g., `../` sequences)

## Performance Bottlenecks

### Large File Sizes
- **Files with excessive lines:**
  - `server/service/system/sys_skills.go` - 785 lines
  - `server/mcp/gva_execute.go` - 750 lines
  - `server/service/system/auto_code_package.go` - 743 lines
  - `server/service/system/sys_export_template.go` - 724 lines
  - `server/utils/autocode/template_funcs.go` - 713 lines
  - `server/api/v1/system/sys_user.go` - 516 lines
- **Impact:** Harder to maintain, test, and understand
- **Improvement path:** Break into smaller, focused files following single responsibility principle

### N+1 Query Patterns with Multiple Preloads
- **Files:** Multiple service files use chained `Preload()` calls
- **Example:** `server/service/system/sys_user.go:53`
- **Code:** `.Preload("Authorities").Preload("Authority")`
- **Impact:** Each preload may trigger additional queries; nested preloads compound this
- **Improvement path:** Use `Preload()` with condition callbacks to limit data, or use JOINs

### Context.Background() Usage in Production Paths
- **Issue:** Uses `context.TODO()` or `context.Background()` where a proper context should be passed
- **Files:**
  - `server/mcp/client/client.go:23` - `ctx := context.Background()`
  - `server/middleware/limit_ip.go:66-88` - Multiple Redis calls with `context.Background()`
  - `server/utils/upload/aws_s3.go:42,68,78` - S3 operations with `context.TODO()`
- **Impact:** Operations cannot be cancelled or traced properly
- **Improvement path:** Accept context as parameter from the request

## Fragile Areas

### Auto Code Generation Template System
- **Files:** `server/utils/autocode/template_funcs.go` (713 lines)
- **Why fragile:** Complex string template processing, difficult to debug when generated code is incorrect
- **Safe modification:** Test all template variations before changes
- **Test coverage:** Likely minimal - auto-generated code is tested indirectly

### MCP Integration
- **Files:** `server/mcp/` directory with multiple files
- **Why fragile:** External MCP protocol integration, error handling across network boundaries
- **Safe modification:** Mock the MCP client in tests
- **Test coverage:** Some test files exist (`*_test.go`) but integration testing is limited

### Database Migration System
- **Files:** `server/initialize/gorm*.go`
- **Why fragile:** Auto-migrate runs on startup, any model mismatch can cause data loss
- **Safe modification:** Always backup database before running migrations, test on staging first
- **Test coverage:** Manual testing required

## Scaling Limits

### In-Memory Session/Token Storage
- **Current capacity:** Limited by single instance memory
- **Limit:** Cannot scale horizontally without external session store
- **Scaling path:** Move to Redis for session/JWT blacklist storage

### Database Connection Pool
- **Configured in:** `server/config/gorm_*.go`
- **Default:** MaxIdleConns=10, MaxOpenConns=100
- **Limit:** May need tuning for high-traffic deployments
- **Scaling path:** Adjust pool size based on load testing

## Dependencies at Risk

### External MCP Library
- **Package:** `github.com/mark3labs/mcp-go`
- **Version:** v0.41.1
- **Risk:** Third-party MCP protocol implementation, may have breaking changes
- **Impact:** Changes in this library could break MCP functionality
- **Migration plan:** Pin to specific version, monitor for security advisories

### Multiple Database Driver Support
- **Drivers:** MySQL, PostgreSQL, SQLite, SQLServer, Oracle, MongoDB
- **Risk:** Each driver has its own quirks and potential vulnerabilities
- **Impact:** Complexity in testing all database combinations
- **Migration plan:** Consider limiting supported databases in production

## Missing Critical Features

### No Request Timeout Middleware
- **Problem:** Only `middleware/timeout.go` exists but it uses panic/recover
- **Missing:** Proper context deadline handling at API layer

### No Rate Limiting on Sensitive Endpoints
- **Problem:** IP rate limiting exists but login attempts are only partially protected
- **Missing:** Better account lockout mechanisms after failed attempts

## Test Coverage Gaps

### Untested Core Business Logic
- **Files:** Most service layer files have no corresponding `*_test.go`
- **What's not tested:** Business logic in services like `sys_authority.go`, `sys_user.go`, `sys_menu.go`
- **Risk:** Refactoring could break functionality without detection
- **Priority:** High - add unit tests for critical paths

### No Integration Tests
- **What's not tested:** Full request lifecycle from API to database
- **Risk:** API and service layers may work in isolation but fail in integration
- **Priority:** Medium - add integration test suite

### Auto-Code Generated Files Not Tested
- **What's not tested:** Code generated by auto-code template system
- **Risk:** Template bugs could generate broken code for all users
- **Priority:** Low - would require testing generated output

---

*Concerns audit: 2026-04-21*
