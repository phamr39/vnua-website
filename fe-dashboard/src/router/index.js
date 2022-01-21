import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  // {
  //   path: '/',
  //   name: 'Dashboard',
  //   component: () => import('@/components/Dashboard.vue')
  // },
  {
    path: '/profiles',
    name: 'Profile',
    component: () => import('@/pages/users/Profile.vue'),
    meta: {
      requiresAuth: true,
      layout: "default",
    },
  },
  {
    path: '/passwords',
    name: 'Password',
    component: () => import('@/pages/users/Password.vue')
  },
  {
    path: '/products',
    name: 'Products',
    component: () => import('@/pages/products/Products.vue')
  },
  {
    path: '/add-product',
    name: 'AddProduct',
    component: () => import('@/pages/products/AddProduct.vue')
  },
  {
    path: '/update-product/:id',
    name: 'UpdateProduct',
    component: () => import('@/pages/products/UpdateProduct.vue')
  },
  {
    path: '/news',
    name: 'News',
    component: () => import('@/pages/posts/News.vue')
  },
  {
    path: '/add-news',
    name: 'AddNews',
    component: () => import('@/pages/posts/AddNews.vue')
  },
  {
    path: '/events',
    name: 'Events',
    component: () => import('@/pages/posts/Events.vue')
  },
  {
    path: '/add-event',
    name: 'AddEvent',
    component: () => import('@/pages/posts/AddEvent.vue')
  },
  {
    path: '/customers',
    name: 'Customers',
    component: () => import('@/pages/customers/Customers.vue')
  },
  {
    path: '/stores',
    name: 'Stores',
    component: () => import('@/pages/customers/Stores.vue')
  },
  {
    path: '/staffs',
    name: 'Staffs',
    component: () => import('@/pages/staffs/Staffs.vue')
  },
  {
    path: '/add-staff',
    name: 'AddStaff',
    component: () => import('@/pages/staffs/AddStaff.vue')
  },
  {
    path: '/videos',
    name: 'Videos',
    component: () => import('@/pages/videos/Videos.vue')
  },
  {
    path: '/add-video',
    name: 'AddVideo',
    component: () => import('@/pages/videos/AddVideo.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/pages/auth/Login.vue'),
    meta: {
      requiresVisitor: true,
      layout: "landing",
    },
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
