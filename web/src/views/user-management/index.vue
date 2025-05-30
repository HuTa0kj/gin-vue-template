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
        <el-form :inline="true" :model="searchForm" @submit.prevent>
          <el-form-item label="用户名">
            <el-input
              v-model="searchForm.username"
              placeholder="请输入用户名"
              clearable
              @keyup.enter="handleSearch"
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
            <el-button type="primary" size="small" @click="handleEdit(row)">
              编辑
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 编辑对话框 -->
      <el-dialog
        v-model="dialogVisible"
        title="编辑用户信息"
        width="30%"
      >
        <el-form :model="editForm" label-width="80px">
          <el-form-item label="用户名">
            <span>{{ editForm.username }}</span>
          </el-form-item>
          <el-form-item label="用户权限">
            <el-select v-model="editForm.role">
              <el-option :value="10" label="管理员" />
              <el-option :value="1" label="普通用户" />
            </el-select>
          </el-form-item>
          <el-form-item label="状态">
            <el-switch
              v-model="editForm.status"
              :active-value="true"
              :inactive-value="false"
              active-text="启用"
              inactive-text="禁用"
            />
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">取消</el-button>
            <el-button type="primary" @click="handleSave">
              确定
            </el-button>
          </span>
        </template>
      </el-dialog>

      <!-- 重置密码链接对话框 -->
      <el-dialog
        v-model="resetDialogVisible"
        title="重置密码链接"
        width="30%"
      >
        <div class="reset-link-container">
          <el-input
            v-model="resetLink"
            readonly
            type="text"
            class="reset-link-input"
          >
            <template #append>
              <el-button @click="handleCopyLink">
                复制链接
              </el-button>
            </template>
          </el-input>
        </div>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="resetDialogVisible = false">关闭</el-button>
          </span>
        </template>
      </el-dialog>

      <!-- 分页 -->
      <el-pagination
        v-model:current-page="searchForm.page"
        v-model:page-size="searchForm.page_size"
        :page-sizes="[10, 20, 50, 100]"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handlePageChange"
        @size-change="handleSizeChange"
      />
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
  username: '',
  page: 1,
  page_size: 10
})

const total = ref(0)

const fetchAllUsers = async () => {
  try {
    const url = `/api/admin/user/all?page=${searchForm.page}&page_size=${searchForm.page_size}`
    const response = await fetch(url)
    const data = await response.json()
    if (data.status === 'ok') {
      userList.value = data.users.map(user => ({
        id: user.id,
        username: user.username,
        role: Number(user.role),
        createTime: user.register_time,
        lastLoginTime: user.last_login_time,
        status: user.status
      }))
      total.value = data.total
    } else {
      ElMessage.error(data.msg || '获取用户列表失败')
    }
  } catch (error) {
    ElMessage.error('获取用户列表失败')
  }
}

const userList = ref<UserInfo[]>([])

const handleSearch = async () => {
  try {
    const response = await fetch('/api/admin/user/search', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(searchForm)
    })
    
    const data = await response.json()
    if (data.status === 'ok') {
      if (Array.isArray(data.users) && data.users.length > 0) {
        // 处理用户数组
        userList.value = data.users.map(user => ({
          id: user.id,
          username: user.username,
          role: Number(user.role),
          createTime: user.register_time,
          lastLoginTime: user.last_login_time,
          status: user.status
        }))
        total.value = data.total || data.users.length
      } else {
        // 未查询到数据，显示空表格
        userList.value = []
        total.value = 0
      }
    } else {
      ElMessage.error(data.msg || '查询失败')
    }
  } catch (error) {
    ElMessage.error('查询失败')
  }
}

const handlePageChange = (page: number) => {
  searchForm.page = page
  fetchAllUsers()
}

const handleSizeChange = (size: number) => {
  searchForm.page_size = size
  searchForm.page = 1 // Reset to first page when changing page size
  fetchAllUsers()
}

const handleReset = () => {
  searchForm.username = ''
}

const resetLink = ref('')
const resetDialogVisible = ref(false)

const handleResetPassword = async (user: UserInfo) => {
  try {
    const response = await fetch('/api/admin/user/reset/password', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        username: user.username
      })
    })
    
    const data = await response.json()
    if (data.status === 'ok') {
      ElMessage.success('密码重置成功')
      resetLink.value = data.link
      resetDialogVisible.value = true
    } else {
      ElMessage.error(data.msg || '密码重置失败')
    }
  } catch (error) {
    ElMessage.error('密码重置失败')
  }
}

onMounted(() => {
  fetchAllUsers()
})

const dialogVisible = ref(false)
const editForm = reactive({
  id: 0,
  username: '',
  role: 1,
  status: true
})

const handleEdit = (row: UserInfo) => {
  editForm.id = row.id
  editForm.username = row.username
  editForm.role = row.role
  editForm.status = row.status
  dialogVisible.value = true
}

const handleSave = async () => {
  try {
    const response = await fetch('/api/admin/user/update', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        user_id: editForm.id,
        role: editForm.role,
        status: editForm.status
      })
    })
    
    const data = await response.json()
    if (data.status === 'ok') {
      ElMessage.success('更新成功')
      dialogVisible.value = false
      // 刷新用户列表
      fetchAllUsers()
    } else {
      ElMessage.error(data.msg || '更新失败')
    }
  } catch (error) {
    ElMessage.error('更新失败')
  }
}

const handleCopyLink = async () => {
  try {
    await navigator.clipboard.writeText(resetLink.value)
    ElMessage.success('链接已复制到剪贴板')
  } catch (error) {
    ElMessage.error('复制失败')
  }
}
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

.pagination-container {
  margin-top: 20px;
  display: flex;
  justify-content: flex-end;
}

:deep(.el-pagination) {
  padding: 10px 0;
  background-color: #fff;
  border-radius: 4px;
}

:deep(.el-pagination .btn-prev),
:deep(.el-pagination .btn-next) {
  padding: 0 8px;
}

:deep(.el-pagination .el-pager li) {
  margin: 0 4px;
  border-radius: 4px;
}

:deep(.el-pagination .el-pager li.active) {
  background-color: var(--el-color-primary);
  color: #fff;
}

.reset-link-container {
  padding: 10px;
}

.reset-link-input {
  width: 100%;
}
</style>