<template>
  <div class="user-management-container">
    <el-card class="user-management-card">
      <template #header>
        <div class="card-header">
          <span>用户管理</span>
        </div>
      </template>

      <!-- 查询区域 -->
      <div class="search-area">
        <el-form :inline="true" :model="searchForm">
          <el-form-item label="用户名">
            <el-input
              v-model="searchForm.username"
              placeholder="请输入用户名"
              clearable
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleSearch">查询</el-button>
            <el-button @click="handleReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>

      <!-- 用户信息表格 -->
      <el-table :data="userList" border style="width: 100%">
        <el-table-column prop="username" label="用户名" width="180" />
        <el-table-column label="用户权限" width="120">
          <template #default="{ row }">
            <el-tag :type="row.role >= 100 ? 'danger' : row.role >= 10 ? 'warning' : 'success'">
              {{ row.role >= 100 ? '超级管理员' : row.role >= 10 ? '管理员' : '普通用户' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="注册时间" />
        <el-table-column prop="lastLoginTime" label="最后登录时间" />
        <el-table-column label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status ? 'success' : 'danger'">
              {{ row.status ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button type="primary" size="small" @click="handleResetPassword(row)">
              重置密码
            </el-button>
            <el-button type="warning" size="small" @click="handleChangeRole(row)">
              修改权限
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

interface UserInfo {
  id: number
  username: string
  role: number
  createTime: string
  lastLoginTime: string
  status: boolean
}

const searchForm = reactive({
  username: ''
})

const userList = ref<UserInfo[]>([])

const handleSearch = async () => {
  try {
    const response = await fetch('/api/user/search', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(searchForm)
    })
    
    const data = await response.json()
    if (data.code === 2000 && data.status === 'ok') {
      userList.value = data.data
    } else {
      ElMessage.error(data.msg || '查询失败')
    }
  } catch (error) {
    ElMessage.error('查询失败')
  }
}

const handleReset = () => {
  searchForm.username = ''
  userList.value = []
}

const handleResetPassword = async (user: UserInfo) => {
  try {
    const response = await fetch('/api/user/reset-password', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        user_id: user.id
      })
    })
    
    const data = await response.json()
    if (data.code === 2000 && data.status === 'ok') {
      ElMessage.success('密码重置成功')
    } else {
      ElMessage.error(data.msg || '密码重置失败')
    }
  } catch (error) {
    ElMessage.error('密码重置失败')
  }
}

const handleChangeRole = async (user: UserInfo) => {
  try {
    const response = await fetch('/api/user/change-role', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        user_id: user.id,
        new_role: user.role === 100 ? 10 : 100
      })
    })
    
    const data = await response.json()
    if (data.code === 2000 && data.status === 'ok') {
      ElMessage.success('权限修改成功')
      // 刷新用户列表
      handleSearch()
    } else {
      ElMessage.error(data.msg || '权限修改失败')
    }
  } catch (error) {
    ElMessage.error('权限修改失败')
  }
}

const fetchAllUsers = async () => {
  try {
    const response = await fetch('/api/root/user/all')
    const data = await response.json()
    if (data.code === 2000 && data.status === 'ok') {
      userList.value = data.users.map(user => ({
        id: user.id,
        username: user.username,
        role: Number(user.role),
        createTime: new Date(user.register_time).toLocaleString(),
        lastLoginTime: new Date(user.last_login_time).toLocaleString(),
        status: user.status
      }))
    } else {
      ElMessage.error(data.msg || '获取用户列表失败')
    }
  } catch (error) {
    ElMessage.error('获取用户列表失败')
  }
}

onMounted(() => {
  fetchAllUsers()
})
</script>

<style scoped>
.user-management-container {
  padding: 20px;
}

.user-management-card {
  max-width: 1200px;
  margin: 0 auto;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.search-area {
  margin-bottom: 20px;
}
</style>