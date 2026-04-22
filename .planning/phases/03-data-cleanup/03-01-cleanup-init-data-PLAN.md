---
description: 清理数据库初始化中的示例数据
wave: 1
depends_on: []
requirements: [DATA-01, DATA-02, DATA-03, DATA-04]
autonomous: true
---

## Plan 03-01: 清理数据库初始化数据

### Objective
清理 server/source/ 目录中的示例数据，保留超级管理员账号

### Tasks

#### Task 1: 检查 source/ 目录结构
<read_first>
- ls server/source/ (查看数据初始化目录)
</read_first>
<action>
检查 server/source/ 目录下的文件结构，了解初始化数据的组织方式：
- 查找菜单初始化文件
- 查找角色初始化文件
- 查找用户初始化文件
</action>
<acceptance_criteria>
- server/source/ 目录存在且包含数据初始化文件
</acceptance_criteria>

#### Task 2: 清理示例菜单数据
<read_first>
- server/source/ 目录下的菜单相关文件
</read_first>
<action>
查找并清理所有示例菜单相关的初始化数据：
- 搜索包含 "example" 或 "示例" 的菜单数据
- 保留系统必要的菜单（如框架相关的菜单）
- 移除业务示例菜单
</action>
<acceptance_criteria>
- source/ 中无示例菜单数据
- 保留必要的框架菜单（如用户管理、角色管理等基础菜单）
</acceptance_criteria>

#### Task 3: 清理示例角色数据
<read_first>
- server/source/ 目录下的角色相关文件
</read_first>
<action>
查找并清理所有示例角色相关的初始化数据：
- 搜索包含 "example" 或示例角色
- 保留超级管理员角色（authorityID = 888）
- 保留普通管理员角色（如果有）
- 移除示例角色
</action>
<acceptance_criteria>
- source/ 中无示例角色数据
- 保留超级管理员角色
</acceptance_criteria>

#### Task 4: 清理示例用户数据
<read_first>
- server/source/ 目录下的用户相关文件
</read_first>
<action>
查找并清理所有示例用户相关的初始化数据：
- 搜索包含 "example" 或示例用户
- 保留超级管理员用户（admin）
- 确保 admin 用户密码为 123456
- 移除其他示例用户
</action>
<acceptance_criteria>
- source/ 中无示例用户数据
- 保留 admin 用户（密码 123456）
</acceptance_criteria>

#### Task 5: 验证数据完整性
<read_first>
- 清理后的 source/ 文件
</read_first>
<action>
验证清理后的数据完整性：
- 超级管理员角色存在
- 超级管理员用户存在且密码正确
- 必要的框架菜单存在
</action>
<acceptance_criteria>
- go build ./server/ 编译成功
- source/ 数据结构完整无语法错误
</acceptance_criteria>
