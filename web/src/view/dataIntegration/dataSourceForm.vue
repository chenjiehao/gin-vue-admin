<template>
  <div class="gva-container">
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button icon="ArrowLeft" @click="goBack">返回列表</el-button>
      </div>

      <el-form
        ref="dataSourceFormRef"
        :model="dataSourceForm"
        :rules="dataSourceRules"
        label-width="120px"
        class="data-source-form"
      >
        <el-form-item label="数据源名称" prop="name">
          <el-input v-model="dataSourceForm.name" placeholder="请输入数据源名称" />
        </el-form-item>

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

        <el-divider content-position="left">连接配置</el-divider>

        <el-form-item label="主机地址" prop="host">
          <el-input v-model="dataSourceForm.host" placeholder="请输入主机地址" />
        </el-form-item>

        <el-form-item label="端口" prop="port">
          <el-input-number v-model="dataSourceForm.port" :min="1" :max="65535" placeholder="端口" />
        </el-form-item>

        <el-form-item label="数据库名称" prop="database">
          <el-input v-model="dataSourceForm.database" placeholder="请输入数据库名称" />
        </el-form-item>

        <el-form-item label="用户名" prop="username">
          <el-input v-model="dataSourceForm.username" placeholder="请输入用户名" />
        </el-form-item>

        <el-form-item label="密码" prop="password">
          <el-input v-model="dataSourceForm.password" type="password" placeholder="请输入密码" show-password />
        </el-form-item>

        <el-divider content-position="left">状态配置</el-divider>

        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="dataSourceForm.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="备注">
          <el-input
            v-model="dataSourceForm.remark"
            type="textarea"
            placeholder="请输入备注"
            :rows="3"
          />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="saveDataSource">{{ isEdit ? '保存修改' : '创建数据源' }}</el-button>
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
import { createDataSource, updateDataSource, getDataSourceById } from '@/api/dataSource'

defineOptions({
  name: 'dataSourceForm'
})

const router = useRouter()
const route = useRoute()

const dataSourceFormRef = ref(null)

const isEdit = computed(() => route.query.id !== undefined && route.query.id !== '')
const dataSourceId = computed(() => route.query.id ? parseInt(route.query.id) : null)

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

const goBack = () => {
  router.push({ name: 'dataSource' })
}

const saveDataSource = async () => {
  if (!dataSourceFormRef.value) return
  await dataSourceFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        if (isEdit.value) {
          const res = await updateDataSource(dataSourceForm.value)
          if (res.code === 0) {
            ElMessage({ type: 'success', message: '修改成功' })
            router.push({ name: 'dataSource' })
          } else {
            ElMessage({ type: 'error', message: res.msg || '修改失败' })
          }
        } else {
          const res = await createDataSource(dataSourceForm.value)
          if (res.code === 0) {
            ElMessage({ type: 'success', message: '创建成功' })
            router.push({ name: 'dataSource' })
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

onMounted(async () => {
  // 编辑模式：加载已有数据
  if (isEdit.value && dataSourceId.value) {
    try {
      const res = await getDataSourceById({ id: dataSourceId.value })
      if (res.code === 0 && res.data) {
        dataSourceForm.value = { ...res.data }
      }
    } catch (e) {
      console.error('加载数据源失败:', e)
    }
  }
})
</script>

<style lang="scss" scoped>
.data-source-form {
  max-width: 600px;
  margin: 20px auto;
  padding: 20px;
  background: #fff;
  border-radius: 8px;
}

.full-width {
  width: 100%;
}

.gva-btn-list {
  margin-bottom: 20px;
}
</style>
