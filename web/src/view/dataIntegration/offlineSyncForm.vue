<template>
  <div class="gva-container">
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button icon="ArrowLeft" @click="goBack">返回列表</el-button>
      </div>

      <el-form
        ref="taskFormRef"
        :model="taskForm"
        :rules="taskRules"
        label-width="140px"
        class="task-form"
      >
        <el-form-item label="任务名称" prop="taskName">
          <el-input v-model="taskForm.taskName" placeholder="请输入任务名称" />
        </el-form-item>

        <el-divider content-position="left">数据源配置</el-divider>

        <el-form-item label="源类型" prop="sourceType">
          <el-select v-model="taskForm.sourceType" placeholder="请选择源类型" class="full-width">
            <el-option label="MySQL" value="MySQL" />
            <el-option label="PostgreSQL" value="PostgreSQL" />
            <el-option label="API" value="API" />
            <el-option label="File" value="File" />
          </el-select>
        </el-form-item>

        <el-form-item label="源连接配置" prop="sourceConfig">
          <el-input
            v-model="taskForm.sourceConfig"
            type="textarea"
            placeholder="请输入源连接配置(JSON)"
            :rows="4"
          />
        </el-form-item>

        <el-divider content-position="left">目标配置</el-divider>

        <el-form-item label="目标类型" prop="targetType">
          <el-select v-model="taskForm.targetType" placeholder="请选择目标类型" class="full-width">
            <el-option label="MySQL" value="MySQL" />
            <el-option label="Elasticsearch" value="Elasticsearch" />
            <el-option label="API" value="API" />
            <el-option label="File" value="File" />
          </el-select>
        </el-form-item>

        <el-form-item label="目标连接配置" prop="targetConfig">
          <el-input
            v-model="taskForm.targetConfig"
            type="textarea"
            placeholder="请输入目标连接配置(JSON)"
            :rows="4"
          />
        </el-form-item>

        <el-divider content-position="left">同步规则</el-divider>

        <el-form-item label="字段映射规则" prop="mappingRules">
          <el-input
            v-model="taskForm.mappingRules"
            type="textarea"
            placeholder="请输入字段映射规则(JSON)"
            :rows="4"
          />
        </el-form-item>

        <el-form-item label="CRON表达式" prop="cronExpression">
          <el-input v-model="taskForm.cronExpression" placeholder="如: 0 0 * * * (每天零点)" />
          <div class="form-tip">格式: 秒 分 时 日 月 周几，每周日凌晨零点执行请使用: 0 0 0 * * 1</div>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="saveTask">{{ isEdit ? '保存修改' : '创建任务' }}</el-button>
          <el-button @click="goBack">取消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'OfflineSyncForm'
})

const router = useRouter()
const route = useRoute()

const taskFormRef = ref(null)

const isEdit = computed(() => route.query.id !== undefined && route.query.id !== '')
const taskId = computed(() => route.query.id ? parseInt(route.query.id) : null)

// 模拟数据 - 实际应从API获取
const mockTaskData = {
  1: {
    id: 1,
    taskName: '用户数据同步',
    sourceType: 'MySQL',
    sourceConfig: JSON.stringify({ host: '127.0.0.1', port: 3306, database: 'users_db', username: 'root', password: 'xxx' }, null, 2),
    targetType: 'Elasticsearch',
    targetConfig: JSON.stringify({ host: '127.0.0.1', port: 9200, index: 'users' }, null, 2),
    mappingRules: JSON.stringify([
      { source: 'id', target: '_id' },
      { source: 'name', target: 'user_name' },
      { source: 'email', target: 'mail' }
    ], null, 2),
    cronExpression: '0 0 * * * *'
  }
}

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

const goBack = () => {
  router.push({ name: 'offlineSync' })
}

const saveTask = async () => {
  if (!taskFormRef.value) return
  await taskFormRef.value.validate((valid) => {
    if (valid) {
      ElMessage({
        type: 'success',
        message: isEdit.value ? '修改成功' : '创建成功'
      })
      router.push({ name: 'offlineSync' })
    }
  })
}

onMounted(() => {
  // 编辑模式：加载已有数据
  if (isEdit.value && taskId.value && mockTaskData[taskId.value]) {
    const data = mockTaskData[taskId.value]
    taskForm.value = { ...data }
  }
})
</script>

<style lang="scss" scoped>
.task-form {
  max-width: 800px;
  margin: 20px auto;
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

.gva-btn-list {
  margin-bottom: 20px;
}
</style>
