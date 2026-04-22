# Technology Stack

**Analysis Date:** 2026-04-21

## Languages

**Primary:**
- Go 1.24.0 - Backend server implementation
- Vue 3.5.31 - Frontend UI framework

**Secondary:**
- JavaScript/TypeScript - Frontend build and tooling

## Runtime

**Backend:**
- Go 1.24.0 with go1.24.2 toolchain

**Frontend:**
- Node.js (npm 11.3.0, core-js 3.38.1)

**Package Manager:**
- npm 11.3.0 (frontend)
- Go modules (backend)

## Backend Frameworks

**Core Web Framework:**
- Gin 1.10.0 (`github.com/gin-gonic/gin`) - HTTP web framework

**ORM:**
- GORM 1.25.12 (`gorm.io/gorm`) - Object-relational mapper
- GORM Gen 0.3.26 (`gorm.io/gen`) - Code generation for GORM

**Database Drivers (GORM):**
- MySQL: `gorm.io/driver/mysql` 1.5.7
- PostgreSQL: `gorm.io/driver/postgres` 1.5.11
- SQL Server: `gorm.io/driver/sqlserver` 1.5.4
- SQLite: `github.com/glebarez/sqlite` 1.11.0
- Oracle: `github.com/dzwvip/gorm-oracle` 0.1.2
- MongoDB: `go.mongodb.org/mongo-driver` 1.17.2 + `github.com/qiniu/qmgo` 1.1.9

**Authentication & Authorization:**
- JWT: `github.com/golang-jwt/jwt/v5` 5.2.2
- Casbin 2.103.0 (`github.com/casbin/casbin/v2`) - RBAC permission framework
- Casbin GORM Adapter 3.32.0 (`github.com/casbin/gorm-adapter/v3`)

**Configuration:**
- Viper 1.19.0 (`github.com/spf13/viper`) - Config management from YAML

**Logging:**
- Zap 1.27.0 (`go.uber.org/zap`) - Structured logging
- gookit/color 1.5.4 - Terminal colors

**Caching:**
- Redis: `github.com/redis/go-redis/v9` 9.7.0

**API Documentation:**
- Swagger: `github.com/swaggo/swag` 1.16.4
- Gin Swagger: `github.com/swaggo/gin-swagger` 1.6.0
- Swagger Files: `github.com/swaggo/files` 1.0.1

**Utilities:**
- UUID: `github.com/google/uuid` 1.6.0
- Excel: `github.com/xuri/excelize/v2` 2.9.0
- Cron: `github.com/robfig/cron/v3` 3.0.1
- Captcha: `github.com/mojocn/base64Captcha` 1.3.8
- Email: `github.com/jordan-wright/email` 4.0.1-0.20210109023952

**Performance:**
- automaxprocs 1.6.0 (`go.uber.org/automaxprocs`) - CPU limit optimization

## Frontend Frameworks

**Core:**
- Vue 3.5.31 - Progressive JavaScript framework
- Vue Router 4.4.3 (`vue-router`) - Client-side routing
- Pinia 2.2.2 - State management

**Build:**
- Vite 6.2.3 (`vite`) - Next-gen frontend build tool
- @vitejs/plugin-vue 5.0.3 - Vue 3 SFC plugin for Vite

**UI Components:**
- Element Plus 2.13.6 (`element-plus`) - Vue 3 UI library
- @element-plus/icons-vue 2.3.1 - Element Plus icons

**CSS:**
- UnoCSS 66.5.0 (`@unocss/vite`, `@unocss/preset-wind3`, `@unocss/extractor-svelte`, `@unocss/transformer-directives`) - Atomic CSS framework
- Sass 1.78.0 - CSS preprocessor

**HTTP Client:**
- Axios 1.8.2 - Promise-based HTTP client

**Charts:**
- ECharts 5.5.1 (`echarts`) - Interactive charting library
- vue-echarts 7.0.3 - ECharts wrapper for Vue

**Utilities:**
- @vueuse/core 11.0.3 - Vue composition utilities
- @vueuse/integrations 12.0.0
- mitt 3.0.1 - Event emitter
- nprogress 0.2.0 - Progress bar
- qs 6.13.0 - Query string parsing
- sortablejs 1.15.3 - Drag and drop
- spark-md5 3.0.2 - MD5 implementation
- universal-cookie 7 - Cookie management

**Form Building:**
- @form-create/designer 3.2.6 - Form designer
- @form-create/element-ui 3.2.10 - Form generation
- vform3-builds 3.0.10 - Form builder

