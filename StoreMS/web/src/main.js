import Vue from 'vue'
import App from './App.vue'
import router from './router/router.js'
import store from './store/store.js'
import axios from 'axios'
import {Lazyload} from 'vant'
import $ from 'jquery'

Vue.prototype.axios = axios
axios.defaults.baseURL='http://47.98.227.139:8088'

Vue.use(Lazyload)

Vue.config.productionTip = false

router.beforeEach((to,from,next)=>{
	window.scrollTo(0,0)
	if (to.matched.some(m=>m.meta.requireAuth)) {
		if (store.state.token) {
			next()
		} else{
			next({path:'/login'})
		}
	} else{
		next()
	}
})

new Vue({
	router,
	store,
  render: h => h(App),
}).$mount('#app')
