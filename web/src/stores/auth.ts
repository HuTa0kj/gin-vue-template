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

    checkAuth() {
      return this.isAuthenticated
    },

    async clearSession() {
      try {
        await axios.get('/api/login/logout', {
          withCredentials: true
        })
        await this.setAuthenticated(false)
        return true
      } catch (error) {
        console.error('Logout failed:', error)
        return false
      }
    }
  }
})
