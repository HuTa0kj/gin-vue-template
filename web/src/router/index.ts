import { createRouter, createWebHistory } from 'vue-router'
import progress from '@bassist/progress'
import routes from './routes'
import { APP_NAME } from '@/constants'
import { useAuthStore } from '@/stores/auth'

progress.configure({ showSpinner: false })
progress.setColor('var(--c-brand)')

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  scrollBehavior: (to, from, savedPosition) => {
    return savedPosition ? savedPosition : { top: 0, left: 0 }
  },
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    next('/login')
  } else if (to.meta.requiresAdmin && authStore.userRole < 10) {
    next('/dashboard')
  } else if (to.meta.requiresSuperAdmin && authStore.userRole < 100) {
    next('/dashboard')
  } else {
    next()
  }
})

router.afterEach((to) => {
  const { title } = to.meta
  document.title = title ? `${title} - ${APP_NAME}` : APP_NAME
  progress.done()
})

export default router
