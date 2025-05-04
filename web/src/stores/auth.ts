import { defineStore } from 'pinia'
import axios from 'axios'

export const useAuthStore = defineStore('auth', {
  persist: true,
  state: () => ({
    isAuthenticated: false,
    userRole: ''
  }),

  actions: {
    async setAuthenticated(value: boolean) {
      if (value) {
        try {
          const response = await axios.get('/api/login/status', {
            withCredentials: true
          })
          this.isAuthenticated = response.data.auth === true
          if (response.data.code === 2000 && response.data.status === "ok") {
            this.userRole = response.data.role
          }
        } catch (error) {
          console.error('Auth check failed:', error)
          this.isAuthenticated = false
        }
      } else {
        this.isAuthenticated = false
        this.userRole = ''
      }
    },

    async checkAuth() {
      try {
        const response = await axios.get('/api/login/status', {
          withCredentials: true
        })
        this.isAuthenticated = response.data.auth === true
        if (response.data.code === 2000 && response.data.status === "ok") {
          this.userRole = response.data.role
        }
        return this.isAuthenticated
      } catch (error) {
        console.error('Auth check failed:', error)
        this.isAuthenticated = false
        return false
      }
    },

    async clearSession() {
      try {
        // 先调用后端登出接口
        await axios.get('/api/logout', {}, {
          withCredentials: true
        })
        // 再清除前端状态
        await this.setAuthenticated(false)
        return true
      } catch (error) {
        console.error('Logout failed:', error)
        return false
      }
    }
  }
})