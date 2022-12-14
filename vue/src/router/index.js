import { createRouter, createWebHistory } from 'vue-router'

import AdminLogin from '@/views/AdminLogin.vue'
import UserIndex from '@/views/UserIndex'
import RoleIndex from '@/views/RoleIndex'
import AccessIndex from '@/views/AccessIndex'

const routes = [
  {
    path: "/",
    name: "home",
    component: AdminLogin,
  },
  {
    path: "/user/",
    name: "user_index",
    component: UserIndex,
  },
  {
    path: "/role/",
    name: "role_index",
    component: RoleIndex,
  },
  {
    path: "/access/",
    name: "access_index",
    component: AccessIndex,
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
