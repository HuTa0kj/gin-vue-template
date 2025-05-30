<template>
    <div class="reset-container">
      <el-card class="reset-card">
        <template #header>
          <div class="card-header">
            <h2>重置密码</h2>
          </div>
        </template>
  
        <el-form
          ref="formRef"
          :model="form"
          :rules="rules"
          label-width="0"
          class="reset-form"
        >
          <el-form-item prop="password">
            <el-input
              v-model="form.password"
              type="password"
              placeholder="请输入密码"
              show-password
            />
          </el-form-item>
          <el-form-item prop="confirmPassword">
            <el-input
              v-model="form.confirmPassword"
              type="password"
              placeholder="请确认密码"
              show-password
            />
          </el-form-item>
          <el-form-item>
            <el-button
              type="primary"
              class="submit-btn"
              :loading="loading"
              @click="handleSubmit"
            >
              确认
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref, reactive } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { ElMessage } from 'element-plus'
  import type { FormInstance, FormRules } from 'element-plus'
  import axios from 'axios'
  
  const route = useRoute()
  const router = useRouter()
  const formRef = ref<FormInstance>()
  const loading = ref(false)
  
  const form = reactive({
    password: '',
    confirmPassword: ''
  })
  
  const validatePass = (_rule: any, value: string, callback: Function) => {
    if (value === '') {
      callback(new Error('请输入密码'))
    } else if (value.length < 6) {
      callback(new Error('密码长度不能小于6位'))
    } else {
      if (form.confirmPassword !== '') {
        formRef.value?.validateField('confirmPassword')
      }
      callback()
    }
  }
  
  const validatePass2 = (_rule: any, value: string, callback: Function) => {
    if (value === '') {
      callback(new Error('请再次输入密码'))
    } else if (value !== form.password) {
      callback(new Error('两次输入密码不一致'))
    } else {
      callback()
    }
  }
  
  const rules = reactive<FormRules>({
    password: [{ validator: validatePass, trigger: 'blur' }],
    confirmPassword: [{ validator: validatePass2, trigger: 'blur' }]
  })
  
  const handleSubmit = async () => {
    if (!formRef.value) return
  
    const resetKey = route.query.reset_key
    const username = route.query.username
    if (!resetKey || !username) {
      ElMessage.error('无效的邀请链接')
      return
    }
  
    try {
      await formRef.value.validate()
      loading.value = true
  
      // 调用check接口验证并设置密码
      const checkResponse = await axios.post('/api/user/reset/check', {
        reset_key: resetKey,
        username: username,
        password: form.password
      })
  
      if (checkResponse.data.status === 'ok') {
        ElMessage.success('密码设置成功')
        router.push('/login')
      } else {
        ElMessage.error(checkResponse.data.msg || '验证失败')
      }
    } catch (error: any) {
      if (error.message) {
        ElMessage.error(error.message)
      } else {
        ElMessage.error('请求失败')
      }
    } finally {
      loading.value = false
    }
  }
  </script>
  
  <style scoped>
  .reset-container {
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: #f5f7fa;
  }
  
  .reset-card {
    width: 400px;
  }
  
  .card-header {
    text-align: center;
  }
  
  .card-header h2 {
    margin: 0;
    font-size: 20px;
    color: #303133;
  }
  
  .reset-form {
    padding: 20px 0;
  }
  
  .submit-btn {
    width: 100%;
  }
  </style>