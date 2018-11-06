// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import _ from 'lodash'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-default/index.css'

axios.defaults.baseURL = ''
axios.defaults.timeout = 1000 * 99999999
axios.defaults.headers['Content-Type'] = 'application/json'
Vue.config.productionTip = false

Vue.use(ElementUI)
Object.defineProperty(Vue.prototype, '$_', { value: _ })
Object.defineProperty(Vue.prototype, '$http', { value: axios })

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
