import Vue from 'vue'
import Router from 'vue-router'
import Login from './views/Login.vue'
import Register from './views/Register.vue'
import About from './views/About.vue'
import Main from './views/Main.vue'
import Justlook from './views/Justlook.vue'
import Changepassword from './views/Changepassword.vue'
import Home from './views/Home.vue'
import haha from './views/haha.vue'
import ManageBooks from './views/ManageBook.vue'
import ManageUsers from './views/ManageUsers.vue'
import LendReturn from './views/LendReturn.vue'
import AllUsers from './views/AllUsers.vue'



Vue.use(Router)

export default new Router({
	assetsPublicPath:'./',
	routes: [{
			path: '/',
			name: 'login',
			component: Login
		},
		{
			path: '/register',
			name: 'register',
			component: Register
		},
		{
			path: '/justlook',
			name: 'justlook',
			component: Justlook
		},
		{
			path: '/changepassword',
			name: 'changepassword',
			component: Changepassword
		},
		{
			path: '/haha',
			name: 'haha',
			component: haha
		},
		{
			path: '/manageusers',
			name: 'manageusers',
			component: ManageUsers
		},
		{
			path: '/managebook',
			name: 'managebook',
			component: ManageBooks
		},
		{
			path: '/allusers',
			name: 'allusers',
			component: AllUsers
		},
		{
			path: '/lendreturn',
			name: 'lendreturn',
			component: LendReturn
		},
		{
			path: '/main',
			name: 'main',
			component: Main,
			children: [{
					path: 'home',
					component: Home
				},
				{
					path: 'about',
					component: About
				}
			]
		}
	]
})
