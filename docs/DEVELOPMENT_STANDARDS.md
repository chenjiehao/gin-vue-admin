# Gin-Vue-Admin 开发规范

## 1. API 权限配置规范 (Casbin)

### 1.1 问题背景

GVA 框架使用 Casbin 做 RBAC 权限控制。当新增 API 接口后，即使使用超级管理员账号访问也返回"权限不足"或"未登录"错误。

### 1.2 常见错误表现

```
{"code": 7, "msg": "未登录或非法访问，请登录"}
{"code": 7, "msg": "权限不足"}
```

### 1.3 解决方案

#### 方式一：通过数据库初始化文件（推荐）

在 `server/source/system/casbin.go` 中添加权限规则：

```go
{Ptype: "p", V0: "888", V1: "/dataSource/getTables", V2: "GET"},
{Ptype: "p", V0: "888", V1: "/dataSource/getDatabases", V2: "GET"},
```

其中：
- `V0`: 角色 ID（`888` = 超级管理员，`9528` = 普通用户）
- `V1`: API 路径
- `V2`: HTTP 方法（GET/POST/PUT/DELETE）

#### 方式二：手动数据库插入

权限规则存储在**业务数据库**中（通常是 `qmPlus`，不是 `gva`）：

```sql
USE qmPlus;
INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES
('p', '888', '/dataSource/getTables', 'GET'),
('p', '888', '/dataSource/getDatabases', 'GET');
```

**注意**：配置后需要**重启后端服务**使 Casbin 缓存失效。

### 1.4 排查步骤

1. 检查 `server/source/system/casbin.go` 是否包含该 API 规则
2. 确认业务数据库（qmPlus）的 `casbin_rule` 表中是否有对应规则
3. 重启后端服务
4. 清除浏览器缓存或重新登录

---

## 2. 前端数据绑定规范

### 2.1 GVA 后端数据模型 ID 字段规范

GVA 后端返回的数据模型中，**主键字段统一使用大写 `ID`**，而不是小写 `id`。

**示例**：
```javascript
// 后端返回的数据结构
{ ID: 1, CreatedAt: "2026-04-22T14:28:46+08:00", name: "test", ... }

// ❌ 错误写法
item.id  // undefined

// ✅ 正确写法
item.ID  // 1
```

### 2.2 el-select 组件 value 绑定建议

为避免 ID 类型转换问题，建议：

```vue
<!-- 方案一：使用 name 作为 value（推荐） -->
<el-option
  v-for="item in dataSourceList"
  :key="item.ID"
  :label="item.name"
  :value="item.name"  <!-- 显示值用于 UI，选中后便于识别 -->
/>

<!-- 方案二：存储整个对象 -->
<el-option
  v-for="item in dataSourceList"
  :key="item.ID"
  :label="item.name"
  :value="item"  <!-- 存储整个对象，需要时取 item.ID -->
/>
```

### 2.3 change 事件处理

```javascript
const handleChange = (selectedItem) => {
  if (selectedItem && typeof selectedItem === 'object') {
    // 使用 item.ID 访问数据库 ID
    fillDatabaseList(selectedItem.ID, 'source')
  } else if (typeof selectedItem === 'string') {
    // 使用 name 查找对应的 item
    const item = dataSourceList.value.find(i => i.name === selectedItem)
    if (item) {
      fillDatabaseList(item.ID, 'source')
    }
  }
}
```

---

## 3. 下拉框联动规范

### 3.1 四级联动场景

典型场景：类型 → 数据源 → 数据库 → 表

### 3.2 联动实现规范

```javascript
// 1. 监听类型变化 → 加载数据源列表
watch(() => taskForm.value.sourceType, (newVal) => {
  if (newVal) {
    getDataSourcesByType(newVal, 'source')
  }
})

// 2. change 事件处理数据源选择 → 加载数据库列表
const handleSourceDataSourceChange = (selectedItem) => {
  if (selectedItem) {
    fillDatabaseList(selectedItem.ID, 'source')
  }
}

// 3. 监听数据库变化 → 加载表列表
watch(() => taskForm.value.sourceDatabase, (newVal) => {
  if (newVal && taskForm.value.sourceDataSourceId) {
    const item = sourceDataSourceList.value.find(
      i => i.name === taskForm.value.sourceDataSourceId
    )
    if (item) {
      loadTableList(item.ID, newVal, 'source')
    }
  }
})
```

### 3.3 清除联动

每当上级选择变化时，需要清除下级已选值：

```javascript
if (target === 'source') {
  taskForm.value.sourceDatabase = ''
  taskForm.value.sourceTable = ''
  sourceDatabaseList.value = []
  sourceTableList.value = []
}
```

---

## 4. 调试技巧

### 4.1 后端日志检查

```bash
# 查看最新的 API 请求日志
tail -f /tmp/gva_server.log

# 检查特定 API 的响应
grep "getDatabases" /tmp/gva_server.log
```

### 4.2 前端调试日志

```javascript
const fillDatabaseList = async (dataSourceId, target) => {
  console.log('fillDatabaseList called:', { dataSourceId, target })
  const res = await getDataSourceDatabases(dataSourceId)
  console.log('response:', res)
  // ...
}
```

### 4.3 数据库权限检查

```bash
# 使用业务数据库用户连接
mysql -u{gva_user} -p{password} -h 127.0.0.1 -P 13306

# 查看现有权限规则
USE qmPlus;
SELECT * FROM casbin_rule WHERE v1 LIKE '%dataSource%';
```

---

## 5. 常见问题排查

| 问题现象 | 可能原因 | 解决方案 |
|---------|---------|---------|
| "权限不足" | Casbin 规则缺失 | 添加规则到 `casbin_rule` 表 |
| "权限不足" 但规则存在 | 规则在错误的数据库 | 确认是 `qmPlus` 而非 `gva` |
| "权限不足" | 服务未重启 | 重启后端使缓存失效 |
| 下拉框无法选择 | el-option value 类型不匹配 | 使用 `name` 或 `item` 作为 value |
| 选中后值为 undefined | ID 字段大小写错误 | 使用 `item.ID` 而非 `item.id` |
