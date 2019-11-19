import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import Catalog from './components/Catalog.vue'

Vue.config.productionTip = false
Vue.use(VueRouter)

const routes = [
  { path: '/catalog/(.*)', component: Catalog, props: (route) => {
    return {id:route.params.pathMatch}
  }},
]

const router = new VueRouter({routes});


new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
