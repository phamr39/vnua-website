import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'HomePage',
    component: () => import('@/pages/home/HomePage.vue')
  },
  {
    path: '/gioi-thieu',
    name: 'About',
    component: () => import('@/pages/about/About.vue')
  },
  {
    path: '/san-pham',
    name: 'Products',
    component: () => import('@/pages/product/Products.vue')
  },
  {
    path: '/sp/:name',
    name: 'ProductDetail',
    component: () => import('@/pages/product/ProductDetail.vue')
  },
  {
    path: '/dao-tao-chuyen-giao',
    name: 'Training',
    component: () => import('@/pages/train/Training.vue')
  },
  {
    path: '/dao-tao-chuyen-giao/chuyen-giao-cong-nghe',
    name: 'TechTrain',
    component: () => import('@/pages/train/TechTrain.vue')
  },
  {
    path: '/dang-ky-dai-ly',
    name: 'SalesAgent',
    component: () => import('@/pages/sale/SalesAgent.vue')
  },
  {
    path: '/su-kien',
    name: 'Event',
    component: () => import('@/pages/event/Event.vue')
  },
  {
    path: '/tin-tuc',
    name: 'News',
    component: () => import('@/pages/news/News.vue')
  },
  {
    path: '/chi-tiet-tin-tuc/:id',
    name: 'NewsDetail',
    component: () => import('@/pages/news/NewsDetail.vue')
  },
  {
    path: '/videos',
    name: 'Video',
    component: () => import('@/pages/video/Video.vue')
  },
  {
    path: '/lien-he',
    name: 'Contact',
    component: () => import('@/pages/contact/Contact.vue')
  },
  {
    path: '/quyen-loi',
    name: 'CustomPolicy',
    component: () => import('@/pages/policy/CustomPolicy.vue')
  },
  {
    path: '/doi-tra',
    name: 'ReturnPolicy',
    component: () => import('@/pages/policy/ReturnPolicy.vue')
  },
  {
    path: '/mua-hang',
    name: 'ShoppingPolicy',
    component: () => import('@/pages/policy/ShoppingPolicy.vue')
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
