<template>
  <div class="invite-user-container">
    <el-card class="invite-user-card">
      <template #header>
        <div class="card-header">
          <span>用户邀请</span>
          <el-button type="primary" @click="handleGenerate">生成邀请链接</el-button>
        </div>
      </template>

      <el-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        label-width="180px"
        class="invite-form"
      >
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="formData.username"
            placeholder="请输入要邀请的用户名"
            clearable
          />
        </el-form-item>

        <el-form-item label="邀请链接" v-if="inviteLink">
          <el-input
            v-model="inviteLink"
            readonly
            placeholder="生成的邀请链接将显示在这里"
          >
            <template #append>
              <el-button @click="copyLink">
                <el-icon><DocumentCopy /></el-icon>
                复制
              </el-button>
            </template>
          </el-input>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { DocumentCopy } from '@element-plus/icons-vue'
import type { FormInstance, FormRules } from 'element-plus'

interface InviteForm {
  username: string
}

const formRef = ref<FormInstance>()
const formData = ref<InviteForm>({
  username: ''
})
const inviteLink = ref('')

const rules = ref<FormRules>({
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '用户名长度应在3-20个字符之间', trigger: 'blur' }
  ]
})

// 生成邀请链接
const handleGenerate = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const response = await fetch('/api/user/invite', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(formData.value)
        })
        const data = await response.json()
        if (data.code === 0) {
          inviteLink.value = data.invite_link
          ElMessage.success(data.msg || '邀请链接生成成功')
        } else {
          ElMessage.error(data.message || '生成邀请链接失败')
        }
      } catch (error) {
        ElMessage.error('生成邀请链接失败')
      }
    }
  })
}

// 复制邀请链接
const copyLink = async () => {
  try {
    await navigator.clipboard.writeText(inviteLink.value)
    ElMessage.success('复制成功')
  } catch (error) {
    ElMessage.error('复制失败')
  }
}
</script>

<style scoped>
.invite-user-container {
  padding: 20px;
}

.invite-user-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.invite-form {
  max-width: 800px;
  margin: 0 auto;
}
</style>