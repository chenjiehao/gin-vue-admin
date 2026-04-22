---
description: 删除前端示例目录和 API 文件
wave: 1
depends_on: []
requirements: [FRONT-01, FRONT-02, FRONT-03, FRONT-04, FRONT-05]
autonomous: true
---

## Plan 02-01: 删除前端示例目录和文件

### Objective
删除 web/src/view/example/ 目录和 web/src/api/example.js 文件

### Tasks

#### Task 1: 删除 web/src/view/example/ 目录
<read_first>
- ls web/src/view/ (查看现有目录)
</read_first>
<action>
删除整个 web/src/view/example/ 目录及其所有文件：
- web/src/view/example/ 目录及所有 .vue 文件
</action>
<acceptance_criteria>
- web/src/view/example/ 目录不存在
- ls web/src/view/ 不包含 example 目录
</acceptance_criteria>

#### Task 2: 删除 web/src/api/example.js 文件
<read_first>
- ls web/src/api/ (查看现有 API 文件)
</read_first>
<action>
删除 web/src/api/example.js 文件
</action>
<acceptance_criteria>
- web/src/api/example.js 文件不存在
- ls web/src/api/ 不包含 example.js
</acceptance_criteria>

#### Task 3: 更新相关引用（如果有）
<read_first>
- web/src/router/index.js (路由配置)
- web/src/view/layout/ (布局目录)
</read_first>
<action>
检查是否有其他地方引用 example 目录或文件，如有则移除
</action>
<acceptance_criteria>
- 路由配置中无 example 相关路由
- 无其他文件引用 example 目录或 example.js
</acceptance_criteria>
