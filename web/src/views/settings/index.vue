<template>
  <div class="settings-container">
    <el-card class="settings-card">
      <template #header>
        <div class="card-header">
          <span>系统设置</span>
          <el-button type="primary" @click="handleSave">保存设置</el-button>
        </div>
      </template>

      <el-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        label-width="180px"
        class="settings-form"
      >
        <el-form-item label="企业微信Key" prop="wechat_key">
          <el-input
            v-model="formData.wechat_key"
            placeholder="请输入企业微信Key"
            clearable
          />
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'

interface SettingsForm {
  wechat_key: string
}

const formRef = ref<FormInstance>()
const formData = ref<SettingsForm>({
  wechat_key: ''
})

const rules = ref<FormRules>({
  wechat_key: [
    { required: true, message: '请输入企业微信Key', trigger: 'blur' }
  ]
})

// 获取系统设置
const fetchSettings = async () => {
  try {
    const response = await fetch('/api/settings/info')
    const data = await response.json()
    if (data.code === 0) {
      const wechatKeySetting = data.data.find(item => item.setting_key === 'wechat_key')
      formData.value.wechat_key = wechatKeySetting ? wechatKeySetting.setting_value : ''
    } else {
      ElMessage.error(data.message || '获取系统设置失败')
    }
  } catch (error) {
    ElMessage.error('获取系统设置失败')
  }
}

// 保存系统设置
const handleSave = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      try {
        const response = await fetch('/api/setting/save', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(formData.value)
        })
        const data = await response.json()
        if (data.code === 0) {
          ElMessage.success('保存成功')
        } else {
          ElMessage.error(data.message || '保存失败')
        }
      } catch (error) {
        ElMessage.error('保存失败')
      }
    }
  })
}

onMounted(() => {
  fetchSettings()
})
</script>

<style scoped>
.settings-container {
  padding: 20px;
}

.settings-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.settings-form {
  max-width: 800px;
  margin: 0 auto;
}
</style>