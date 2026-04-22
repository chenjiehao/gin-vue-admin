<template>
  <div>
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
        <el-button icon="plus" type="primary" @click="openDrawer('create')">新增任务</el-button>
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
            <el-button icon="edit" type="primary" link @click="openDrawer('edit', scope.row)">编辑</el-button>
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

    <!-- 新增/编辑抽屉 -->
    <el-drawer v-model="drawerVisible" :title="drawerTitle" size="50%" :before-close="handleDrawerClose" destroy-on-close>
      <el-form ref="taskFormRef" :model="taskForm" :rules="taskRules" label-width="120px">
        <el-form-item label="任务名称" prop="taskName">
          <el-input v-model="taskForm.taskName" placeholder="请输入任务名称" />
        </el-form-item>
        <el-form-item label="源类型" prop="sourceType">
          <el-select v-model="taskForm.sourceType" placeholder="请选择源类型">
            <el-option label="MySQL" value="MySQL" />
            <el-option label="PostgreSQL" value="PostgreSQL" />
            <el-option label="API" value="API" />
            <el-option label="File" value="File" />
          </el-select>
        </el-form-item>
        <el-form-item label="源配置" prop="sourceConfig">
          <el-input v-model="taskForm.sourceConfig" type="textarea" placeholder="请输入源配置(JSON)" :rows="3" />
        </el-form-item>
        <el-form-item label="目标类型" prop="targetType">
          <el-select v-model="taskForm.targetType" placeholder="请选择目标类型">
            <el-option label="MySQL" value="MySQL" />
            <el-option label="Elasticsearch" value="Elasticsearch" />
            <el-option label="API" value="API" />
            <el-option label="File" value="File" />
          </el-select>
        </el-form-item>
        <el-form-item label="目标配置" prop="targetConfig">
          <el-input v-model="taskForm.targetConfig" type="textarea" placeholder="请输入目标配置(JSON)" :rows="3" />
        </el-form-item>
        <el-form-item label="字段映射" prop="mappingRules">
          <el-input v-model="taskForm.mappingRules" type="textarea" placeholder="请输入字段映射规则(JSON)" :rows="3" />
        </el-form-item>
        <el-form-item label="CRON表达式" prop="cronExpression">
          <el-input v-model="taskForm.cronExpression" placeholder="如: 0 0 * * * * (每天零点)" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div style="flex: 1">
          <el-button @click="handleDrawerClose">取消</el-button>
          <el-button type="primary" @click="saveTask">确定</el-button>
        </div>
      </template>
    </el-drawer>

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
  import { formatDate } from '@/utils/format'
  import { ref } from 'vue'
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

  // 抽屉相关
  const drawerVisible = ref(false)
  const drawerTitle = ref('')
  const dialogFlag = ref('')
  const taskFormRef = ref(null)
  const taskForm = ref({
    id: null,
    taskName: '',
    sourceType: '',
    sourceConfig: '',
    targetType: '',
    targetConfig: '',
    mappingRules: '',
    cronExpression: ''
  })

  const taskRules = {
    taskName: [{ required: true, message: '请输入任务名称', trigger: 'blur' }],
    sourceType: [{ required: true, message: '请选择源类型', trigger: 'change' }],
    targetType: [{ required: true, message: '请选择目标类型', trigger: 'change' }]
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
    // 使用模拟数据
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

  // 抽屉操作
  const openDrawer = (flag, row) => {
    dialogFlag.value = flag
    if (flag === 'edit' && row) {
      drawerTitle.value = '编辑任务'
      taskForm.value = {
        id: row.id,
        taskName: row.taskName,
        sourceType: row.sourceType,
        sourceConfig: JSON.stringify({ host: 'localhost', port: 3306 }, null, 2),
        targetType: row.targetType,
        targetConfig: JSON.stringify({ host: 'localhost', port: 9200 }, null, 2),
        mappingRules: JSON.stringify([{ source: 'id', target: '_id' }], null, 2),
        cronExpression: '0 0 * * * *'
      }
    } else {
      drawerTitle.value = '新增任务'
      taskForm.value = {
        id: null,
        taskName: '',
        sourceType: '',
        sourceConfig: '',
        targetType: '',
        targetConfig: '',
        mappingRules: '',
        cronExpression: ''
      }
    }
    drawerVisible.value = true
  }

  const handleDrawerClose = () => {
    drawerVisible.value = false
  }

  const saveTask = async () => {
    if (!taskFormRef.value) return
    await taskFormRef.value.validate(async (valid) => {
      if (valid) {
        ElMessage({
          type: 'success',
          message: dialogFlag.value === 'create' ? '创建成功' : '更新成功'
        })
        drawerVisible.value = false
        getTableData()
      }
    })
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
  .error-text {
    color: #f56c6c;
  }
</style>
