import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Camera from '../views/Camera.vue'
import Ride from '../views/Ride.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  }, {
    path: '/ride',
    name: 'Ride',
    component: Ride,
    meta: { back: true }
  }, {
    path: '/camera',
    name: 'Camera',
    component: Camera,
    meta: { back: true }
  }, {
    path: '/help',
    name: 'Help',
    component: () => import('../views/Help.vue'),
    meta: { back: true }
  }, {
    path: '/settings',
    name: 'Settings',
    component: () => import('../views/Settings.vue'),
    meta: { back: true }
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
