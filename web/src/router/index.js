import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/init',
    name: 'Init',
    component: () => import('@/view/init/index.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/view/login/index.vue')
  },
  {
    path: '/dataIntegration/offlineSyncForm',
    name: 'OfflineSyncForm',
    component: () => import('@/view/dataIntegration/offlineSyncForm.vue')
  },
  {
    path: '/dataIntegration/dataSourceForm',
    name: 'dataSourceForm',
    component: () => import('@/view/dataIntegration/dataSourceForm.vue')
  },
  {
    path: '/:catchAll(.*)',
    meta: {
      closeTab: true
    },
    component: () => import('@/view/error/index.vue')
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
