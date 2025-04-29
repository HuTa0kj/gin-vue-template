<template>
  <div class="login-container">
    <el-card class="login-card">
      <div class="login-header">
        <img class="logo" src="https://element-plus.org/images/element-plus-logo.svg" alt="logo">
        <h2>系统登录</h2>
      </div>

      <el-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        @keyup.enter="handleSubmit"
      >
        <el-form-item prop="username">
          <el-input
            v-model="formData.username"
            placeholder="请输入用户名"
            :prefix-icon="User"
            size="large"
          />
        </el-form-item>

        <el-form-item prop="password">
          <el-input
            v-model="formData.password"
            type="password"
            placeholder="请输入密码"
            :prefix-icon="Lock"
            size="large"
            show-password
          />
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            :loading="loading"
            class="submit-btn"
            @click="handleSubmit"
          >
            {{ loading ? '登录中...' : '登录' }}
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { User, Lock } from '@element-plus/icons-vue'
import type { FormInstance, FormRules } from 'element-plus'
import axios from 'axios'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()
const loading = ref(false)
const formRef = ref<FormInstance>()

const formData = reactive({
  username: '',
  password: ''
})

const rules = reactive<FormRules>({
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ]
})

const handleSubmit = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        const response = await axios.post('/api/login/verify', formData, {
          headers: {
            'Content-Type': 'application/json'
          },
          withCredentials: true
        })
        if (response.data.code === 2000 && response.data.token.length > 0) {
          await authStore.setAuthenticated(true)
          try {
            await router.push('/dashboard')
            ElMessage.success('登录成功')
          } catch (navigationError) {
            console.error('Navigation failed:', navigationError)
            ElMessage.error('页面跳转失败')
          }
        } else {
          ElMessage.error(response.data.msg || '登录失败')
        }
      } catch (error) {
        console.error('Login failed:', error)
        ElMessage.error('用户名或密码错误')
      } finally {
        loading.value = false
      }
    }
  })
}
</script>

<style scoped>
.login-container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

.login-card {
  width: 400px;
  padding: 20px;
  border-radius: 15px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(10px);
  background: rgba(255, 255, 255, 0.95);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.login-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15);
}

.login-header {
  text-align: center;
  margin-bottom: 30px;
}

.logo {
  width: 80px;
  margin-bottom: 16px;
}

.submit-btn {
  width: 100%;
  padding: 12px;
  font-size: 16px;
  transition: all 0.3s ease;
}

:deep(.el-input__wrapper) {
  padding: 4px 11px;
  border-radius: 8px;
  transition: all 0.3s ease;
  font-size: 14px;

  &:hover {
    box-shadow: 0 0 0 1px var(--el-color-primary) inset;
  }

  &.is-focus {
    box-shadow: 0 0 0 1px var(--el-color-primary) inset;
  }
}

:deep(.el-input__inner) {
  height: 40px;
  color: var(--el-text-color-primary);
  font-size: 14px;

  &::placeholder {
    color: var(--el-text-color-placeholder);
    font-size: 14px;
  }
}
</style>
