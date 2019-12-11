<template>
	<div id="main">
		<div class="content">
			<keep-alive>
				<router-view v-if="$route.meta.keepAlive"></router-view>
			</keep-alive>
			<router-view v-if="!$route.meta.keepAlive"></router-view>
		</div>
		<foot-bar v-if="$store.state.footbarShow" :footBarItems="footBarItems" :active="active" @footBarChange="footBarChange"></foot-bar>
	</div>
</template>

<script>
	import footBar from '@/components/foot-bar.vue'
	import {Images} from '../resources/index.js'
	export default{
		components:{
			footBar
		},
		data(){
			return{
				// active:'home',
				footBarItems:[
					{
						name:'home',
						icon:[Images.tabController.ic_home_normal,Images.tabController.ic_home_selected],
						text:'首页',
						to:'/main/home'
					},
					{
						name:'mall',
						icon:[Images.tabController.ic_shop_normal,Images.tabController.ic_shop_selected],
						text:'商城',
						to:'/main/mall'
					},
					{
						name:'merchant',
						icon:[Images.tabController.ic_business_normal,Images.tabController.ic_business_selected],
						text:'商家',
						to:'/main/merchant'
					},
					{
						name:'activity',
						icon:[Images.tabController.ic_activity_normal,Images.tabController.ic_activity_selected],
						text:'活动',
						to:'/main/activity'
					},
					{
						name:'mine',
						icon:[Images.tabController.ic_mine_normal,Images.tabController.ic_mine_selected],
						text:'我的',
						to:'/main/mine'
					}
				]
			}
		},
		methods:{
			footBarChange(active){
				window.console.log(active)
				this.$store.commit('changeTabbarActive',active)
			}
		},
		mounted() {
			this.$store.commit('changeTabbarActive','home')
		},
		computed:{
			active(){
				return this.$store.state.active
			}
		}
	}
</script>

<style scoped>
	.content{
		margin-top: 2.875rem;
		margin-bottom: 3.125rem;
	}
</style>