**Rich Text Editor:**
- @wangeditor/editor 5.1.23 - Rich text editor
- @wangeditor/editor-for-vue 5.1.12 - Vue bindings

**Code Editor:**
- ace-builds 1.36.4 - Code editor
- vue3-ace-editor 2.2.4 - Ace editor for Vue 3

**Office Documents:**
- @vue-office/docx 1.6.2 - Word document viewer
- @vue-office/excel 1.7.11 - Excel document viewer
- @vue-office/pdf 2.0.2 - PDF document viewer

**Image Processing:**
- vue-cropper 1.1.4 - Image cropping
- marked 14.1.1 - Markdown parser
- marked-highlight 2.1.4 - Markdown with highlight.js

**QR Code:**
- vue-qr 4.0.9 - QR code generator

**Vite Plugins:**
- vite-plugin-vue-devtools 7.0.16 - Vue DevTools integration
- vite-plugin-banner 0.8.0 - Add banner to output
- vite-plugin-importer 0.2.5 - Import resolution
- vite-auto-import-svg 2.5.0 - Auto import SVG
- vite-check-multiple-dom 0.2.1 - Multi-domain checking
- @vitejs/plugin-legacy 6.0.0 - Legacy browser support

**Linting/Formatting:**
- ESLint 8.57.0 with eslint-plugin-vue 9.19.2
- @babel/eslint-parser 7.25.1

## Cloud Storage Integrations

**Supported Providers (via `server/utils/upload/upload.go`):**

1. **Local** - Default local filesystem storage
2. **MinIO** - `github.com/minio/minio-go/v7` 7.0.84 - S3-compatible object storage
3. **Aliyun OSS** - `github.com/aliyun/aliyun-oss-go-sdk` 3.0.2 - Alibaba Cloud Object Storage
4. **AWS S3** - `github.com/aws/aws-sdk-go-v2` with S3 service - Amazon Simple Storage Service
5. **Tencent COS** - `github.com/tencentyun/cos-go-sdk-v5` 0.7.60 - Tencent Cloud Object Storage
6. **Qiniu** - `github.com/qiniu/go-sdk/v7` 7.25.2 - Qiniu Cloud Object Storage
7. **Huawei OBS** - `github.com/huaweicloud/huaweicloud-sdk-go-obs` 3.24.9 - Huawei Cloud Object Storage
8. **Cloudflare R2** - S3-compatible storage (via AWS S3 SDK)

**Configuration location:** `server/config.yaml` under sections:
- `minio` - MinIO configuration
- `aliyun-oss` - Aliyun OSS configuration
- `aws-s3` - AWS S3 / R2 configuration
- `tencent-cos` - Tencent COS configuration
- `qiniu` - Qiniu configuration
- `hua-wei-obs` - Huawei OBS configuration
- `cloudflare-r2` - Cloudflare R2 configuration

**Upload factory:** `server/utils/upload/upload.go` - `NewOss()` function selects OSS provider based on `global.GVA_CONFIG.System.OssType`

## Deployment

**Docker:**
- Base image: `centos:7`
- Location: `deploy/docker/Dockerfile`
- Entrypoint: `deploy/docker/entrypoint.sh`

**Docker Compose:**
- Configuration: `deploy/docker-compose/docker-compose.yaml`
- Includes MySQL, Redis, Nginx, and application services

**Kubernetes:**
- Server manifests: `deploy/kubernetes/server/`
  - `gva-server-configmap.yaml` - ConfigMap for server
  - `gva-server-deployment.yaml` - Deployment for server
  - `gva-server-service.yaml` - Service for server
- Web manifests: `deploy/kubernetes/web/`
  - `gva-web-configmap.yaml` - ConfigMap for web
  - `gva-web-deploymemt.yaml` - Deployment for web
  - `gva-web-service.yaml` - Service for web
  - `gva-web-ingress.yaml` - Ingress for web

## Additional Backend Components

**File Processing:**
- `github.com/mholt/archives` 0.1.1 - Archive extraction
- `github.com/otiai10/copy` 1.14.1 - File copy utility

**System Info:**
- `github.com/shirou/gopsutil/v3` 3.24.5 - System/utility library

**Security:**
- `golang.org/x/crypto` 0.37.0 - Cryptographic functions
- `github.com/unrolled/secure` 1.17.0 - Security headers middleware

**Code Generation Tools:**
- `github.com/songzhibin97/gkit` 1.2.13 - Go utility toolkit

---

*Stack analysis: 2026-04-21*
