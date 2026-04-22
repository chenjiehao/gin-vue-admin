# External Integrations

**Analysis Date:** 2026-04-21

## Cloud Storage (OSS)

**Provider Selection:** The OSS provider is configured via `server/config.yaml` under `system.oss-type`. The factory function `NewOss()` in `server/utils/upload/upload.go` instantiates the appropriate driver.

**Supported Providers:**

| Provider | SDK | Config Section | Auth Fields |
|----------|-----|----------------|-------------|
| Local | Built-in | N/A | N/A |
| MinIO | `github.com/minio/minio-go/v7` | `minio` | `endpoint`, `access-key-id`, `access-key-secret`, `bucket-name`, `use-ssl` |
| Aliyun OSS | `github.com/aliyun/aliyun-oss-go-sdk` | `aliyun-oss` | `endpoint`, `access-key-id`, `access-key-secret`, `bucket-name`, `bucket-url` |
| AWS S3 | `github.com/aws/aws-sdk-go-v2/service/s3` | `aws-s3` | `secret-id`, `secret-key`, `bucket`, `region` |
| Tencent COS | `github.com/tencentyun/cos-go-sdk-v5` | `tencent-cos` | `secret-id`, `secret-key`, `bucket`, `region` |
| Qiniu | `github.com/qiniu/go-sdk/v7` | `qiniu` | `access-key`, `secret-key`, `bucket` |
| Huawei OBS | `github.com/huaweicloud/huaweicloud-sdk-go-obs` | `hua-wei-obs` | `access-key`, `secret-key`, `endpoint`, `bucket` |
| Cloudflare R2 | AWS S3 SDK (S3-compatible) | `cloudflare-r2` | `access-key-id`, `secret-access-key`, `account-id` |

**Upload Interface:** `server/utils/upload/upload.go` defines the `OSS` interface with methods:
- `UploadFile(file *multipart.FileHeader) (string, string, error)` - Returns URL, key, error
- `DeleteFile(key string) error`

**MinIO Special Handling:** MinIO requires client initialization with connection verification. If initialization fails, the server panics to prevent unsafe operation.

## Database Engines

**Supported via GORM drivers:**

| Database | Driver | Config Section |
|----------|--------|----------------|
| MySQL | `gorm.io/driver/mysql` | `mysql` |
| PostgreSQL | `gorm.io/driver/postgres` | `pgsql` |
| SQL Server | `gorm.io/driver/sqlserver` | `mssql` |
| Oracle | `github.com/dzwvip/gorm-oracle` | `oracle` |
| SQLite | `github.com/glebarez/sqlite` | `sqlite` |
| MongoDB | `go.mongodb.org/mongo-driver` | `mongo` |

**Multi-Database Support:** `server/config.yaml` defines `db-list` for additional databases. The `server/initialize/db_list.go` handles initialization of multiple database connections.

**MongoDB Integration:**
- Uses `github.com/qiniu/qmgo` 1.1.9 as MongoDB ODM
- Configuration in `config.yaml` under `mongo` section:
  - `coll`, `database`, `username`, `password`
  - `hosts` array with `host` and `port`
  - `options`, `auth-source`, `min-pool-size`, `max-pool-size`
  - `socket-timeout-ms`, `connect-timeout-ms`

**Database Selection:** Configured via `system.db-type` in `server/config.yaml`.

## Caching

**Redis Integration:**
- Client: `github.com/redis/go-redis/v9` 9.7.0
- Configuration: `server/config.yaml` under `redis` and `redis-list` sections

**Redis Configuration Options:**
- `useCluster` - Enable cluster mode
- `addr` - Single instance address
- `password` - Authentication
- `db` - Database number
- `clusterAddrs` - Cluster node addresses

**Redis Usage:**
- JWT blacklisting (`server/service/system/jwt_black_list.go`)
- Rate limiting/IP blocking (`server/middleware/limit_ip.go`)
- Captcha storage (`server/utils/captcha/redis.go`)
- General caching via `redis-list` configuration

## MCP Server Integration

**Implementation:** The project includes a Model Context Protocol (MCP) server integration for AI-assisted development.

**MCP Package:** `github.com/mark3labs/mcp-go` v0.41.1

**MCP Configuration:** `server/config/mcp.go`
```go
type MCP struct {
    Name            string  // Server name
    Version         string  // Protocol version
    Path            string  // HTTP path (default: "/mcp")
    Addr            int     // Port number
    BaseURL         string  // Base URL for requests
    UpstreamBaseURL string  // Upstream server URL
    AuthHeader      string  // Authentication header
    RequestTimeout  int     // Request timeout in seconds
}
```

**Config in `server/config.yaml`:**
```yaml
mcp:
    name: GVA_MCP
    version: v1.0.0
    addr: 8889
    separate: false
```

