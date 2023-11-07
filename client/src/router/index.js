import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import ItemsList from '../components/items/itemslist.vue'
import Review from '../components/review/review.vue'
import Picklist from '../components/picklist/picklist.vue'
import BackAdmin from '../components/backadmin/backadmin.vue'

const routes = [
  {
    path: '/',
    name: 'Home-component',
    component: Home
  },
  {
    path: '/itemslist',
    name: 'Items-component',
    component: ItemsList
  },
  {
    path: '/review',
    name:'review-component',
    component: Review
  },
  {
    path: '/picklist',
    name: 'picklist-component',
    component: Picklist
  },
  {
    path:'/backadmin',
    name:'backadmin-component',
    component: BackAdmin
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router