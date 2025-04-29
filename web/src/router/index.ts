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

router.beforeEach(async (to, from, next) => {
  progress.start()
  const authStore = useAuthStore()
  
  if (to.meta.requiresAuth) {
    await authStore.setAuthenticated(true)
    if (!authStore.checkAuth()) {
      next({ name: 'Login' })
      return
    }

    if (to.meta.requiresAdmin && authStore.userRole < 10) {
      next({ name: 'Dashboard' })
      return
    }
  }
  next()
})

router.afterEach((to) => {
  const { title } = to.meta
  document.title = title ? `${title} - ${APP_NAME}` : APP_NAME
  progress.done()
})

export default router
