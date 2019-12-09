import Vue from 'vue'
import Router from 'vue-router'
import Swapi from '@/components/Swapi'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Swapi',
      component: Swapi
    }
  ]
})
