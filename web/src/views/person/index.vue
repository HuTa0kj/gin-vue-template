<template>
  <div class="person-container">
    <el-card class="person-card">
      <template #header>
        <div class="card-header">
          <span>个人信息</span>
        </div>
      </template>

      <el-form label-width="120px" class="person-form">
        <el-form-item label="用户名">
          <el-input v-model="userInfo.username" disabled />
        </el-form-item>
        
        <el-form-item label="用户角色">
          <el-tag :type="userInfo.role >= 100 ? 'danger' : userInfo.role >= 10 ? 'warning' : 'success'">
            {{ userInfo.role >= 100 ? '超级管理员' : userInfo.role >= 10 ? '管理员' : '普通用户' }}
          </el-tag>
        </el-form-item>

        <el-form-item label="注册时间">
          <span>{{ userInfo.createTime }}</span>
        </el-form-item>

        <el-form-item label="最后登录时间">
          <span>{{ userInfo.lastLoginTime }}</span>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="handleUpdatePassword">修改密码</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <!-- 修改密码对话框 -->
    <el-dialog
      v-model="dialogVisible"
      title="修改密码"
      width="500px"
    >
      <el-form
        ref="passwordFormRef"
        :model="passwordForm"
        :rules="passwordRules"
        label-width="100px"
      >
        <el-form-item label="原密码" prop="oldPassword">
          <el-input
            v-model="passwordForm.oldPassword"
            type="password"
            show-password
            placeholder="请输入原密码"
          />
        </el-form-item>
        <el-form-item label="新密码" prop="newPassword">
          <el-input
            v-model="passwordForm.newPassword"
            type="password"
            show-password
            placeholder="请输入新密码"
          />
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input
            v-model="passwordForm.confirmPassword"
            type="password"
            show-password
            placeholder="请确认新密码"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitPasswordUpdate">
            确认
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const dialogVisible = ref(false)
const passwordFormRef = ref<FormInstance>()

const userInfo = reactive({
  username: '',
  role: 1,
  createTime: '',
  lastLoginTime: ''
})

const passwordForm = reactive({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const validatePass = (_rule: any, value: string, callback: Function) => {
  if (value === '') {
    callback(new Error('请输入密码'))
  } else if (value.length < 6) {
    callback(new Error('密码长度不能小于6位'))
  } else {
    callback()
  }
}

const validatePass2 = (_rule: any, value: string, callback: Function) => {
  if (value === '') {
    callback(new Error('请再次输入密码'))
  } else if (value !== passwordForm.newPassword) {
    callback(new Error('两次输入密码不一致'))
  } else {
    callback()
  }
}

const passwordRules = reactive<FormRules>({
  oldPassword: [
    { required: true, message: '请输入原密码', trigger: 'blur' }
  ],
  newPassword: [
    { validator: validatePass, trigger: 'blur' }
  ],
  confirmPassword: [
    { validator: validatePass2, trigger: 'blur' }
  ]
})

const fetchUserInfo = async () => {
  try {
    const response = await fetch('/api/user/info')
    if (response.ok) {
      const data = await response.json()
      if (data.status === 'ok') {
        userInfo.username = data.username
        userInfo.role = data.role
        userInfo.createTime = data.register_time
        userInfo.lastLoginTime = data.last_login_time
      } else {
        ElMessage.error(data.msg || '获取用户信息失败')
      }
    }
  } catch (error) {
    ElMessage.error('获取用户信息失败')
  }
}

const handleUpdatePassword = () => {
  dialogVisible.value = true
}

const submitPasswordUpdate = async () => {
  if (!passwordFormRef.value) return

  await passwordFormRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const response = await fetch('/api/user/password/update', {
          method: 'PATCH',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            old_password: passwordForm.oldPassword,
            new_password: passwordForm.newPassword
          })
        })
        
        const data = await response.json()
        if (data.status === 'ok') {
          ElMessage.success('密码修改成功')
          dialogVisible.value = false
          // 清空表单
          passwordForm.oldPassword = ''
          passwordForm.newPassword = ''
          passwordForm.confirmPassword = ''
        } else {
          ElMessage.error(data.msg || '密码修改失败')
        }
      } catch (error) {
        ElMessage.error('密码修改失败')
      }
    }
  })
}

onMounted(() => {
  fetchUserInfo()
})
</script>

<style scoped>
.person-container {
  padding: 20px;
}

.person-card {
  max-width: 800px;
  margin: 0 auto;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.person-form {
  max-width: 600px;
  margin: 0 auto;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style>