import Vue from 'vue'
import Router from 'vue-router'
import Haoma from '@/components/Haoma'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: '电话号码标记查询',
      component: Haoma
    }
  ]
})
