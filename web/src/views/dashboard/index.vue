<template>
  <div class="dashboard-container">
    <!-- 顶部导航栏 -->
    <el-header class="header">
      <div class="header-left">
        <h2>管理面板</h2>
      </div>
      <div class="header-right">
        <el-dropdown @command="handleCommand">
          <el-avatar :size="32" :icon="UserFilled" />
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="settings">个人设置</el-dropdown-item>
              <el-dropdown-item command="logout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </el-header>

    <el-container class="main-container">
      <!-- 侧边菜单栏 -->
      <el-aside width="200px" class="aside">
        <el-menu
          :default-active="$route.path"
          class="menu"
          :collapse="false"
          router
        >
          <el-menu-item index="/dashboard">
            <el-icon><Monitor /></el-icon>
            <span>系统概览</span>
          </el-menu-item>
          <el-menu-item index="/api-token">
            <el-icon><Key /></el-icon>
            <span>API Token管理</span>
          </el-menu-item>
          <el-menu-item index="/user-management" v-if="authStore.userRole >= 100">
            <el-icon><User /></el-icon>
            <span>用户管理</span>
          </el-menu-item>
          <el-menu-item index="/invite-user" v-if="authStore.userRole >= 10">
            <el-icon><UserFilled /></el-icon>
            <span>用户邀请</span>
          </el-menu-item>
          <el-menu-item index="/settings" v-if="authStore.userRole >=100">
            <el-icon><Setting /></el-icon>
            <span>系统设置</span>
          </el-menu-item>
        </el-menu>
      </el-aside>

      <!-- 主要内容区域 -->
      <el-main class="main">
        <router-view v-if="$route.path !== '/dashboard'" />
        <template v-else>
          <!-- 数据概览卡片 -->
          <el-row :gutter="20" class="data-overview">
            <el-col :span="6">
              <el-card shadow="hover">
                <template #header>
                  <div class="card-header">
                    <el-icon><Document /></el-icon>
                    <span>总日志数</span>
                  </div>
                </template>
                <div class="card-value">{{ totalLogs.toLocaleString() }}</div>
                <div class="card-footer">今日新增: +{{ todayLogs.toLocaleString() }}</div>
              </el-card>
            </el-col>
            <el-col :span="6">
              <el-card shadow="hover">
                <template #header>
                  <div class="card-header">
                    <el-icon><Warning /></el-icon>
                    <span>今日攻击数</span>
                  </div>
                </template>
                <div class="card-value">{{ pendingVuls.toLocaleString() }}</div>
                <div class="card-footer">待验证: {{ unverifiedVuls.toLocaleString() }}</div>
              </el-card>
            </el-col>
            <el-col :span="6">
              <el-card shadow="hover">
                <template #header>
                  <div class="card-header">
                    <el-icon><Monitor /></el-icon>
                    <span>系统状态</span>
                  </div>
                </template>
                <div class="card-value">正常</div>
                <div class="card-footer">CPU使用率: 32%</div>
              </el-card>
            </el-col>
            <el-col :span="6">
              <el-card shadow="hover">
                <template #header>
                  <div class="card-header">
                    <el-icon><Timer /></el-icon>
                    <span>运行时间</span>
                  </div>
                </template>
                <div class="card-value">20天</div>
                <div class="card-footer">上次重启: 2025-03-05</div>
              </el-card>
            </el-col>
          </el-row>

          <!-- 图表区域 -->
          <el-row :gutter="20" class="charts-container">
            <el-col :span="24">
              <el-card shadow="hover">
                <template #header>
                  <div class="card-header">
                    <span>最近日志趋势</span>
                  </div>
                </template>
                <div class="chart-placeholder">
                  此处可以集成 ECharts 等图表库展示实时数据（后续补充）
                </div>
              </el-card>
            </el-col>
          </el-row>
        </template>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { Document, Monitor, Setting, Key, User, UserFilled,  Warning, Timer } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { ref, onMounted } from 'vue'
import axios from 'axios'

const router = useRouter()
const authStore = useAuthStore()

const totalLogs = ref(0)
const todayLogs = ref(0)
const unverifiedVuls = ref(0)
const pendingVuls = ref(0)


const handleCommand = async (command: string) => {
  if (command === 'logout') {
    await handleLogout()
  } else if (command === 'settings') {
    router.push('/person')  // 修改这里，跳转到个人信息页面
  }
}

const handleLogout = async () => {
  const success = await authStore.clearSession()
  if (success) {
    router.push('/login')
  }else {
    router.push('/login')
  }
}
</script>

<style scoped>
.dashboard-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  background-color: #fff;
  border-bottom: 1px solid #dcdfe6;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
}

.header-left h2 {
  margin: 0;
  font-size: 18px;
  color: #303133;
}

.main-container {
  flex: 1;
  overflow: hidden;
}

.aside {
  background-color: #fff;
  border-right: 1px solid #dcdfe6;
}

.menu {
  height: 100%;
  border-right: none;
}

.main {
  background-color: #f5f7fa;
  padding: 20px;
  overflow-y: auto;
}

.data-overview {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 16px;
  color: #303133;
}

.card-value {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
  margin: 16px 0;
}

.card-footer {
  font-size: 14px;
  color: #909399;
}

.charts-container {
  margin-top: 20px;
}

.chart-placeholder {
  height: 300px;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #909399;
  background-color: #f5f7fa;
  border-radius: 4px;
}
</style>