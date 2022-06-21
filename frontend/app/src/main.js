import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import BootstrapVue3 from 'bootstrap-vue-3'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue-3/dist/bootstrap-vue-3.css'

// createAppについて(Vue3から変更されたらしい) https://blog.capilano-fw.com/?p=6393
const app = createApp(App)
app.use(BootstrapVue3)
app.use(router)
app.mount('#app')