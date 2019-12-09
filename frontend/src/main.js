import Vue from 'vue'
import VueX from 'vuex'
import VueRouter from 'vue-router'
import App from './App.vue'
import Page from './components/Page.vue'
import Schema from './schema.json'
import store from './store'

Vue.config.productionTip = false
Vue.use(VueX)
Vue.use(VueRouter)

// Build routes using schema
const routes = Schema.map( (i) => {
  const parts = i.key.split("_");
  const named = parts.map( (part, idx) => idx > 0 ? `:${part}` : `${part}` ); // add colons to param name
  const name = i.key;
  const path =`/${named.join('/')}`;
  return {
    path,
    name,
    component: Page,
    props: (route) => { return {name, params:route.params, fields:i.fields} },
  }
});

// Build top level links using schema
const links = Schema.filter( (i) => i.key.split("_").length === 1 )
  .map( (i) => i.key.split("_")[0] );

const router = new VueRouter({routes});

new Vue({
  render: h => h(App, {props:{links}}),
  store,
  router,
}).$mount('#app')