**MCP Server Implementation:**
- `server/mcp/server.go` - `NewMCPServer()` and `NewStreamableHTTPServer()` functions
- Uses `mark3labs/mcp-go` library for protocol handling
- Registers all MCP tools via `RegisterAllTools(s)`
- Supports both standard MCP server and Streamable HTTP server mode

**MCP Tools Registered:** Located in `server/mcp/`:
- `api_lister.go` - List API endpoints
- `api_creator.go` - Create new API endpoints
- `menu_lister.go` - List menu items
- `menu_creator.go` - Create menu items
- `dictionary_query.go` - Query dictionary data
- `dictionary_generator.go` - Generate dictionary entries
- `gva_analyze.go` - GVA code analysis
- `gva_execute.go` - Execute GVA operations
- `gva_review.go` - Code review functionality
- `autocode_http.go` - Auto-code HTTP operations
- `requirement_analyzer.go` - Analyze requirements
- `page.go` - Page generation

**MCP CLI:** `server/cmd/mcp/main.go` - Standalone MCP server executable with `config.yaml` at `server/cmd/mcp/config.yaml`

**Frontend MCP Integration:**
- `web/src/view/systemTools/autoCode/mcp.vue` - MCP configuration UI
- `web/src/view/systemTools/autoCode/mcpTest.vue` - MCP testing interface
- `web/src/api/autoCode.js` - API endpoints for MCP operations
- `web/src/view/systemTools/aiWrokflow/index.vue` - AI workflow interface

## Authentication & Authorization

**JWT Authentication:**
- Library: `github.com/golang-jwt/jwt/v5` 5.2.2
- Configuration: `server/config.yaml` under `jwt` section
  - `signing-key` - Secret key for signing
  - `expires-time` - Token expiration (default: 7d)
  - `buffer-time` - Token refresh buffer (default: 1d)
  - `issuer` - Token issuer

**RBAC Authorization:**
- Library: Casbin 2.103.0 (`github.com/casbin/casbin/v2`)
- Adapter: `github.com/casbin/gorm-adapter/v3` 3.32.0
- Policy storage in GORM-compatible database
- Configuration: `server/config.yaml` under `system.use-strict-auth` for hierarchical roles

**Middleware:**
- `server/middleware/jwt.go` - JWT validation middleware
- `server/middleware/limit_ip.go` - IP rate limiting (15000 requests/hour by default)

## Email

**Library:** `github.com/jordan-wright/email` 4.0.1-0.20210109023952

**Configuration:** `server/config.yaml` under `email` section:
```yaml
email:
    to: xxx@qq.com
    port: 465
    from: xxx@163.com
    host: smtp.163.com
    is-ssl: true
    secret: xxx  # SMTP password
    nickname: test
```

## File Processing

**Archives:** `github.com/mholt/archives` 0.1.1 - Handles zip extraction for uploaded archives

**Excel:** `github.com/xuri/excelize/v2` 2.9.0 - Excel file reading/writing

**Captcha:**
- Library: `github.com/mojocn/base64Captcha` 1.3.8
- Redis storage for captcha state
- Configuration: `server/config.yaml` under `captcha` section

## Monitoring

**System Metrics:**
- Library: `github.com/shirou/gopsutil/v3` 3.24.5
- Used for disk usage monitoring (`server/config.yaml` `disk-list`)

**Structured Logging:**
- Library: `go.uber.org/zap` 1.27.0
- Configuration in `server/config.yaml` under `zap`:
  - `level`, `format`, `prefix`, `director`
  - `show-line`, `encode-level`, `stacktrace-key`
  - `log-in-console`, `retention-day`

## Deployment Platforms

**Docker:**
- Single container with all services (MySQL, Redis, Nginx, Go, npm)
- CentOS 7 base image
- Entrypoint script: `deploy/docker/entrypoint.sh`

**Docker Compose:**
- Full stack deployment via `deploy/docker-compose/docker-compose.yaml`
- Includes: MySQL, Redis, Nginx, Web, Server services

**Kubernetes:**
- Server deployment: `deploy/kubernetes/server/gva-server-*.yaml`
- Web deployment: `deploy/kubernetes/web/gva-web-*.yaml`
- ConfigMaps for configuration management
- Ingress for web routing

## Cross-Origin (CORS)

**Configuration:** `server/config.yaml` under `cors` section

**Modes:**
- `allow-all` - Allow all origins
- `whitelist` - Allow listed origins with credentials
- `strict-whitelist` - Reject non-whitelisted origins

**Whitelist Entry Structure:**
```yaml
cors:
    mode: strict-whitelist
    whitelist:
        - allow-origin: example.com
          allow-headers: Content-Type,Authorization,Token
          allow-methods: POST,GET
          expose-headers: Content-Length
          allow-credentials: true
```

---

*Integration audit: 2026-04-21*
