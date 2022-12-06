

// #ifndef VUE3
import Vue from 'vue'
import App from './App'

Vue.config.productionTip = false

App.mpType = 'app'

const app = new Vue({
    ...App
})
app.$mount()
// #endif

// #ifdef VUE3
import { createSSRApp } from 'vue'
import App from './App.vue'
import api from './common/request/api.js'
// #ifdef H5
import './common/js/vconsole.min.js'
const mode = import.meta.env.MODE;
if(mode==="development"){
	new window.VConsole({ theme: 'white' })
}
// #endif
export function createApp() {
  const app = createSSRApp(App)
  app.config.globalProperties.$api = api
  return {
    app
  }
}
// #endif