<template>
  <div>
    <!-- 列表视图 -->
    <div v-if="!formVisible">
      <div class="gva-search-box">
        <el-form :inline="true" :model="searchInfo">
          <el-form-item label="任务名称">
            <el-input v-model="searchInfo.taskName" placeholder="搜索任务名称" />
          </el-form-item>
          <el-form-item label="状态">
            <el-select v-model="searchInfo.status" placeholder="请选择状态" clearable>
              <el-option label="待执行" value="pending" />
              <el-option label="运行中" value="running" />
              <el-option label="成功" value="success" />
              <el-option label="失败" value="failed" />
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
          <el-button icon="plus" type="primary" @click="goCreate">新增任务</el-button>
          <el-button icon="refresh" @click="getTableData">刷新</el-button>
        </div>
        <el-table
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="id"
          @selection-change="handleSelectionChange"
        >
          <el-table-column align="left" type="selection" width="55" />
          <el-table-column align="left" label="任务名称" prop="taskName" width="180">
            <template #default="scope">
              <el-link type="primary" :underline="false">{{ scope.row.taskName }}</el-link>
            </template>
          </el-table-column>
          <el-table-column align="left" label="源类型" prop="sourceType" width="100">
            <template #default="scope">
              <el-tag>{{ scope.row.sourceType }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column align="left" label="目标类型" prop="targetType" width="100">
            <template #default="scope">
              <el-tag>{{ scope.row.targetType }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column align="left" label="状态" prop="status" width="100">
            <template #default="scope">
              <el-tag :type="getStatusType(scope.row.status)">{{ getStatusLabel(scope.row.status) }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column align="left" label="上次同步时间" prop="lastSyncTime" width="180">
            <template #default="scope">
              {{ scope.row.lastSyncTime ? formatDate(scope.row.lastSyncTime) : '-' }}
            </template>
          </el-table-column>
          <el-table-column align="left" label="下次同步时间" prop="nextSyncTime" width="180">
            <template #default="scope">
              {{ scope.row.nextSyncTime ? formatDate(scope.row.nextSyncTime) : '-' }}
            </template>
          </el-table-column>
          <el-table-column align="left" label="操作" min-width="180">
            <template #default="scope">
              <el-button icon="edit" type="primary" link @click="goEdit(scope.row)">编辑</el-button>
              <el-button icon="delete" type="primary" link @click="deleteTask(scope.row)">删除</el-button>
              <el-button icon="video-play" type="primary" link @click="handleTriggerSync(scope.row)">触发</el-button>
              <el-button icon="histogram" type="primary" link @click="viewHistory(scope.row)">历史</el-button>
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

      <!-- 同步历史对话框 -->
      <el-dialog v-model="historyDialogVisible" title="同步历史" width="60%" :before-close="handleHistoryDialogClose" destroy-on-close>
        <el-table :data="historyTableData" style="width: 100%">
          <el-table-column align="left" label="执行时间" prop="executeTime" width="180">
            <template #default="scope">
              {{ formatDate(scope.row.executeTime) }}
            </template>
          </el-table-column>
          <el-table-column align="left" label="状态" prop="status" width="100">
            <template #default="scope">
              <el-tag :type="getStatusType(scope.row.status)">{{ getStatusLabel(scope.row.status) }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column align="left" label="执行耗时" prop="duration" width="100">
            <template #default="scope">
              {{ scope.row.duration }}ms
            </template>
          </el-table-column>
          <el-table-column align="left" label="同步数量" prop="syncCount" width="100" />
          <el-table-column align="left" label="错误信息" prop="errorMsg" min-width="200">
            <template #default="scope">
              <span v-if="scope.row.errorMsg" class="error-text">{{ scope.row.errorMsg }}</span>
              <span v-else>-</span>
            </template>
          </el-table-column>
        </el-table>
        <template #footer>
          <el-button @click="handleHistoryDialogClose">关闭</el-button>
        </template>
      </el-dialog>
    </div>

    <!-- 表单视图 -->
    <div v-else class="form-view">
      <div class="gva-table-box">
        <div class="gva-btn-list">
          <el-button icon="ArrowLeft" @click="closeForm">返回列表</el-button>
        </div>

        <el-form
          ref="taskFormRef"
          :model="taskForm"
          :rules="taskRules"
          label-width="100px"
          class="task-form"
        >
          <el-divider content-position="left">基本信息</el-divider>
          <el-row :gutter="16">
            <el-col :span="24">
              <el-form-item label="任务名称" prop="taskName">
                <el-input v-model="taskForm.taskName" placeholder="请输入任务名称" />
              </el-form-item>
            </el-col>
          </el-row>

          <el-divider content-position="left">数据源配置</el-divider>
          <el-row :gutter="16">
            <!-- 源配置 -->
            <el-col :span="12">
              <el-card shadow="never" class="source-card">
                <template #header>
                  <div class="card-header">
                    <span>源配置</span>
                  </div>
                </template>
                <el-form-item label="源类型" prop="sourceType">
                  <el-select v-model="taskForm.sourceType" placeholder="请选择源类型" class="full-width" clearable placement="bottom-start">
                    <el-option label="MySQL" value="MySQL" />
                    <el-option label="PostgreSQL" value="PostgreSQL" />
                    <el-option label="达梦" value="达梦" />
                    <el-option label="人大金仓" value="人大金仓" />
                    <el-option label="Oracle" value="Oracle" />
                    <el-option label="SQLServer" value="SQLServer" />
                  </el-select>
                </el-form-item>
                <el-form-item label="源数据源" prop="sourceDataSourceId">
                  <el-select
                    v-model="taskForm.sourceDataSourceId"
                    placeholder="请先选择源类型"
                    class="full-width"
                    clearable
                    placement="bottom-start"
                    :disabled="!taskForm.sourceType"
                    @change="handleSourceDataSourceChange"
                  >
                    <el-option
                      v-for="item in sourceDataSourceList"
                      :key="item.ID"
                      :label="item.name"
                      :value="item.name"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item label="源数据库" prop="sourceDatabase">
                  <el-select
                    v-model="taskForm.sourceDatabase"
                    placeholder="请先选择源数据源"
                    class="full-width"
                    clearable
                    placement="bottom-start"
                    :disabled="!taskForm.sourceDataSourceId"
                  >
                    <el-option
                      v-for="item in sourceDatabaseList"
                      :key="item"
                      :label="item"
                      :value="item"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item label="源表" prop="sourceTable">
                  <el-select
                    v-model="taskForm.sourceTable"
                    placeholder="请先选择源数据库"
                    class="full-width"
                    clearable
                    placement="bottom-start"
                    :disabled="!taskForm.sourceDatabase"
                  >
                    <el-option
                      v-for="item in sourceTableList"
                      :key="item"
                      :label="item"
                      :value="item"
                    />
                  </el-select>
                </el-form-item>
              </el-card>
            </el-col>
            <!-- 目标配置 -->
            <el-col :span="12">
              <el-card shadow="never" class="target-card">
                <template #header>
                  <div class="card-header">
                    <span>目标配置</span>
                  </div>
                </template>
                <el-form-item label="目标类型" prop="targetType">
                  <el-select v-model="taskForm.targetType" placeholder="请选择目标类型" class="full-width" clearable placement="bottom-start">
                    <el-option label="MySQL" value="MySQL" />
                    <el-option label="PostgreSQL" value="PostgreSQL" />
                    <el-option label="达梦" value="达梦" />
                    <el-option label="人大金仓" value="人大金仓" />
                    <el-option label="Oracle" value="Oracle" />
                    <el-option label="SQLServer" value="SQLServer" />
                  </el-select>
                </el-form-item>
                <el-form-item label="目标数据源" prop="targetDataSourceId">
                  <el-select
                    v-model="taskForm.targetDataSourceId"
                    placeholder="请先选择目标类型"
                    class="full-width"
                    clearable
                    placement="bottom-start"
                    :disabled="!taskForm.targetType"
                    @change="handleTargetDataSourceChange"
                  >
                    <el-option
                      v-for="item in targetDataSourceList"
                      :key="item.ID"
                      :label="item.name"
                      :value="item.name"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item label="目标数据库" prop="targetDatabase">
                  <el-select
                    v-model="taskForm.targetDatabase"
                    placeholder="请先选择目标数据源"
                    class="full-width"
                    clearable
                    placement="bottom-start"
                    :disabled="!taskForm.targetDataSourceId"
                  >
                    <el-option
                      v-for="item in targetDatabaseList"
                      :key="item"
                      :label="item"
                      :value="item"
                    />
                  </el-select>
                </el-form-item>
                <el-form-item label="目标表" prop="targetTable">
                  <el-select
                    v-model="taskForm.targetTable"
                    placeholder="请先选择目标数据库"
                    class="full-width"
                    clearable
                    placement="bottom-start"
                    :disabled="!taskForm.targetDatabase"
                  >
                    <el-option
                      v-for="item in targetTableList"
                      :key="item"
                      :label="item"
                      :value="item"
                    />
                  </el-select>
                </el-form-item>
              </el-card>
            </el-col>
          </el-row>

          <el-divider content-position="left">同步规则</el-divider>
          <el-row :gutter="16">
            <el-col :span="12">
              <el-form-item label="字段映射规则" prop="mappingRules">
                <el-input
                  v-model="taskForm.mappingRules"
                  type="textarea"
                  placeholder="请输入字段映射规则(JSON)"
                  :rows="3"
                />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="CRON表达式" prop="cronExpression">
                <el-input v-model="taskForm.cronExpression" placeholder="如: 0 0 * * *" />
                <div class="form-tip">格式: 秒 分 时 日 月 周，每周日凌晨零点: 0 0 0 * * 1</div>
              </el-form-item>
            </el-col>
          </el-row>

          <el-form-item>
            <el-button type="primary" @click="saveTask">{{ isEdit ? '保存修改' : '创建任务' }}</el-button>
            <el-button @click="closeForm">取消</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup>
  import {
    getOfflineSyncList,
    createOfflineSync,
    updateOfflineSync,
    deleteOfflineSync,
    triggerSync,
    getSyncHistory
  } from '@/api/offlineSync'
  import { getDataSourceList, getDataSourceTables, getDataSourceDatabases } from '@/api/dataSource'
  import { formatDate } from '@/utils/format'
  import { ref, watch } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'

  defineOptions({
    name: 'OfflineSync'
  })

  // 模拟数据
  const mockTableData = [
    {
      id: 1,
      taskName: '用户数据同步',
      sourceType: 'MySQL',
      targetType: 'Elasticsearch',
      status: 'success',
      lastSyncTime: '2026-04-21T10:00:00Z',
      nextSyncTime: '2026-04-22T10:00:00Z',
      createdAt: '2026-04-01T08:00:00Z'
    },
    {
      id: 2,
      taskName: '订单数据备份',
      sourceType: 'MySQL',
      targetType: 'MySQL',
      status: 'running',
      lastSyncTime: '2026-04-21T09:30:00Z',
      nextSyncTime: '2026-04-21T12:30:00Z',
      createdAt: '2026-04-05T08:00:00Z'
    },
    {
      id: 3,
      taskName: '日志采集',
      sourceType: 'API',
      targetType: 'Elasticsearch',
      status: 'pending',
      lastSyncTime: null,
      nextSyncTime: '2026-04-21T14:00:00Z',
      createdAt: '2026-04-10T08:00:00Z'
    },
    {
      id: 4,
      taskName: '商品数据同步',
      sourceType: 'MySQL',
      targetType: 'API',
      status: 'failed',
      lastSyncTime: '2026-04-21T08:00:00Z',
      nextSyncTime: null,
      createdAt: '2026-04-15T08:00:00Z'
    },
    {
      id: 5,
      taskName: '文件归档任务',
      sourceType: 'File',
      targetType: 'File',
      status: 'success',
      lastSyncTime: '2026-04-20T23:00:00Z',
      nextSyncTime: '2026-04-21T23:00:00Z',
      createdAt: '2026-04-18T08:00:00Z'
    },
    {
      id: 6,
      taskName: '第三方API数据拉取',
      sourceType: 'API',
      targetType: 'MySQL',
      status: 'success',
      lastSyncTime: '2026-04-21T06:00:00Z',
      nextSyncTime: '2026-04-22T06:00:00Z',
      createdAt: '2026-04-19T08:00:00Z'
    }
  ]

  const mockHistoryData = [
    { id: 1, executeTime: '2026-04-21T10:00:00Z', status: 'success', duration: 1500, syncCount: 1000, errorMsg: '' },
    { id: 2, executeTime: '2026-04-20T10:00:00Z', status: 'success', duration: 1200, syncCount: 980, errorMsg: '' },
    { id: 3, executeTime: '2026-04-19T10:00:00Z', status: 'failed', duration: 500, syncCount: 50, errorMsg: 'Connection timeout' },
    { id: 4, executeTime: '2026-04-18T10:00:00Z', status: 'success', duration: 1100, syncCount: 950, errorMsg: '' }
  ]

  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(10)
  const tableData = ref([])
  const searchInfo = ref({})
  const multipleSelection = ref([])

  // 表单相关状态
  const formVisible = ref(false)
  const taskFormRef = ref(null)
  const taskForm = ref({
    id: null,
    taskName: '',
    sourceType: '',
    sourceDataSourceId: null,
    sourceDatabase: '',
    sourceTable: '',
    targetType: '',
    targetDataSourceId: null,
    targetDatabase: '',
    targetTable: '',
    mappingRules: '',
    cronExpression: ''
  })

  const taskRules = {
    taskName: [{ required: true, message: '请输入任务名称', trigger: 'blur' }],
    sourceType: [{ required: true, message: '请选择源类型', trigger: 'change' }],
    targetType: [{ required: true, message: '请选择目标类型', trigger: 'change' }]
  }

  const isEdit = ref(false)

  // 数据源列表
  const sourceDataSourceList = ref([])
  const targetDataSourceList = ref([])

  // 数据库列表（从数据源的 database 字段提取）
  const sourceDatabaseList = ref([])
  const targetDatabaseList = ref([])

  // 表列表
  const sourceTableList = ref([])
  const targetTableList = ref([])

  // 获取数据源列表（根据类型筛选）
  const getDataSourcesByType = async (type, target) => {
    if (!type) {
      if (target === 'source') {
        sourceDataSourceList.value = []
      } else {
        targetDataSourceList.value = []
      }
      return
    }
    try {
      const res = await getDataSourceList({ page: 1, pageSize: 100, keyword: '', type: type })
      if (res.code === 0) {
        const list = res.data.list || []
        if (target === 'source') {
          sourceDataSourceList.value = list.filter(item => item.status === 1)
          // 清除已选数据源（如果类型变了）
          if (taskForm.value.sourceType !== type) {
            taskForm.value.sourceDataSourceId = null
          }
          // 清除数据库和表
          taskForm.value.sourceDatabase = ''
          taskForm.value.sourceTable = ''
          sourceDatabaseList.value = []
          sourceTableList.value = []
        } else {
          targetDataSourceList.value = list.filter(item => item.status === 1)
          if (taskForm.value.targetType !== type) {
            taskForm.value.targetDataSourceId = null
          }
          taskForm.value.targetDatabase = ''
          taskForm.value.targetTable = ''
          targetDatabaseList.value = []
          targetTableList.value = []
        }
      }
    } catch (e) {
      console.error('获取数据源列表失败:', e)
    }
  }

  // 监听源类型变化
  watch(() => taskForm.value.sourceType, (newVal) => {
    getDataSourcesByType(newVal, 'source')
  })

  // 监听目标类型变化
  watch(() => taskForm.value.targetType, (newVal) => {
    getDataSourcesByType(newVal, 'target')
  })

  // 获取表列表
  const loadTableList = async (dataSourceId, database, target) => {
    // 防护：没有选中有效数据源时不请求（防止 dataSourceId=undefined/null/"null"）
    const validId = dataSourceId && dataSourceId !== 'undefined' && dataSourceId !== 'null' && dataSourceId !== ''
    if (!validId) {
      if (target === 'source') {
        sourceTableList.value = []
      } else {
        targetTableList.value = []
      }
      return
    }
    try {
      const res = await getDataSourceTables(dataSourceId, database)
      if (res.code === 0) {
        if (target === 'source') {
          sourceTableList.value = res.data.tables || []
          // 清除已选表（如果数据源或库变了）
          if (taskForm.value.sourceTable) {
            taskForm.value.sourceTable = ''
          }
        } else {
          targetTableList.value = res.data.tables || []
          if (taskForm.value.targetTable) {
            taskForm.value.targetTable = ''
          }
        }
      }
    } catch (e) {
      console.error('获取表列表失败:', e)
    }
  }

  // 填充数据库列表（调用 API 获取数据源实例下的所有数据库）
  const fillDatabaseList = async (dataSourceId, target) => {
    if (!dataSourceId) {
      if (target === 'source') {
        sourceDatabaseList.value = []
      } else {
        targetDatabaseList.value = []
      }
      return
    }
    try {
      const res = await getDataSourceDatabases(dataSourceId)
      if (res.code === 0) {
        const databases = res.data.databases || []
        if (target === 'source') {
          sourceDatabaseList.value = databases
          // 自动选中第一个库（仅在没有选中库的情况下）
          if (databases.length > 0 && !taskForm.value.sourceDatabase) {
            taskForm.value.sourceDatabase = databases[0]
          }
        } else {
          targetDatabaseList.value = databases
          if (databases.length > 0 && !taskForm.value.targetDatabase) {
            taskForm.value.targetDatabase = databases[0]
          }
        }
      }
    } catch (e) {
      console.error('获取数据库列表失败:', e)
    }
  }

  // 监听源数据源变化 - 只在清空时处理
  // handleSourceDataSourceChange 已处理 fillDatabaseList 调用
  watch(() => taskForm.value.sourceDataSourceId, (newVal) => {
    if (!newVal) {
      sourceDatabaseList.value = []
      taskForm.value.sourceDatabase = ''
      sourceTableList.value = []
      taskForm.value.sourceTable = ''
    }
  }, { immediate: false })

  // 直接处理源数据源变化
  const handleSourceDataSourceChange = (name) => {
    if (name && typeof name === 'string') {
      const selectedItem = sourceDataSourceList.value.find(item => item.name === name)
      if (selectedItem) {
        taskForm.value.sourceDataSourceId = selectedItem.name
        fillDatabaseList(selectedItem.ID, 'source')
      }
    } else {
      taskForm.value.sourceDataSourceId = null
      sourceDatabaseList.value = []
      taskForm.value.sourceDatabase = ''
    }
  }

  // 监听目标数据源变化 - 只处理清空逻辑
  // 注意：handleTargetDataSourceChange 已处理 fillDatabaseList 调用
  watch(() => taskForm.value.targetDataSourceId, (newVal, oldVal) => {
    if (!newVal) {
      targetDatabaseList.value = []
      taskForm.value.targetDatabase = ''
      targetTableList.value = []
      taskForm.value.targetTable = ''
    }
  })

  // 直接处理目标数据源变化
  const handleTargetDataSourceChange = (name) => {
    if (name && typeof name === 'string') {
      const selectedItem = targetDataSourceList.value.find(item => item.name === name)
      if (selectedItem) {
        taskForm.value.targetDataSourceId = selectedItem.name
        fillDatabaseList(selectedItem.ID, 'target')
      }
    } else {
      taskForm.value.targetDataSourceId = null
      targetDatabaseList.value = []
      taskForm.value.targetDatabase = ''
    }
  }

  // 监听源数据库变化 - 加载表列表
  watch(() => taskForm.value.sourceDatabase, (newVal) => {
    if (newVal && taskForm.value.sourceDataSourceId) {
      // sourceDataSourceId 存的是 name，需要查找对应的 ID
      const item = sourceDataSourceList.value.find(i => i.name === taskForm.value.sourceDataSourceId)
      if (item) {
        loadTableList(item.ID, newVal, 'source')
      }
    } else {
      sourceTableList.value = []
      taskForm.value.sourceTable = ''
    }
  })

  // 监听目标数据库变化 - 加载表列表
  watch(() => taskForm.value.targetDatabase, (newVal) => {
    if (newVal && taskForm.value.targetDataSourceId) {
      const item = targetDataSourceList.value.find(i => i.name === taskForm.value.targetDataSourceId)
      if (item) {
        loadTableList(item.ID, newVal, 'target')
      }
    } else {
      targetTableList.value = []
      taskForm.value.targetTable = ''
    }
  })

  // 打开表单（新增）
  const goCreate = () => {
    isEdit.value = false
    taskForm.value = {
      id: null,
      taskName: '',
      sourceType: '',
      sourceDataSourceId: null,
      sourceDatabase: '',
      sourceTable: '',
      targetType: '',
      targetDataSourceId: null,
      targetDatabase: '',
      targetTable: '',
      mappingRules: '',
      cronExpression: ''
    }
    sourceDataSourceList.value = []
    targetDataSourceList.value = []
    sourceDatabaseList.value = []
    targetDatabaseList.value = []
    sourceTableList.value = []
    targetTableList.value = []
    formVisible.value = true
  }

  // 打开表单（编辑）
  const goEdit = (row) => {
    isEdit.value = true
    taskForm.value = { ...row }
    formVisible.value = true
  }

  // 关闭表单
  const closeForm = () => {
    formVisible.value = false
    taskForm.value = {
      id: null,
      taskName: '',
      sourceType: '',
      sourceDataSourceId: null,
      sourceDatabase: '',
      sourceTable: '',
      targetType: '',
      targetDataSourceId: null,
      targetDatabase: '',
      targetTable: '',
      mappingRules: '',
      cronExpression: ''
    }
    sourceDataSourceList.value = []
    targetDataSourceList.value = []
    sourceDatabaseList.value = []
    targetDatabaseList.value = []
    sourceTableList.value = []
    targetTableList.value = []
  }

  // 保存任务
  const saveTask = async () => {
    if (!taskFormRef.value) return
    await taskFormRef.value.validate((valid) => {
      if (valid) {
        ElMessage({
          type: 'success',
          message: isEdit.value ? '修改成功' : '创建成功'
        })
        closeForm()
        getTableData()
      }
    })
  }

  // 历史对话框相关
  const historyDialogVisible = ref(false)
  const historyTableData = ref([])
  const currentTaskId = ref(null)

  // 状态相关
  const getStatusType = (status) => {
    const typeMap = {
      success: 'success',
      running: 'warning',
      pending: 'info',
      failed: 'danger'
    }
    return typeMap[status] || 'info'
  }

  const getStatusLabel = (status) => {
    const labelMap = {
      success: '成功',
      running: '运行中',
      pending: '待执行',
      failed: '失败'
    }
    return labelMap[status] || status
  }

  // 获取表格数据
  const getTableData = async () => {
    let filteredData = [...mockTableData]
    if (searchInfo.value.taskName) {
      filteredData = filteredData.filter(item => item.taskName.includes(searchInfo.value.taskName))
    }
    if (searchInfo.value.status) {
      filteredData = filteredData.filter(item => item.status === searchInfo.value.status)
    }
    total.value = filteredData.length
    const start = (page.value - 1) * pageSize.value
    const end = start + pageSize.value
    tableData.value = filteredData.slice(start, end)
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

  // 删除任务
  const deleteTask = async (row) => {
    ElMessageBox.confirm('确定要删除该任务吗?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      getTableData()
    })
  }

  // 触发同步
  const handleTriggerSync = async (row) => {
    ElMessageBox.confirm(`确定要触发同步任务 "${row.taskName}" 吗?`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(async () => {
      ElMessage({
        type: 'success',
        message: '触发成功'
      })
      getTableData()
    })
  }

  // 查看历史
  const viewHistory = async (row) => {
    currentTaskId.value = row.id
    historyTableData.value = mockHistoryData
    historyDialogVisible.value = true
  }

  const handleHistoryDialogClose = () => {
    historyDialogVisible.value = false
    historyTableData.value = []
  }

  // 初始化
  getTableData()
</script>

<style lang="scss" scoped>
.form-view {
  .task-form {
    padding: 20px;
    background: #fff;
    border-radius: 8px;
  }

  .full-width {
    width: 100%;
  }

  .form-tip {
    font-size: 12px;
    color: #909399;
    margin-top: 4px;
    line-height: 1.4;
  }

  :deep(.el-divider--horizontal) {
    margin: 16px 0;
  }

  :deep(.el-card) {
    border: 1px solid #e4e7ed;

    .card-header {
      font-weight: 600;
      color: #303133;
    }
  }

  :deep(.el-card + .el-card) {
    margin-left: 16px;
  }

  .source-card,
  .target-card {
    :deep(.el-card__header) {
      padding: 12px 20px;
      background: #f5f7fa;
    }

    :deep(.el-card__body) {
      padding: 20px;
    }
  }
}

.error-text {
  color: #f56c6c;
}
</style>
