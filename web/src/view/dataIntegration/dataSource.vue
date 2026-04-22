<template>
  <div>
    <!-- 列表视图 -->
    <div v-if="!formVisible">
      <div class="gva-search-box">
        <el-form :inline="true" :model="searchInfo">
          <el-form-item label="数据源名称">
            <el-input v-model="searchInfo.name" placeholder="搜索数据源名称" />
          </el-form-item>
          <el-form-item label="类型">
            <el-select v-model="searchInfo.type" placeholder="请选择类型" clearable>
              <el-option label="MySQL" value="MySQL" />
              <el-option label="PostgreSQL" value="PostgreSQL" />
              <el-option label="达梦" value="达梦" />
              <el-option label="人大金仓" value="人大金仓" />
              <el-option label="Oracle" value="Oracle" />
              <el-option label="SQLServer" value="SQLServer" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
            <el-button icon="refresh" @click="onReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div class="gva-table-box">
        <div class="gva-btn-list">
          <el-button icon="plus" type="primary" @click="goCreate">新增数据源</el-button>
          <el-button icon="refresh" @click="getTableData">刷新</el-button>
          <span class="selection-count" v-if="multipleSelection.length > 0">已选择 {{ multipleSelection.length }} 项</span>
          <el-button icon="delete" type="danger" @click="handleBatchDelete" :disabled="multipleSelection.length === 0">批量删除</el-button>
          <el-button icon="check" type="success" @click="handleBatchEnable" :disabled="multipleSelection.length === 0">批量启用</el-button>
          <el-button icon="close" type="warning" @click="handleBatchDisable" :disabled="multipleSelection.length === 0">批量禁用</el-button>
        </div>
        <el-table
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="id"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="left" />
          <el-table-column align="left" label="数据源名称" prop="name" width="180" />
          <el-table-column align="left" label="类型" prop="type" width="120">
            <template #default="scope">
              <el-tag>{{ scope.row.type }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column align="left" label="主机" prop="host" width="180" />
          <el-table-column align="left" label="端口" prop="port" width="100" />
          <el-table-column align="left" label="数据库" prop="database" width="150" />
          <el-table-column align="left" label="状态" prop="status" width="100">
            <template #default="scope">
              <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
                {{ scope.row.status === 1 ? '启用' : '禁用' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column align="left" label="备注" prop="remark" min-width="150" />
          <el-table-column align="left" label="操作" min-width="150">
            <template #default="scope">
              <el-button icon="edit" type="primary" link @click="goEdit(scope.row)">编辑</el-button>
              <el-button icon="delete" type="primary" link @click="deleteDataSource(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination">
          <el-pagination
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
          />
        </div>
      </div>
    </div>

    <!-- 表单视图 -->
    <div v-else class="form-view">
      <div class="gva-table-box">
        <div class="gva-btn-list">
          <el-button icon="ArrowLeft" @click="closeForm">返回列表</el-button>
        </div>

        <el-form
          ref="dataSourceFormRef"
          :model="dataSourceForm"
          :rules="dataSourceRules"
          label-width="100px"
          class="data-source-form"
        >
          <el-divider content-position="left">基本信息</el-divider>
          <el-row :gutter="16">
            <el-col :span="12">
              <el-form-item label="数据源名称" prop="name">
                <el-input v-model="dataSourceForm.name" placeholder="请输入数据源名称" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="数据源类型" prop="type">
                <el-select v-model="dataSourceForm.type" placeholder="请选择数据源类型" class="full-width">
                  <el-option label="MySQL" value="MySQL" />
                  <el-option label="PostgreSQL" value="PostgreSQL" />
                  <el-option label="达梦" value="达梦" />
                  <el-option label="人大金仓" value="人大金仓" />
                  <el-option label="Oracle" value="Oracle" />
                  <el-option label="SQLServer" value="SQLServer" />
                </el-select>
              </el-form-item>
            </el-col>
          </el-row>

          <el-divider content-position="left">连接配置</el-divider>
          <el-row :gutter="16">
            <el-col :span="10">
              <el-form-item label="主机地址" prop="host">
                <el-input v-model="dataSourceForm.host" placeholder="请输入主机地址" />
              </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="端口" prop="port">
                <el-input-number v-model="dataSourceForm.port" :min="1" :max="65535" placeholder="端口" class="full-width" />
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="数据库名称" prop="database">
                <el-input v-model="dataSourceForm.database" placeholder="请输入数据库名称" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="16">
            <el-col :span="12">
              <el-form-item label="用户名" prop="username">
                <el-input v-model="dataSourceForm.username" placeholder="请输入用户名" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="密码" prop="password">
                <el-input v-model="dataSourceForm.password" type="password" placeholder="请输入密码" show-password />
              </el-form-item>
            </el-col>
          </el-row>

          <el-divider content-position="left">状态配置</el-divider>
          <el-row :gutter="16">
            <el-col :span="12">
              <el-form-item label="状态" prop="status">
                <el-radio-group v-model="dataSourceForm.status">
                  <el-radio :label="1">启用</el-radio>
                  <el-radio :label="0">禁用</el-radio>
                </el-radio-group>
              </el-form-item>
            </el-col>
          </el-row>

          <el-divider content-position="left">备注</el-divider>
          <el-row :gutter="16">
            <el-col :span="24">
              <el-form-item label="备注">
                <el-input
                  v-model="dataSourceForm.remark"
                  type="textarea"
                  placeholder="请输入备注"
                  :rows="3"
                />
              </el-form-item>
            </el-col>
          </el-row>

          <el-form-item>
            <el-button type="primary" @click="saveDataSource">{{ isEdit ? '保存修改' : '创建数据源' }}</el-button>
            <el-button type="success" @click="handleTestConnection" :loading="testingConnection">测试连接</el-button>
            <el-button @click="closeForm">取消</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup>
  import { getDataSourceList, deleteDataSource as apiDeleteDataSource, createDataSource, updateDataSource, testConnection, batchDelete, batchUpdateStatus } from '@/api/dataSource'
  import { ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'

  defineOptions({
    name: 'dataSource'
  })

  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})
  const multipleSelection = ref([])

  // 表单相关状态
  const formVisible = ref(false)
  const dataSourceFormRef = ref(null)
  const isEdit = ref(false)
  const testingConnection = ref(false)

  const dataSourceForm = ref({
    id: null,
    name: '',
    type: '',
    host: '',
    port: 3306,
    database: '',
    username: '',
    password: '',
    status: 1,
    remark: ''
  })

  const dataSourceRules = {
    name: [{ required: true, message: '请输入数据源名称', trigger: 'blur' }],
    type: [{ required: true, message: '请选择数据源类型', trigger: 'change' }],
    host: [{ required: true, message: '请输入主机地址', trigger: 'blur' }],
    port: [{ required: true, message: '请输入端口', trigger: 'blur' }],
    database: [{ required: true, message: '请输入数据库名称', trigger: 'blur' }]
  }

  // 打开表单（新增）
  const goCreate = () => {
    isEdit.value = false
    dataSourceForm.value = {
      id: null,
      name: '',
      type: '',
      host: '',
      port: 3306,
      database: '',
      username: '',
      password: '',
      status: 1,
      remark: ''
    }
    formVisible.value = true
  }

  // 打开表单（编辑）
  const goEdit = (row) => {
    isEdit.value = true
    dataSourceForm.value = { ...row }
    formVisible.value = true
  }

  // 关闭表单
  const closeForm = () => {
    formVisible.value = false
    dataSourceForm.value = {
      id: null,
      name: '',
      type: '',
      host: '',
      port: 3306,
      database: '',
      username: '',
      password: '',
      status: 1,
      remark: ''
    }
  }

  // 保存数据源
  const saveDataSource = async () => {
    if (!dataSourceFormRef.value) return
    await dataSourceFormRef.value.validate(async (valid) => {
      if (valid) {
        try {
          if (isEdit.value) {
            const res = await updateDataSource(dataSourceForm.value)
            if (res.code === 0) {
              ElMessage({ type: 'success', message: '修改成功' })
              closeForm()
              getTableData()
            } else {
              ElMessage({ type: 'error', message: res.msg || '修改失败' })
            }
          } else {
            const res = await createDataSource(dataSourceForm.value)
            if (res.code === 0) {
              ElMessage({ type: 'success', message: '创建成功' })
              closeForm()
              getTableData()
            } else {
              ElMessage({ type: 'error', message: res.msg || '创建失败' })
            }
          }
        } catch (e) {
          ElMessage({ type: 'error', message: '保存失败' })
        }
      }
    })
  }

  // 获取表格数据
  const getTableData = async () => {
    try {
      const res = await getDataSourceList({
        page: page.value,
        pageSize: pageSize.value,
        keyword: searchInfo.value.name || ''
      })
      if (res.code === 0) {
        tableData.value = res.data.list || []
        total.value = res.data.total || 0
      }
    } catch (e) {
      console.error('获取数据源列表失败:', e)
    }
  }

  const onReset = () => {
    searchInfo.value = {}
    page.value = 1
    getTableData()
  }

  const onSubmit = () => {
    page.value = 1
    getTableData()
  }

  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }

  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }

  // 多选
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }

  // 删除数据源
  const deleteDataSource = async (row) => {
    ElMessageBox.confirm('确定要删除该数据源吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      try {
        const res = await apiDeleteDataSource({ id: row.id })
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '删除成功' })
          getTableData()
        } else {
          ElMessage({ type: 'error', message: res.msg || '删除失败' })
        }
      } catch (e) {
        ElMessage({ type: 'error', message: '删除失败' })
      }
    })
  }

  // 测试连接
  const handleTestConnection = async () => {
    if (!dataSourceFormRef.value) return
    await dataSourceFormRef.value.validateField(['name', 'type', 'host', 'port', 'database'], async (valid) => {
      if (valid) {
        testingConnection.value = true
        try {
          const res = await testConnection(dataSourceForm.value)
          if (res.code === 0) {
            const data = res.data || {}
            if (data.success) {
              ElMessage({ type: 'success', message: data.message || '连接成功' })
              // 连接成功后返回列表页面
              closeForm()
              getTableData()
            } else {
              ElMessage({ type: 'error', message: data.message || '连接失败' })
            }
          } else {
            ElMessage({ type: 'error', message: res.msg || '测试连接失败' })
          }
        } catch (e) {
          ElMessage({ type: 'error', message: '测试连接失败' })
        } finally {
          testingConnection.value = false
        }
      }
    })
  }

  // 批量删除
  const handleBatchDelete = async () => {
    if (multipleSelection.value.length === 0) {
      ElMessage({ type: 'warning', message: '请先选择要删除的数据源' })
      return
    }
    ElMessageBox.confirm(`确定要删除选中的 ${multipleSelection.value.length} 个数据源吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const ids = multipleSelection.value.map(item => item.id)
      try {
        const res = await batchDelete({ ids })
        if (res.code === 0) {
          const data = res.data || {}
          ElMessage({ type: 'success', message: `成功删除 ${data.successCount || 0} 个数据源` })
          getTableData()
          multipleSelection.value = []
        } else {
          ElMessage({ type: 'error', message: res.msg || '批量删除失败' })
        }
      } catch (e) {
        ElMessage({ type: 'error', message: '批量删除失败' })
      }
    }).catch(() => {})
  }

  // 批量启用
  const handleBatchEnable = async () => {
    if (multipleSelection.value.length === 0) {
      ElMessage({ type: 'warning', message: '请先选择要启用的数据源' })
      return
    }
    ElMessageBox.confirm(`确定要启用选中的 ${multipleSelection.value.length} 个数据源吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const ids = multipleSelection.value.map(item => item.id)
      try {
        const res = await batchUpdateStatus({ ids, status: 1 })
        if (res.code === 0) {
          const data = res.data || {}
          ElMessage({ type: 'success', message: `成功启用 ${data.successCount || 0} 个数据源` })
          getTableData()
          multipleSelection.value = []
        } else {
          ElMessage({ type: 'error', message: res.msg || '批量启用失败' })
        }
      } catch (e) {
        ElMessage({ type: 'error', message: '批量启用失败' })
      }
    }).catch(() => {})
  }

  // 批量禁用
  const handleBatchDisable = async () => {
    if (multipleSelection.value.length === 0) {
      ElMessage({ type: 'warning', message: '请先选择要禁用的数据源' })
      return
    }
    ElMessageBox.confirm(`确定要禁用选中的 ${multipleSelection.value.length} 个数据源吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      const ids = multipleSelection.value.map(item => item.id)
      try {
        const res = await batchUpdateStatus({ ids, status: 0 })
        if (res.code === 0) {
          const data = res.data || {}
          ElMessage({ type: 'success', message: `成功禁用 ${data.successCount || 0} 个数据源` })
          getTableData()
          multipleSelection.value = []
        } else {
          ElMessage({ type: 'error', message: res.msg || '批量禁用失败' })
        }
      } catch (e) {
        ElMessage({ type: 'error', message: '批量禁用失败' })
      }
    }).catch(() => {})
  }

  // 初始化
  getTableData()
</script>

<style lang="scss" scoped>
.form-view {
  .data-source-form {
    padding: 20px;
    background: #fff;
    border-radius: 8px;
  }

  .full-width {
    width: 100%;
  }

  :deep(.el-divider--horizontal) {
    margin: 16px 0;
  }
}

.selection-count {
  margin-left: 16px;
  color: #909399;
  font-size: 14px;
}
</style>
