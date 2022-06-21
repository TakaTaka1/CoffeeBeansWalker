import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/beansInfo',
    name: 'beansInfo',
    component: () => import(/* webpackChunkName: "BeansInfo" */ '@/views/BeansInfoView.vue')

  },
  {
    path: '/editBeansInfo/:id',
    name: 'editBeansInfo',
    component: () => import(/* webpackChunkName: "EditBeansInfo" */ '@/views/EditBeansInfoView.vue')

  },  
  {
    path: '/shopList', // urlについて参考 : https://qiita.com/ksh-fthr/items/a4ac1d04d9923c550cd7
    name: 'shopList',
    component: () => import(/* webpackChunkName: "ShopList" */ '@/views/ShopListView.vue')

  },
  {
    path: '/shopInfo/:id', // urlについて参考 : https://qiita.com/ksh-fthr/items/a4ac1d04d9923c550cd7
    name: 'shopInfo',
    component: () => import(/* webpackChunkName: "ShopInfo" */ '@/views/ShopInfoView.vue')

  },  
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
