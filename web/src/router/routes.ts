import { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/dashboard',
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/index.vue'),
    meta: {
      title: '登录',
      requiresAuth: false,
    },
  },
  {
    path: '/invite',
    name: 'Invite',
    component: () => import('@/views/invite/index.vue'),
    meta: {
      title: '设置账户密码',
      requiresAuth: false,
    },
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: () => import('@/views/dashboard/index.vue'),
    meta: {
      title: '仪表盘',
      requiresAuth: true,
    },
    children: [
      {
        path: '/api-token',
        name: 'ApiToken',
        component: () => import('@/views/api-token/index.vue'),
        meta: {
          title: 'API Token管理',
          requiresAuth: true,
        },
      },
      {
        path: '/invite-user',
        name: 'InviteUser',
        component: () => import('@/views/invite-user/index.vue'),
        meta: {
          title: '用户邀请',
          requiresAuth: true,
          requiresAdmin: true,
        },
      },
      {
        path: '/person',
        name: 'Person',
        component: () => import('@/views/person/index.vue'),
        meta: {
          title: '个人主页',
          requiresAuth: true,
        },
      },
      {
        path: '/settings',
        name: 'Settings',
        component: () => import('@/views/settings/index.vue'),
        meta: {
          title: '系统设置',
          requiresAuth: true,
          requiresAdmin: true,
        },
      },
      {
        path: '/user-management',
        name: 'UserManagement',
        component: () => import('@/views/user-management/index.vue'),
        meta: {
          title: '用户管理',
          requiresAuth: true,
          requiresSuperAdmin: true,
        },
      },
    ],
  },
]

export default routes
