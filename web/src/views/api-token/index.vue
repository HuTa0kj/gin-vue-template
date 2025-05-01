<template>
  <div class="api-token-container">
    <el-card class="api-token-card">
      <template #header>
        <div class="card-header">
          <span>API Token管理</span>
        </div>
      </template>

      <el-descriptions :column="1" border>
        <el-descriptions-item label="用户名">
          {{ username }}
        </el-descriptions-item>
        <el-descriptions-item label="API Token">
          <div class="token-wrapper">
            <span>{{ token }}</span>
            <el-button
              type="primary"
              link
              :icon="CopyDocument"
              @click="copyToken"
            >
              复制
            </el-button>
          </div>
        </el-descriptions-item>
      </el-descriptions>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { CopyDocument } from '@element-plus/icons-vue'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const username = ref('')
const token = ref('')

const fetchToken = async () => {
  try {
    const response = await fetch('/api/user/token')
    if (response.ok) {
      const data = await response.json()  
      if (data.code === 2000 && data.status === "ok") {
        token.value = data.token
        username.value = data.username
      } else {
        ElMessage.error(data.msg || '获取Token失败')
      }
    }
  } catch (error) {
    ElMessage.error('获取Token失败')
  }
}

const copyToken = async () => {
  try {
    await navigator.clipboard.writeText(token.value)
    ElMessage.success('Token已复制到剪贴板')
  } catch (error) {
    ElMessage.error('复制失败')
  }
}

onMounted(() => {
  fetchToken()
})
</script>

<style scoped>
.api-token-container {
  padding: 20px;
}

.api-token-card {
  max-width: 800px;
  margin: 0 auto;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.token-wrapper {
  display: flex;
  align-items: center;
  gap: 10px;
}
</style>