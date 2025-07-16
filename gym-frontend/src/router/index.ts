import { createRouter, createWebHistory } from 'vue-router'
import Register from '../views/Register.vue'
import Login from '../views/Login.vue'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
      {
        path: '/register',
        name: 'register',
        component: Register,
      },
      {
        path: '/login',
        name: 'login',
        component: Login,
      },
      {
        path: '/dashboard',
        name: 'dashboard',
        component: () => import('../views/Dashboard.vue'),
        meta: { requiresAuth: true }
      }
  ],
})
// Guardia global
router.beforeEach((to, _from, next) => {
  const isAuthenticated = !!localStorage.getItem('token') // O tu lógica de auth
  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login')
  } else {
    next()
  }
})
export default router
